package googlepubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"newsletter-pub/utils/config"
	"sync"

	models_email "newsletter-pub/models/email"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

var appConfig = config.AppCfg

func PublisherLoop(data *models_email.EmailPayload) (string, error) {
	// Create a context and a Pub/Sub client.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, appConfig.PubSub.ProjectId, option.WithCredentialsFile("credential.json"))
	if err != nil {
		return "", fmt.Errorf("failed to create Pub/Sub client: %v", err)
	}

	// Get a reference to the topic.
	topic := client.Topic(appConfig.PubSub.TopicName)

	var wg sync.WaitGroup

	for _, i := range data.RecipientList {

		wg.Add(1)
		go func(i models_email.EmailPayloadList) {
			defer wg.Done()
			stringData, _ := json.Marshal(i)

			// Create a Pub/Sub message.
			msg := &pubsub.Message{
				Data: []byte(stringData),
			}

			// Publish the message.
			result := topic.Publish(ctx, msg)

			// Get the server-generated message ID.
			messageID, err := result.Get(ctx)
			if err != nil {
				fmt.Printf("failed to publish message: %v", err)
			}

			// Your message to publish.
			fmt.Printf("Message published with Message ID: %s\n", messageID)
		}(i)

	}
	wg.Wait()

	return "Success", nil
}

func Publisher(data string) (string, error) {
	// Create a context and a Pub/Sub client.
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, appConfig.PubSub.ProjectId, option.WithCredentialsFile("credential.json"))
	if err != nil {
		return "", fmt.Errorf("failed to create Pub/Sub client: %v", err)
	}

	// Get a reference to the topic.
	topic := client.Topic(appConfig.PubSub.TopicName)

	// Create a Pub/Sub message.
	msg := &pubsub.Message{
		Data: []byte(data),
	}

	// Publish the message.
	result := topic.Publish(ctx, msg)

	// Get the server-generated message ID.
	messageID, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %v", err)
	}

	// Your message to publish.
	fmt.Printf("Message published with Message ID: %s\n", messageID)
	return messageID, nil
}
