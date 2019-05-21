package models

// FusionResponse from yelps fusion api
type FusionResponse struct {
	Total      int        `json:"total"`
	Businesses []Business `json:"businesses"`
	Region     struct {
		Center struct {
			Latitude  float32 `json:"latitude"`
			Longitude float32 `json:"longitude"`
		} `json:"center"`
	} `json:"region"`
}

// BusinessResponse struct
type BusinessResponse struct {
	Rating     int    `json:"rating"`
	Price      string `json:"price"`
	Phone      string `json:"phone"`
	ID         string `json:"id"`
	Alias      string `json:"alias"`
	IsClosed   bool   `json:"is_closed"`
	Categories []struct {
		Alias string `json:"alias"`
		Title string `json:"title"`
	} `json:"categories"`
	ReviewCount int    `json:"review_count"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Coordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	} `json:"coordinates"`
	ImageURL string `json:"image_url"`
	Location struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		Address2 string `json:"address2"`
		Address3 string `json:"address3"`
		State    string `json:"state"`
		Address1 string `json:"address1"`
		ZipCode  string `json:"zip_code"`
	} `json:"location"`
	Distance     float32  `json:"distance"`
	Transactions []string `json:"transactions"`
	RID          int      `json:"r_id"`
}
