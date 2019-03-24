package routes

var yelpQuery = `query ($name: String!) {
	search(
		term: $name,
		longitude: -74.0060,
		latitude: 40.7128,
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
