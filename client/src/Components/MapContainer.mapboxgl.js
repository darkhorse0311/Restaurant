import React, { Component } from 'react';
import mapboxgl from 'mapbox-gl'

const mapBoxToken = "pk.eyJ1Ijoicm5sZCIsImEiOiJjanBzdGdjODkxb3JwM3ltZm1vOWQxOGFqIn0._OcNwJ_Q7XvH6a0hEahVnQ";

mapboxgl.accessToken = mapBoxToken;

// mapbox://styles/mapbox/satellite-v9
// mapbox://styles/mapbox/dark-v9

class MapContainer extends Component {

    state = {
        lng: -74.0060,
        lat: 40.7128,
        zoom: 12
    }

    componentDidMount() {
        const { lng, lat, zoom } = this.state;
    
        const map = new mapboxgl.Map({
          container: this.mapContainer,
          style: 'mapbox://styles/mapbox/dark-v9',
          center: [lng, lat],
          zoom
        });
    
        map.on('move', () => {
          const { lng, lat } = map.getCenter();
    
          this.setState({
            lng: lng.toFixed(4),
            lat: lat.toFixed(4),
            zoom: map.getZoom().toFixed(2)
          });
        });
      }

    render() {

        return (
            <div 
                ref={el => this.mapContainer = el} 
                style={{
                    width: '100vw',
                    height: '100vh',
                    position: 'fixed',
                    top: 0,
                    left: 0
                }}
            />
        );
    }
}

export default MapContainer;
