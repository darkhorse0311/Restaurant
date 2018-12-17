import React, { Component } from "react";
import { Map, InfoWindow, Marker, GoogleApiWrapper } from "google-maps-react";
const key = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;

class MapContainer extends Component {
    state = {
        lat: 40.7644,
        lng: -73.9235,
        restaurants: [],
        selectedPlace: {},
    }

    fetchPlaces = (mapProps, map) => {
        const { lat, lng } = this.state;
        const { google } = mapProps;

        const current = new google.maps.LatLng(lat, lng);

        var request = {
            query: 'mcdonald',
            location: current,
            radius: 200,
            type: 'restaurant'
        };

        const service = new google.maps.places.PlacesService(map);
        service.textSearch(request, (res, status) => {
            this.setState({restaurants: res})
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
        this.setState({selectedPlace: place})
    }

    render() {
        return (
        <Map 
            style={{width: '500px', height: '300px', marginLeft: '-250px', marginTop: '-150px' }}
            google={this.props.google} 
            zoom={15}
            initialCenter={
                {
                    lat: this.state.lat,
                    lng: this.state.lng
                }
            }
            onReady={this.fetchPlaces}
            mapTypeControl={false}
            streetViewControl={false}
            rotateControl={false}
            fullscreenControl={false}
            
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

            <InfoWindow onClose={this.onInfoWindowClose}>
                <div>
                <h1>{this.state.selectedPlace.name}</h1>
                </div>
            </InfoWindow>
        </Map>
        );
    }
}

export default GoogleApiWrapper({
  apiKey: key
})(MapContainer);
