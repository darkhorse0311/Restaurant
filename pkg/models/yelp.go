package models

// Business db struct
type Business struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Coordinates struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	Photos   []string `json:"photos"`
	Distance float64  `json:"distance"`
	RID      int      `json:"r_id"`
}

// Search db response
type Search struct {
	Total    int        `json:"total"`
	Business []Business `json:"business"`
}

// YelpResponse db response
type YelpResponse struct {
	Search Search `json:"search"`
}

// YelpQuery for business info
var YelpQuery = `query ($name: String!, $lat:Float, $lon: Float) {
	search(
		term: $name,
		latitude: $lat,
		longitude: $lon,
		radius: 500
	) {
		total
		business {
			id
			name
			coordinates {
				latitude
				longitude
			}
			photos
			distance
		}
	}
}`
