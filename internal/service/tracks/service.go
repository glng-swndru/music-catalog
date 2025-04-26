package tracks

import (
	"context"

	"github.com/glng-swndru/music-catalog/internal/models/trackactivities"
	"github.com/glng-swndru/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
	GetRecommendation(ctx context.Context, limit int, trackID string) (*spotify.SpotifyRecommendationResponse, error)
}

type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackactivities.TrackActivity) error
	Update(ctx context.Context, model trackactivities.TrackActivity) error
	Get(ctx context.Context, UserID uint, SpotifyID string) (*trackactivities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, UserID uint, SpotifyIDs []string) (map[string]trackactivities.TrackActivity, error)
}

type service struct {
	spotifyOutbound     spotifyOutbound
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbound spotifyOutbound, trackActivitiesRepo trackActivitiesRepository) *service {
	return &service{spotifyOutbound: spotifyOutbound, trackActivitiesRepo: trackActivitiesRepo}
}
