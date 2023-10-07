package main

import (
	"fmt"
	"log"

	"github.com/DeanThompson/ginpprof"
	"github.com/SkyAPM/go2sky"
	v3 "github.com/SkyAPM/go2sky-plugins/gin/v3"
	"github.com/SkyAPM/go2sky/reporter"
	"github.com/TuringCup/TuringBackend/config"
	"github.com/TuringCup/TuringBackend/middleware"
	"github.com/TuringCup/TuringBackend/repository/cache"
	"github.com/TuringCup/TuringBackend/repository/db/dao"
	"github.com/TuringCup/TuringBackend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.InitConfig("")
	fmt.Println(config.Conf.System.Host)
	fmt.Println(config.Conf.System.Port)
	dao.ConnectDB()
	cache.InitCache()
	r := gin.Default()
	reporter, err := reporter.NewGRPCReporter(config.Conf.Skywalking.Host + ":" + config.Conf.Skywalking.Port)
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	defer reporter.Close()
	tracer, err := go2sky.NewTracer("TuringCup", go2sky.WithReporter(reporter))
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
	}
	r.Use(v3.Middleware(r, tracer))
	r.Use(middleware.CorsMiddle())
	routes.NewRouter(r)
	ginpprof.Wrap(r)
	r.Run(config.Conf.System.Host + ":" + config.Conf.System.Port)
}
