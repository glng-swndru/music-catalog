package tracks

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/music-catalog/internal/models/spotify"
	"github.com/glng-swndru/music-catalog/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
		expectedBody       spotify.SearchResponse
		wantErr            bool
	}{
		{
			name:               "success",
			expectedStatusCode: 200,
			expectedBody: spotify.SearchResponse{
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
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1, uint(1)).Return(
					&spotify.SearchResponse{
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
					}, nil)
			},
		},
		{
			name:               "fail",
			expectedStatusCode: 400,
			expectedBody:       spotify.SearchResponse{},
			wantErr:            true,
			mockFn: func() {
				mockSvc.EXPECT().Search(gomock.Any(), "bohemian rhapsody", 10, 1, uint(1)).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()

			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			h.RegisterRoutes()
			w := httptest.NewRecorder()
			endpoint := `/tracks/search?query=bohemian+rhapsody&pageSize=10&pageIndex=1`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()

				response := spotify.SearchResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)

				assert.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
