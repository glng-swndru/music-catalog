package spotify

type SearchResponse struct {
	Limit  int                  `json:"limit"`
	Offset int                  `json:"offset"`
	Items  []SpotifyTrackObject `json:"items"`
	Total  int                  `json:"total"`
}

type SpotifyTrackObject struct {
	// album related field
	AlbumType       string   `json:"albumType"`
	AlbumTotalTrack int      `json:"totaTrack"`
	AlbumImagesURL  []string `json:"albumImagesURL"`
	AlbumName       string   `json:"albumName"`

	// artist related field
	ArtistName []string `json:"artistName"`

	// track related field
	Explicit bool   `json:"explicit"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsLiked  *bool  `json:"isLiked"`
}
