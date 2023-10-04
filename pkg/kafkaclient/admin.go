package kafkaclient

import (
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
)

type AdminClient interface {
	Print()
}

func newAdminClient(c *kgo.Client) (AdminClient, error) {
	client := kadm.NewClient(c)

	kc := &kafkaAdminClient{
		client: client,
	}
	return kc, nil
}

type kafkaAdminClient struct {
	client *kadm.Client
}

func (c *kafkaAdminClient) Print() {

}
