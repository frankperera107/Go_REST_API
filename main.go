package main

import (
	"net/http"

	"github.com/frankperera107/golang/rest_api/db"
	"github.com/frankperera107/golang/rest_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}


func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured while retrieving events."})
		return 
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse received data."})
		return 
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()

	if err!= nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}