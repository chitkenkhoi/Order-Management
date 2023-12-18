package firebaseconfig

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"log"
)

var Client *firestore.Client

func ConnectDB() {
	opt := option.WithCredentialsFile("firebaseconfig/quan-ly-don-hang-b7bdf-firebase-adminsdk-2vybh-edfac5cf00.json")
	// Initialize the app with the SDK
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}
	Client = client
}
