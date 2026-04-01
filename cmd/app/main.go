package main

import (
	. "project_go/internal/game/adapter"
	. "project_go/internal/game/application"

	"github.com/gin-gonic/gin"
)

func main() {
	//_ := gin.Default()
	router := gin.Default()

	t := NewTicker(TPS(20))
	a := NewAdapter(t)
	go a.Start()

	router.GET("/connect", a.Connect)

	router.Run(":8080")
}
