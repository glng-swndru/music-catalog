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
	membershipRepo "github.com/glng-swndru/music-catalog/internal/repository/memberships"
	"github.com/glng-swndru/music-catalog/internal/repository/spotify"
	trackactivityRepo "github.com/glng-swndru/music-catalog/internal/repository/trackactivities"
	membershipSvc "github.com/glng-swndru/music-catalog/internal/service/memberships"
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
	db.AutoMigrate(&trackactivities.TrackActivity{})

	r := gin.Default()

	httpClient := httpclient.NewClient(&http.Client{})
	spotifyOutbond := spotify.NewSpotifyOutbond(cfg, httpClient)

	membershipRepo := membershipRepo.NewRepository(db)
	trackActivityRepo := trackactivityRepo.NewRepository(db)

	membershipSvc := membershipSvc.NewService(cfg, membershipRepo)
	trackSvc := tracks.NewService(spotifyOutbond, trackActivityRepo)

	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoutes()

	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
