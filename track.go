package soundcloud

// Track (http://developers.soundcloud.com/docs/api/reference#tracks)
type Track struct {
	Id                  int `json:"id"`
	CreatedAt           string `json:"created_at"`
	UserId              int `json:"user_id"`
	User                struct{}
	Title               string `json:"title"`
	Permalink           string `json:"permalink"`
	PermalinkUrl        string `json:"permalink_url"`
	Uri                 string `json:"uri"`
	Sharing             string `json:"sharing"`
	EmbeddableBy        string `json:"embeddable_by"`
	PurchaseUrl         string `json:"purchase_url, omitempty"`
	ArtworkUrl          string `json:"artwork_url, omitempty"`
	Description         string `json:"description, omitempty"`
	Label               struct{}
	Duration            int `json:"duration"`
	Genre               string `json:"genre, omitempty"`
	TagList             string `json:"tag_list"`
	LabelId             int `json:"label_id, omitempty"`
	LabelName           string `json:"label_name, omitempty"`
	//Release             int `json:"release, omitempty"` // string value
	ReleaseDay          int `json:"release_day, omitempty"`
	ReleaseMonth        int `json:"release_month, omitempty"`
	ReleaseYear         int `json:"release_year, omitempty"`
	Streamable          bool `json:"streamable"`
	Downloadable        bool `json:"downloadable"`
	State               string `json:"state"`
	License             string `json:"license"`
	TrackType           string `json:"track_type"`
	WaveformUrl         string `json:"waveform_url"`
	DownloadUrl         string `json:"download_url"`
	StreamUrl           string `json:"stream_url"`
	VideoUrl            string `json:"video_url, omitempty"`
	Bpm                 int `json:"bpm, omitempty"`
	Commentable         bool `json:"commentable"`
	ISRC                string `json:"isrc, omitempty"`
	KeySignature        string `json:"key_signature, omitempty"`
	CommentCount        int `json:"comment_count"`
	DownloadCount       int `json:"download_count"`
	PlaybackCount       int `json:"playback_count"`
	FavoritingsCount    int `json:"favoritings_count"`
	OriginalFormat      string `json:"original_format"`
	OriginalContentSize int `json:"original_content_size"`
	CreatedWith         struct{}
	AssetData           []byte `json:"asset_data"`
	ArtworkData         []byte `json:"artwork_data"`
	UserFavorite        bool   `json:"user_favorite"`
	AttachmentsUri      string `json:"attachments_url, omitempty"`
}
