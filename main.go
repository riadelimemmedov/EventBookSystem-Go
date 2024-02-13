package main

import (
	"book_event/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// *main
func main() {
	//Set enviroment mode
	// gin.SetMode(gin.DebugMode)

	//create path instance from gin.Default
	server := gin.Default()

	//api endpoint
	server.GET("events/", getEvents)
	server.POST("event/", createEvent)

	//run the server``
	server.Run(":8080") // http://localhost:8080/
}

// !getEvents
func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

// !createEvent
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) //Check request body,user send required  data or not

	message, isSuccess := checkError(err, "Could not parse request data")

	if !isSuccess {
		context.JSON(http.StatusBadRequest, gin.H{
			"message":   message,
			"isSuccess": isSuccess,
		})
		return
	}

	//referance data to event instance
	event.Id = 1
	event.UserID = 1

	//Save event to database
	event.Save()

	//return created event to client
	context.JSON(http.StatusOK, gin.H{
		"message": "Event created!",
		"event":   event,
	})
}

// !checkError
func checkError(err error, message string) (string, bool) {
	if err != nil {
		return message, false
	}
	return "Success", true
}
