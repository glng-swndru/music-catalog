package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/music-catalog/internal/configs"
	membershipsHandler "github.com/glng-swndru/music-catalog/internal/handler/memberships"
	"github.com/glng-swndru/music-catalog/internal/models/memberships"
	membershipRepo "github.com/glng-swndru/music-catalog/internal/repository/memberships"
	membershipSvc "github.com/glng-swndru/music-catalog/internal/service/memberships"
	"github.com/glng-swndru/music-catalog/pkg/internalsql"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs",
			"./internal/configs", // for local configs file path
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v\n", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err)
	}
	db.AutoMigrate(&memberships.User{})
	r := gin.Default()

	membershipRepo := membershipRepo.NewRepository(db)

	membershipSvc := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
