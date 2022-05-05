package main

import (
	db2 "gitee.com/zince/micro-service/common/pkg/db"
	"gitee.com/zince/micro-service/user/pkg/model"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"gitee.com/zince/micro-service/user/handler"
	"gitee.com/zince/micro-service/user/subscriber"

	"gitee.com/zince/micro-service/user/proto/user"
)

func main() {
	db := db2.GetDB()
	db.AutoMigrate(&model.User{})

	// New Service
	service := micro.NewService(
		micro.Name("micro.service.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("micro.service.user", service.Server(), new(subscriber.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
