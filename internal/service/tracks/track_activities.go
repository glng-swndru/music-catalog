package tracks

import (
	"context"
	"fmt"

	"github.com/glng-swndru/music-catalog/internal/models/trackactivities"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userID uint, request trackactivities.TrackActivityRequest) error {
	activity, err := s.trackActivityRepo.Get(ctx, userID, request.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get record from database")
		return err
	}
	if err == gorm.ErrRecordNotFound || activity == nil {
		// create record activity

		err = s.trackActivityRepo.Create(ctx, trackactivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil {
			log.Error().Err(err).Msg("error create record to database")
			return err
		}
		return nil
	}
	activity.IsLiked = request.IsLiked
	err = s.trackActivityRepo.Update(ctx, *activity)

	if err != nil {
		log.Error().Err(err).Msg("error update record to database")
		return err
	}
	return nil
}
