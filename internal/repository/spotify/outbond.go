package spotify

import (
	"time"

	"github.com/glng-swndru/music-catalog/internal/configs"
	"github.com/glng-swndru/music-catalog/pkg/httpclient"
)

type outbond struct {
	cfg         *configs.Config
	client      httpclient.HTTPClient
	AccessToken string
	TokenType   string
	ExpiredAt   time.Time
}

func NewSpotifyOutbond(cfg *configs.Config, client httpclient.HTTPClient) *outbond {
	return &outbond{
		cfg:    cfg,
		client: client,
	}
}
