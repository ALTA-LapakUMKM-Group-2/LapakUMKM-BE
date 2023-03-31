package helpers

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func PushNotifFirebase(title, body, topic string) error {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	// Define the message to send
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: topic,
		Token: "<client_token>",
	}

	// Send the message
	if _, err := client.Send(ctx, message); err != nil {
		return err
	}

	return nil
}
