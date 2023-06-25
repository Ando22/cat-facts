package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"user-message/backend/config"
	"user-message/backend/database"
	"user-message/backend/models"
	"user-message/backend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetMessages handles the retrieval of messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.GetDB()
	collection := db.Collection(config.CollectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch messages")
		return
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch messages")
		return
	}

	jsonResponse(w, http.StatusOK, messages)
}

// CreateMessage handles the creation of a new message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var message models.Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	message.ID = primitive.NewObjectID()
	message.Created = time.Now().Unix()

	db := database.GetDB()
	collection := db.Collection(config.CollectionName)

	_, err := collection.InsertOne(ctx, message)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create message")
		return
	}

	jsonResponse(w, http.StatusCreated, message)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
