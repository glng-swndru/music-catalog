package tracks

import (
	"context"

	"github.com/glng-swndru/music-catalog/internal/models/spotify"
	"github.com/glng-swndru/music-catalog/internal/models/trackactivities"
	spotifyRepo "github.com/glng-swndru/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}

	trackIDs := make([]string, len(trackDetails.Tracks.Items))
	for idx, item := range trackDetails.Tracks.Items {
		trackIDs[idx] = item.ID
	}

	trackAtivities, err := s.trackActivityRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from database")
	}

	return modelToResponse(trackDetails, trackAtivities), nil
}

func modelToResponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackactivities.TrackActivity) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
		artistsName := make([]string, len(item.Artists))
		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))
		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL
		}

		items = append(items, spotify.SpotifyTrackObject{
			// album related field
			AlbumType:       item.Album.AlbumType,
			AlbumTotalTrack: item.Album.TotalTrack,
			AlbumImagesURL:  imageUrls,
			AlbumName:       item.Album.Name,

			// artist related field
			ArtistName: artistsName,

			// track related field
			Explicit: item.Explicit,
			ID:       item.ID,
			Name:     item.Name,
			IsLiked:  mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Items:  items,
		Total:  data.Tracks.Total,
	}
}
