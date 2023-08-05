package main

import (
	"log"

	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/TuringCup/TuringBackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	reporter, err := reporter.NewGRPCReporter("skywalking-oap:11800")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer reporter.Close()
	tracer, err := go2sky.NewTracer("TuringCup", go2sky.WithReporter(reporter))
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	r.Use(v3.Middleware(r, tracer))
	routes.NewRouter(r)
	r.Run(":5001")
}
