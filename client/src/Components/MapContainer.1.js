import React, { Component } from "react";
import { Map, Marker, GoogleApiWrapper } from "google-maps-react";
import { connect } from 'react-redux';
import axios from 'axios';
import { updateSelected } from '../redux/actions/MapActions';


const key = process.env.REACT_APP_GOOGLE_MAPS_API_KEY;

class MapContainer extends Component {
    state = {
        lat: 40.7536,
        lng: -73.9832,
        // lat: 40.703486,
        // lng: -73.808929,
        restaurants: [],
        selectedPlace: {},
    }

    fetchPlaces = async (mapProps, map) => {
        const { google } = mapProps;
        const mapBounds = map.getBounds();
        const bounds = new google.maps.LatLngBounds(mapBounds.getSouthWest(), mapBounds.getNorthEast())
        
        const names = await axios.get('http://localhost:9001/api/res/names');
        const namesString = names.data.map(obj => obj.name).join(' OR ');
        console.log('namesString', namesString);
        // const namesString = 'Mcdonald OR Wendy';

        var request = {
            keyword: namesString,
            name: namesString,
            bounds: bounds,
            // fields: ['formatted_address', 'geometry', 'icon', 'id', 'name', 'permanently_closed', 'photos', 'place_id', 'plus_code', 'types']
        };

        const service = new google.maps.places.PlacesService(map);

        service.nearbySearch(request, (res, status) => {
            console.log('status', status);
            console.log('res', res);
            if (status === google.maps.places.PlacesServiceStatus.OK) {
                this.setState(previusState => {
                    return {
                        restaurants: res
                        // restaurants: [...previusState.restaurants, ...res]
                    }
                })
            }
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
        console.log('place', place.name);
        this.setState({selectedPlace: place})
        this.props.updateSelected(place)
    }

    render() {
        return (
        <Map 
            style={{
                width: '600px', height: '350px', 
                marginLeft: '-600px', 
                marginTop: '-175px'
            }}
            mapTypeControl={false}
            streetViewControl={false}
            rotateControl={false}
            fullscreenControl={false}
            zoom={14}
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

export default connect(null, { updateSelected })(GoogleApiWrapper({apiKey: key})(MapContainer));
