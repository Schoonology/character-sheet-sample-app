package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createRouter(productionMode bool) *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("../shared/*")
	router.StaticFile("/main.js", "../shared/main.js")
	router.StaticFile("/styles.css", "../shared/styles.css")

	router.GET("/", showCharacterSheet)
	router.GET("/roll/:dice", getDiceRoll)

	return router
}

func showCharacterSheet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.go.html", gin.H{
		"title": "Character sheet",
	})
}

func getDiceRoll(ctx *gin.Context) {
	dice := ctx.Param("dice")

	var amount int32
	var size int32

	i, _ := fmt.Sscanf(dice, "%dd%d", &amount, &size)

	if i < 2 {
		amount = 1
		fmt.Sscanf(dice, "d%d", &size)
	}

	result, err := rollDice(amount, size)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Printf("Dice: %s => %dd%d => %d\n", dice, amount, size, result)

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
