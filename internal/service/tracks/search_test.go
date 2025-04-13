package tracks

import (
	"context"
	"reflect"
	"testing"

	"github.com/glng-swndru/music-catalog/internal/models/spotify"
	spotifyRepo "github.com/glng-swndru/music-catalog/internal/repository/spotify"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_service_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSpotifyOutbound := NewMockSpotifyOutbound(mockCtrl)
	next := "https://api.spotify.com/v1/search?offset=10&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8"
	type args struct {
		query     string
		pageSize  int
		pageIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    *spotify.SearchResponse
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want: &spotify.SearchResponse{
				Limit:  10,
				Offset: 0,
				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:       "album",
						AlbumTotalTrack: 22,
						AlbumImagesURL:  []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:       "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistName:      []string{"Queen"},
						Explicit:        false,
						ID:              "3z8h0TU7ReDPLIbEnYhWZb",
						Name:            "Bohemian Rhapsody",
					},
					{
						AlbumType:       "compilation",
						AlbumTotalTrack: 17,
						AlbumImagesURL:  []string{"https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263", "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263"},
						AlbumName:       "Greatest Hits (Remastered)",
						ArtistName:      []string{"Queen"},
						Explicit:        false,
						ID:              "2OBofMJx94NryV2SK8p8Zf",
						Name:            "Bohemian Rhapsody - Remastered 2011",
					},
				},
				Total: 26,
			},
			wantErr: false,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(&spotifyRepo.SpotifySearchResponse{
					Tracks: spotifyRepo.SpotifyTracks{
						Href:   "https://api.spotify.com/v1/search?offset=0&limit=10&query=bohemian%20rhapsody&type=track&market=ID&locale=en-US,en;q%3D0.9,id;q%3D0.8",
						Limit:  10,
						Next:   &next,
						Offset: 0,
						Total:  26,
						Items: []spotifyRepo.SpotifyTrackObject{
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:  "album",
									TotalTrack: 22,
									Images: []spotifyRepo.SpotifyAlbumImage{
										{
											URL: "https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b",
										},
									},
									Name: "Bohemian Rhapsody (The Original Soundtrack)",
								},
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/3z8h0TU7ReDPLIbEnYhWZb",
								ID:       "3z8h0TU7ReDPLIbEnYhWZb",
								Name:     "Bohemian Rhapsody",
							},
							{
								Album: spotifyRepo.SpotifyAlbumObject{
									AlbumType:  "compilation",
									TotalTrack: 17,
									Images: []spotifyRepo.SpotifyAlbumImage{
										{
											URL: "https://i.scdn.co/image/ab67616d0000b273bb19d0c22d5709c9d73c8263",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00001e02bb19d0c22d5709c9d73c8263",
										},
										{
											URL: "https://i.scdn.co/image/ab67616d00004851bb19d0c22d5709c9d73c8263",
										},
									},
									Name: "Greatest Hits (Remastered)",
								},
								Artists: []spotifyRepo.SpotifyArtistObject{
									{
										Href: "https://api.spotify.com/v1/artists/1dfeR4HaWDbWqFHLkxsg1d",
										Name: "Queen",
									},
								},
								Explicit: false,
								Href:     "https://api.spotify.com/v1/tracks/2OBofMJx94NryV2SK8p8Zf",
								ID:       "2OBofMJx94NryV2SK8p8Zf",
								Name:     "Bohemian Rhapsody - Remastered 2011",
							},
						},
					},
				}, nil)
			},
		},
		{
			name: "fail",
			args: args{
				query:     "bohemian rhapsody",
				pageSize:  10,
				pageIndex: 1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mockSpotifyOutbound.EXPECT().Search(gomock.Any(), args.query, 10, 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				spotifyOutbound: mockSpotifyOutbound,
			}
			got, err := s.Search(context.Background(), tt.args.query, tt.args.pageSize, tt.args.pageIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
