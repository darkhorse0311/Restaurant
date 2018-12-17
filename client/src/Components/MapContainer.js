import React, { Component } from "react";
import { Map, Marker, GoogleApiWrapper } from "google-maps-react";
const key = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;

class MapContainer extends Component {
    state = {
        lat: 40.703486,
        lng: -73.808929,
        restaurants: [],
        selectedPlace: {},
    }

    fetchPlaces = async (mapProps, map) => {
        const { google } = mapProps;
        const mapBounds = map.getBounds();
        const bounds = new google.maps.LatLngBounds(mapBounds.getSouthWest(), mapBounds.getNorthEast())
        
        // const names = await axios.get('http://localhost:9001/api/res/names');
        // const namesString = names.data.map(obj => obj.name).join(' OR ');
        // console.log('namesString', namesString);
        const namesString = 'Mcdonald OR Wendy';

        var request = {
            bounds,
            query: namesString,
            type: 'restaurant'
        };

        const service = new google.maps.places.PlacesService(map);
        service.textSearch(request, (res, status, pagination) => {
            // console.log('res', res);
            this.setState(previusState => {
                return {
                    restaurants: [...previusState.restaurants, ...res]
                }
            })
            // if (pagination.hasNextPage) {
            //     pagination.nextPage();
            // }
        });
    }

    getLocation = () => {
        const options = {
            enableHighAccuracy: true,
            timeout: 5000,
            maximumAge: 0
        };

        const error = (err) => console.log(err);

        const success = (pos) => {
            const crd = pos.coords;
            this.setState({lat: crd.latitude, lng: crd.longitude});
        }

        navigator.geolocation.getCurrentPosition(success, error, options)
    }

    onMarkerClick = (e, place) => {
        console.log('place', place);
        this.setState({selectedPlace: place})
    }

    render() {
        return (
        <Map 
            style={{width: '500px', height: '300px', marginLeft: '-250px', marginTop: '-150px' }}
            mapTypeControl={false}
            streetViewControl={false}
            rotateControl={false}
            fullscreenControl={false}
            zoom={15}
            google={this.props.google} 
            initialCenter={
                {
                    lat: this.state.lat,
                    lng: this.state.lng
                }
            }
            // onReady={this.fetchPlaces}
            onTilesloaded={this.fetchPlaces}
        >

            {
                this.state.restaurants.map((place, index) => {
                    return <Marker
                        name={place.name}
                        position={place.geometry.location}
                        onClick={e => this.onMarkerClick(e, place)}
                        key={index}
                    /> 
                })
            }
        </Map>
        );
    }
}

export default GoogleApiWrapper({
  apiKey: key
})(MapContainer);
