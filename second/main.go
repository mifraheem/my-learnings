package main

import (
	"fmt"
	"my-app/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8080")
	fmt.Println("Server is running on port 8080\n localhost:8080")
}
