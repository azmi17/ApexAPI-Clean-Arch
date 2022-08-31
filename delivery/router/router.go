package router

import (
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kpango/glg"
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)

	//Discard semua output yang dicatat oleh gin karena print out akan dicetak sesuai kebutuhan programmer
	gin.DefaultWriter = ioutil.Discard

	router := gin.Default() //create router engine by default
	api := router.Group("api/v1")

	router.Use(gin.Recovery())

	// ! ADD ANOTHER HANDLER BELOW..
	RegisterHandler(router, api)

	listenerPort := os.Getenv("app.listener_port") // get port from .env

	_ = glg.Logf("[HTTP] Listening at : %s", listenerPort)
	return router.Run(":" + listenerPort)

}
