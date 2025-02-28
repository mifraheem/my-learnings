package controllers

import "github.com/gin-gonic/gin"

type NotesController struct{}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes)
	notes.POST("/", n.CreateNote)
	notes.GET("/:id", n.GetNote)
	notes.PUT("/:id", n.UpdateNote)
	notes.DELETE("/:id", n.DeleteNote)

}

func (n *NotesController) GetNotes(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all notes",
	})
}

func (n *NotesController) CreateNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create note",
	})
}

func (n *NotesController) GetNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get note",
	})
}
func (n *NotesController) UpdateNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update note",
	})
}
func (n *NotesController) DeleteNote(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete note",
	})
}
