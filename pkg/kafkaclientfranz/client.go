package kafkaclientfranz

import (
	"crypto/tls"
	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/plain"
	"net"
	"time"
)

var currentKafkaClient *kgo.Client

func fetchClient(topic, consumerGroup string, info kafkaClientInfo) (*kgo.Client, error) {
	if currentKafkaClient != nil {
		return currentKafkaClient, nil
	}

	seeds := []string{info.broker}

	tlsDialer := &tls.Dialer{NetDialer: &net.Dialer{Timeout: 10 * time.Second}}
	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),

		// SASL Options
		kgo.SASL(plain.Auth{
			User: info.userName,
			Pass: info.password,
		}.AsMechanism()),

		// Configure TLS. Uses SystemCertPool for RootCAs by default.
		kgo.Dialer(tlsDialer.DialContext),

		kgo.ConsumerGroup(consumerGroup),
		kgo.ConsumeTopics(topic),
		//kgo.ConsumeResetOffset(),
		kgo.ClientID(info.clientID),
	}
	c, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	currentKafkaClient = c
	return c, nil
}

type kafkaClientInfo struct {
	broker   string
	userName string
	password string
	clientID string
}
