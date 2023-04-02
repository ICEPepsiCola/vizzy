package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
	"vizzy/common"
	"vizzy/dao"
	"vizzy/job"
	"vizzy/service"
)

func Bootstrap(mode string) {
	db := initSqlite()
	var page1Dao = dao.Page1Dao{
		DB: db,
	}
	var page2Dao = dao.Page2Dao{
		DB: db,
	}
	var page3Dao = dao.Page3Dao{
		DB: db,
	}
	var page4Dao = dao.Page4Dao{
		DB: db,
	}
	var page5Dao = dao.Page5Dao{
		DB: db,
	}

	initCron(func(cron *cron.Cron) {
		spec := "*/5 * * * *"
		err := cron.AddFunc(spec, func() {
			job.Page1Job(page1Dao)
			job.Page2Job(page2Dao)
			job.Page3Job(page3Dao)
			job.Page4Job(page4Dao)
			job.Page5Job(page5Dao)
		})
		if err != nil {
			logrus.Error(err)
		}
	})
	if mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
		f, err := os.OpenFile("app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		gin.DefaultWriter = io.MultiWriter(f)
		log.SetOutput(f)
	} else {
		// 开发环境，日志输出到控制台
		gin.DefaultWriter = logrus.StandardLogger().Writer()
	}
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	var page0Srv = service.Page0Srv{}
	var page1Srv = service.Page1Srv{Dao: page1Dao}
	var page2Srv = service.Page2Srv{Dao: page2Dao}
	var page3Srv = service.Page3Srv{Dao: page3Dao}
	var page4Srv = service.Page4Srv{Dao: page4Dao}
	var page5Srv = service.Page5Srv{Dao: page5Dao}
	r.GET("/page-0", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page0Srv.Query(cc)
		})
	})
	r.GET("/page-1", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page1Srv.Query(cc)
		})
	})
	r.GET("/page-2", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page2Srv.Query(cc)
		})
	})

	r.GET("/page-3", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page3Srv.Query(cc)
		})
	})

	r.GET("/page-4", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page4Srv.Query(cc)
		})
	})

	r.GET("/page-5", func(c *gin.Context) {
		common.CustomizeHandler(c, func(cc *common.RouteContext) {
			page5Srv.Query(cc)
		})
	})

	err := r.Run(":8888")
	if err != nil {
		logrus.Fatal(err)
	}
}
