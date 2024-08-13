package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL = "localhost:19092"
	topic    = "user_topic"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // specify the kafka brokers []string{"localhost:9092", "localhost:9093"}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		// StartOffset: kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua bán chứng khoán
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

func actionStock(c *gin.Context) {
	s := newStock(c.Query("message"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("stock"),
		Value: []byte(jsonBody),
	}

	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}

// Consumer hóng mua ATC
func RegisterConsumerATC(id int) {
	// group consumer
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, topic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer %d Hong phien ATC", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer %d Error: %v ", id, err)
		}

		fmt.Printf("Consumer(%d) - Message at offset %d: %s = %s\n", id, m.Offset, string(m.Key), string(m.Value))
	}
}

func main() {
	// gin
	r := gin.Default()

	kafkaProducer = getKafkaWriter(kafkaURL, topic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)
	// đăng ký 2 user để mua Stock trong phiên ATC (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)

	r.Run(":8999")
}
