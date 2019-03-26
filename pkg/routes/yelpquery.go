package routes

var yelpQuery = `query ($name: String!, $lat:Float, $lon: Float) {
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
