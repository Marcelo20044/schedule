package kafka

import (
	"github.com/Shopify/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer(brokers []string) (*Consumer, error) {
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer}, nil
}

func (c *Consumer) Consume(topic string, handler func(message *sarama.ConsumerMessage)) error {
	partitionList, err := c.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitionList {
		pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				handler(message)
			}
		}(pc)
	}

	return nil
}
