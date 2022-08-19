package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProduceKafka(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	fmt.Println("subham: produce 1")
	fmt.Println("subham: produce 2")

	diler := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}
	p, err := diler.LookupPartition(context.TODO(), "tcp", BROKER, TOPIC, 0)
	if err != nil {
		//log.Fatal("can not lookup partition: ", err.Error())
	}
	fmt.Println("subham: produce 3")
	p.Leader.Host = "localhost"
	conn, err := diler.DialPartition(context.TODO(), "tcp", BROKER, p)
	if err != nil {
		log.Fatal("bbb   subham failed to dial leader:", err)
	}
	fmt.Println("subham: produce 4")
	fmt.Println("Reading partitions again ")
	partitions, err := conn.ReadPartitions()
	if err != nil {
		fmt.Println("ListPartitions: Error while getting partition list: ", err)
		return
	}
	for _, val := range partitions {
		fmt.Println("Partition id", val.ID, "topic is", val.Topic)
	}
	var data string
	for {
		fmt.Println("enter Input")
		fmt.Scanln(&data)
		fmt.Println("enter Input data : ", data)
		msg := kafka.Message{
			Key:   []byte("key-1"),
			Value: []byte(data),
		}
		_, err := conn.WriteMessages(msg)
		if err != nil {
			log.Fatal("ccc   subham failed to write : ", err)
			break
		}
	}

}

func ProduceWindows(wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{BROKER},
		Topic:        TOPIC,
		MaxAttempts:  3,
		RequiredAcks: -1,
		Async:        true,
	})
	writer.AllowAutoTopicCreation = false
	fmt.Println("ProduceWindows 1")
	var data string
	for {
		fmt.Println("enter Input")
		fmt.Scanln(&data)
		msg := kafka.Message{
			Key:   []byte("key-1"),
			Value: []byte(data),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatal("ccc   subham failed to write : ", err.Error(), " -> TOPIC : ", TOPIC)
			break
		}
	}

}
