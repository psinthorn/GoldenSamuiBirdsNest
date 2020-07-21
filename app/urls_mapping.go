package app

import "github.com/psinthorn/GoldenSamuiBirdsNest/controllers"

func urlsMapping() {
	router.GET("/", controllers.Index.Welcome)
}
