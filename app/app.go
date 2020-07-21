package app

import (
	"log"

	"github.com/gin-gonic/gin"
	util "github.com/psinthorn/GoldenSamuiBirdsNest/utils"
)

//Create router as gin engine type
var (
	router *gin.Engine
)

//Init gin router
func init() {
	router = gin.Default()
}

func AppStart() {
	// Check current port running
	port := util.ServerPort.PortRunning("8083")

	// Teplate: Load HTML Template global folder
	router.LoadHTMLGlob("views/*/*.html")

	// Static: Load static file folder this is for images, css ans scripts
	router.Static("/assets", "./assets")
	urlsMapping()
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("ListenAndServe:"+port, err)
	}

}
