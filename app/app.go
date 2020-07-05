package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Create router as gin engine type
var (
	router *gin.Engine
)

//Init gin router
func init(){
	router = gin.Default()
}

func ServerApp() {
	// Check current port running
	port := app.Server.ServerRunningPort("8080")

	// Teplate: Load HTML Template global folder
	router.LoadHTMLGlob("views/*/*.html")

	// Static: Load static file folder this is for images, css ans scripts
	router.Static("/assets", "./assets")

	err := router.Run(":" +port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}	
}