package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var listenAddress = ":8080"
var productionMode = false
var randomOrgApiKey = ""

func init() {
	var present bool
	randomOrgApiKey, present = os.LookupEnv("RANDOM_ORG_API_KEY")
	if !present {
		panic("RANDOM.ORG API key is required.")
	}

	portString, present := os.LookupEnv("PORT")
	if present {
		listenAddress = fmt.Sprintf(":%s", portString)
	}

	productionModeStr, _ := os.LookupEnv("ENV")
	productionMode = productionModeStr == "production"

	if productionMode {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	createRouter(productionMode).Run(listenAddress)
}
