import React, { Component } from 'react';

import ReactMapBoxGl, { Layer, Feature, Popup } from 'react-mapbox-gl';

const token= "pk.eyJ1Ijoicm5sZCIsImEiOiJjanBzdGdjODkxb3JwM3ltZm1vOWQxOGFqIn0._OcNwJ_Q7XvH6a0hEahVnQ";

const Mapbox = ReactMapBoxGl({
    minZoom: 8,
    maxZoom: 15,
    accessToken: token
  });

const mapStyle = {
flex: 1,
width: '100vw',
height: '100vh'
};

// mapbox://styles/mapbox/satellite-v9
// mapbox://styles/mapbox/dark-v9

class MapContainer extends Component {

    state = {
        fitBounds: undefined,
        center: [-74.0060, 40.7128],
        zoom: [12],
        station: undefined,
        stations: {}
    }

    render() {

        return (
            <Mapbox 
                // eslint-disable-next-line react/style-prop-object
                style="mapbox://styles/mapbox/dark-v9"
                containerStyle={mapStyle}

            />
        );
    }
}

export default MapContainer;
