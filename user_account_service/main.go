package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type UserAccount struct {
	AccountNumber string    `json:"account_number"`
	Amount        int       `json:"amount"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}

//
//func mainConsumer(partition int32) {
//	kafka := kafka.NewKafkaConsumer()
//	defer kafka.Close()
//
//	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
//	if err != nil {
//		fmt.Printf("Kafka error: %s\n", err)
//		os.Exit(-1)
//	}
//
//	go ConsumeEvents(consumer)
//
//	fmt.Println("Press [enter] to exit consumer\n")
//	bufio.NewReader(os.Stdin).ReadString('\n')
//	fmt.Println("Terminating...")
//}

func main() {

	var Accounts map[string]UserAccount

	Accounts = map[string]UserAccount{
		"1111": {
			AccountNumber: "1111",
			Amount:        1000000,
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		},
		"2222": {

			AccountNumber: "1111",
			Amount:        0,
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		},
	}

	for i, _ := range Accounts {
		e, err := json.Marshal(Accounts[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(e))
	}

	r := gin.Default()
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/accounts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": Accounts,
		})
	})

	r.Run(":8082")
}
