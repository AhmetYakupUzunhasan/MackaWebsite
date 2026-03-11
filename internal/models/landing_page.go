package models

type LandingPage struct {
	Title     string  `json:"title" binding:"required"`
	SubTitle  string  `json:"subtitle" binding:"required"`
	Text      string  `json:"text" binding:"required"`
	ImageLink *string `json:"image_link,omitempty"`
}

type Blog struct {
	Title     string  `json:"title" binding:"required"`
	SubTitle  string  `json:"subtitle" binding:"required"`
	Text      string  `json:"text" binding:"required"`
	ImageLink *string `json:"image_link,omitempty"`
}
