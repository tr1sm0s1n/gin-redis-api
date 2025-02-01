package middlewares

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/tr1sm0s1n/gin-redis-api/models"
)

func Pub(client *redis.Client, cert models.Certificate) error {
	data, err := json.Marshal(cert)
	if err != nil {
		return err
	}
	if err := client.Publish(context.Background(), "redis-queue", data).Err(); err != nil {
		return err
	}
	return nil
}

func Sub(client *redis.Client) {
	defer client.Close()
	sub := client.Subscribe(context.Background(), "redis-queue")

	for {
		log.Println("[*] Waiting for message..")
		msg, err := sub.ReceiveMessage(context.Background())
		if err != nil {
			log.Println(err)
		}
		log.Printf("[*] Received new payload: %s\n", msg.Payload)
	}
}
