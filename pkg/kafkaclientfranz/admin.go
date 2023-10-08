package kafkaclientfranz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"log"
)

type AdminClient interface {
	Print(ctx context.Context)
	Offsets(ctx context.Context, group string) error
}

func newAdminClient(consumerGroup string, c *kgo.Client) (AdminClient, func(), error) {
	client := kadm.NewClient(c)

	kc := &kafkaAdminClient{
		client:        client,
		consumerGroup: consumerGroup,
	}

	return kc, client.Close, nil
}

type kafkaAdminClient struct {
	client        *kadm.Client
	consumerGroup string
}

func (c *kafkaAdminClient) Offsets(ctx context.Context, group string) error {
	offsets, err := c.client.FetchOffsets(ctx, group)
	if err != nil {
		return err
	}

	_ = offsets
	return nil
}

func (c *kafkaAdminClient) Print(ctx context.Context) {
	//c.committedOffsets(ctx)
	c.describeGroups(ctx)

}

func (c *kafkaAdminClient) committedOffsets(ctx context.Context) {
	offsets, err := c.client.ListCommittedOffsets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for topic, offset := range offsets.Offsets() {
		fmt.Printf("=== %s\n", topic)
		for k, ofs := range offset {
			fmt.Printf("    KEY: \t%d\n", k)
			fmt.Printf("    TOPIC: \t%s\n", ofs.Topic)
			fmt.Printf("    PARTITION: \t%d\n", ofs.Partition)
			fmt.Printf("    AT: \t%d\n", ofs.At)
			fmt.Printf("    LEADER EPOCH: \t%d\n", ofs.LeaderEpoch)
			fmt.Printf("    METADATA: \t%s\n", ofs.Metadata)
		}
	}
}

func (c *kafkaAdminClient) describeGroups(ctx context.Context) {
	groups, err := c.client.DescribeGroups(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range groups {
		b, err := json.Marshal(group)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}

	//c.client.FetchOffsetsForTopics()
}
