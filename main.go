package main

import (
	"github.com/TuringCup/TuringBackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.NewRouter(r)
	r.Run()
}
