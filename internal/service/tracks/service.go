package tracks

import (
	"context"

	"github.com/glng-swndru/music-catalog/internal/repository/spotify"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type SpotifyOutbound interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}

type service struct {
	spotifyOutbound SpotifyOutbound
}

func NewService(SpotifyOutbound SpotifyOutbound) *service {
	return &service{spotifyOutbound: SpotifyOutbound}
}
