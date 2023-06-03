package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// This function register the particular adapter based on providerID
func Login() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from Login Handler")
	}
}

func SingUp() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from SingUp Handler")
	}
}

func CreateNotes() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from CreateNotes Handler")
	}
}
func GetNotes() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from GetNotes Handler")
	}
}

func GetNotesByID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from GetNotesByID Handler")
	}
}

func DeleteNotesByID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("Hello from DeleteNotesByID Handler")
	}
}
