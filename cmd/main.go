package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/music-catalog/internal/configs"
	membershipsHandler "github.com/glng-swndru/music-catalog/internal/handler/memberships"
	tracksHandler "github.com/glng-swndru/music-catalog/internal/handler/tracks"
	"github.com/glng-swndru/music-catalog/internal/models/memberships"
	"github.com/glng-swndru/music-catalog/internal/models/trackactivities"
	membershipsRepo "github.com/glng-swndru/music-catalog/internal/repository/memberships"
	"github.com/glng-swndru/music-catalog/internal/repository/spotify"
	trackactivitiesRepo "github.com/glng-swndru/music-catalog/internal/repository/trackactivities"
	membershipsSvc "github.com/glng-swndru/music-catalog/internal/service/memberships"
	"github.com/glng-swndru/music-catalog/internal/service/tracks"
	"github.com/glng-swndru/music-catalog/pkg/httpclient"
	"github.com/glng-swndru/music-catalog/pkg/internalsql"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/", // for local configs file path
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v", err)
	}
	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&trackactivities.TrackActivity{})

	r := gin.Default()

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbound := spotify.NewSpotifyOutbound(cfg, httpClient)

	membershipRepo := membershipsRepo.NewRepository(db)
	trackAvtivitiesRepo := trackactivitiesRepo.NewRepository(db)

	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
	tracksSvc := tracks.NewService(spotifyOutbound, trackAvtivitiesRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	tracksHandler := tracksHandler.NewHandler(r, tracksSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
