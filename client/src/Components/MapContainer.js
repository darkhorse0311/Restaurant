import React, { Component } from 'react';
import axios from 'axios';
import ReactMapBoxGl, { Layer, Feature, Popup } from 'react-mapbox-gl';

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
    minZoom: 10,
    accessToken: token
  });

const mapStyle = {
    flex: 1,
    width: '100vw',
    height: '100vh',
    position: 'fixed',
    top: 0,
    left: 0
};

// mapbox://styles/mapbox/satellite-v9
// mapbox://styles/mapbox/dark-v9

class MapContainer extends Component {

    state = {
        fitBounds: undefined,
        center: [-74.0060, 40.7128],
        zoom: [15],
        station: undefined,
        stations: {}
    }

    componentDidMount() {
        this.getPlaces();
    }

    getPlaces = async () => {
        
        const yelpApi = "https://api.yelp.com/v3/graphql"


        
    }

    render() {
        const { center, zoom } = this.state; 
        return (
            <Mapbox 
                // eslint-disable-next-line react/style-prop-object
                style="mapbox://styles/mapbox/dark-v9"
                containerStyle={mapStyle}
                center={center}
                zoom={zoom}
            />
        );
    }
}

export default MapContainer;
