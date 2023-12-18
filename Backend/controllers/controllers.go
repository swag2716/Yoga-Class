package controllers

import (
	"context"
	"net/http"
	"time"
	"yoga-class/database"
	"yoga-class/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var participantsCollection = database.OpenCollection(database.Client, "participantsCollection")

func EnrollController(c *gin.Context) {
	var participant models.Participant

	if err := c.BindJSON(&participant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !participant.Payment {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment is not Done"})
		return
	}

	participant.EnrollDate = time.Now()
	participant.ModifiedDate = participant.EnrollDate

	// Save participant to MongoDB
	result, err := participantsCollection.InsertOne(context.Background(), participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error enrolling participant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"participantId": result.InsertedID})
}

func GetParticipantsController(c *gin.Context) {
	var participants []models.Participant

	cursor, err := participantsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving participants"})
		return
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &participants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding participants"})
		return
	}

	c.JSON(http.StatusOK, participants)
}

func UpdateBatchController(c *gin.Context) {
	id := c.Param("id")

	var updatedParticipant models.Participant
	if err := c.BindJSON(&updatedParticipant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if participant.modifiedDate

	// Find the participant by ID
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"batch":        updatedParticipant.Batch,
			"modifiedDate": time.Now(),
		},
	}

	result, err := participantsCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating participant batch"})
		return
	}

	if result.ModifiedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Batch updated successfully"})
}
