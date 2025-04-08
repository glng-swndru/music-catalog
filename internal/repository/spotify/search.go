package spotify

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type SpotifySearchResponse struct {
	Tracks SpotifyTracks `json:"tracks"`
}

type SpotifyTracks struct {
	Href     string               `json:"href"`
	Limit    int                  `json:"limit"`
	Next     *string              `json:"next"`
	Offset   int                  `json:"offset"`
	Previous *string              `json:"previous"`
	Total    int                  `json:"total"`
	Items    []SpotifyTrackObject `json:"items"`
}

type SpotifyTrackObject struct {
	Album    SpotifyAlbumObject    `json:"album"`
	Artist   []SpotifyArtistObject `json:"artist"`
	Explicit bool                  `json:"explicit"`
	Href     string                `json:"href"`
	ID       string                `json:"id"`
	Name     string                `json:"name"`
}

type SpotifyAlbumObject struct {
	AlbumType  string              `json:"album_type"`
	TotalTrack int                 `json:"total_track"`
	Images     []SpotifyAlbumImage `json:"images"`
	Name       string              `json:"name"`
}

type SpotifyAlbumImage struct {
	URL string `json:"url"`
}

type SpotifyArtistObject struct {
	Href string `json:"href"`
	Name string `json:"name"`
}

func (o *outbond) Search(ctx context.Context, query string, limit, offset int) (*SpotifySearchResponse, error) {
	url := `https://api.spotify.com/v1/search`
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create search request for Spotify")
		return nil, err
	}

	bearerToken := ""
	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error execute request for spotify")
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifySearchResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal search response from spotify")
		return nil, err
	}
	return &response, nil
}
