package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	TOPIC  = "TEST"
	BROKER = "localhost:9093"
)

func ConsumeAgain(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println("subham: Consume 1")
	fmt.Println("subham: Consume 2")
	diler := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	p, err := diler.LookupPartition(context.TODO(), "tcp", BROKER, TOPIC, 0)
	if err != nil {
		log.Fatal("can not lookup partition: ", err.Error())
	}
	fmt.Println("subham: Consume 3")
	p.Leader.Host = "localhost"
	conn, err := diler.DialPartition(context.TODO(), "tcp", BROKER, p)
	if err != nil {
		log.Fatal("bbb   subham failed to dial leader:", err)
	}
	fmt.Println("subham: Consume 4")
	fmt.Println("Reading partitions again ")
	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Println("ListPartitions: Error while getting partition list: ", err)
		return
	}
	for _, val := range partitions {
		fmt.Println("Partition id", val.ID, "topic is", val.Topic)
	}
	fmt.Println("subham: Consume 5")
	a, _ := conn.ReadLastOffset()
	conn.Seek(a, 0)
	fmt.Println("laste offset : ", a, kafka.LastOffset)
	for {
		msg, err := conn.ReadMessage(10e6)
		if err != nil {
			log.Fatal("ccc   subham failed to Read : ", err)
			break
		}
		fmt.Println("Partition: ", msg.Partition, " Offset: ", msg.Offset, " Topic: ", msg.Topic, "key : ", string(msg.Key))

		fmt.Println("subham: received msg :", string(msg.Value))
	}

}

func ConsumeWindows(wg *sync.WaitGroup, gpID string) {
	fmt.Println("subham: ConsumeWindows 1 ")
	if wg != nil {
		defer wg.Done()
	}
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{BROKER},
		Topic:       TOPIC,
		GroupID:     gpID,
		StartOffset: kafka.LastOffset,
	})
	fmt.Println("subham: ConsumeWindows 2 ", kafka.LastOffset)
	for {
		msg, err := reader.FetchMessage(context.Background())
		fmt.Println("Starting to fetch from Kafka: reader: ", reader.Config().Topic, reader.Config().GroupID, reader.Config().Brokers)
		if err != nil {
			fmt.Println("Failed to fetch from Kafka: Error: ", err.Error())
			panic("could not read message " + err.Error())
		}
		fmt.Println("Partition: ", msg.Partition, " Offset: ", msg.Offset, " Topic: ", msg.Topic, "key : ", string(msg.Key))

		fmt.Println("subham: received msg :", string(msg.Value))
	}

}
