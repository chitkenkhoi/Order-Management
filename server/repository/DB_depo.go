package repository

import (
	"context"
	"google.golang.org/api/iterator"
	"log"
	FB "server/firebaseconfig"
	model "server/model"
)

func CheckUser(user model.User) bool {
	client := FB.Client
	query := client.Collection("users").
		Where("username", "==", user.Username).
		Where("password", "==", user.Password).
		Documents(context.Background())

	// Check if there is any matching document
	for {
		_, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error querying documents: %v", err)
		}

		// Document with the specified username and password exists

		return true
	}
	return false
}
func GetItems(name string) []model.Item {
	client := FB.Client
	query := client.Collection("items").
		Where("name", "==", name).
		Documents(context.Background())
	// Iterate through the result set
	var items []model.Item
	for {
		doc, err := query.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error querying documents: %v", err)
		}
		var map_item map[string]interface{}
		doc.DataTo(&map_item)
		item := model.MapToItem(map_item, doc.Ref.ID)
		items = append(items, item)
	}
	return items
}
