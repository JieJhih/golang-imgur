package imgur

type ImgurResponse struct {
	Data    Imgdata `json:"data"`
	Success bool    `json:"success"`
	Status  int     `json:"status"`
}

type Imgdata struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Datetime    int64    `json:"datetime"`
	Type        string   `json:"type"`
	Animated    bool     `json:"animated"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	Size        int      `json:"size"`
	Views       int      `json:"views"`
	Bandwidth   int      `json:"bandwidth"`
	Vote        string   `json:"vote"`
	Favorite    bool     `json:"favorite"`
	Nsfw        string   `json:"nsfw"`
	Section     string   `json:"section"`
	AcountURL   string   `json:"account_url"`
	AcountID    int      `json:"account_id"`
	IsAD        bool     `json:"is_ad"`
	InMostViral bool     `json:"in_most_viral"`
	Tags        []string `json:"tags"`
	ADType      int      `json:"ad_type"`
	ADUrl       string   `json:"ad_url"`
	InGallery   bool     `json:"in_gallery"`
	Deletehash  string   `json:"deletehash"`
	Name        string   `json:"name"`
	Link        string   `json:"link"`
}
