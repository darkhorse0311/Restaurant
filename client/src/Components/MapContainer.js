import React, { Component } from "react";
import axios from "axios";
import ReactMapBoxGl, { Layer, Feature, Popup } from "react-mapbox-gl";
import BottomPopup from './BottomPopup';

import mockPlaces from './mockPlaces';

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token
});

const mapStyle = {
  flex: 1,
  width: "100vw",
  height: "100vh",
  position: "fixed",
  top: 0,
  left: 0
};

class MapContainer extends Component {
  state = {
    fitBounds: undefined,
    center: [-74.006, 40.7128],
    zoom: [14],
    place: null,
    places: mockPlaces
  }

  componentDidMount() {
    this.getPlaces();
  }

  getPlaces = async () => {};

  clearPlace = () => {
    this.setState({ place: null, zoom: [14] });
  };

  markerClick = async (place, coord) => {
    const items = await axios.get(
      `http://localhost:9001/items/${place.r_id}`
    );

    this.setState({
      center: coord,
      zoom: [16],
      place: { ...place, coord, items: items.data }
    });
  };

  render() {
    const { center, zoom, places, place } = this.state;
    const flyToOptions = {
      speed: 0.8
    };
    return (
      <Mapbox
        // eslint-disable-next-line react/style-prop-object
        style="mapbox://styles/mapbox/dark-v9"
        containerStyle={mapStyle}
        center={center}
        zoom={zoom}
        flyToOptions={flyToOptions}
      >
        <Layer
          type="symbol"
          id="marker"
          layout={{ "icon-image": "marker-15" }}
          style={{
            zIndex: 5
          }}
        >
          {places.map((place, i) => {
            const { longitude, latitude } = place.coordinates;
            const coord = [longitude, latitude];
            return (
              <Feature
                coordinates={coord}
                key={i}
                onClick={() => this.markerClick(place, coord)}
              />
            );
          })}
        </Layer>
        {place && (
          <BottomPopup place={place} clearPlace={this.clearPlace} />
        )}
      </Mapbox>
    );
  }
}

export default MapContainer;
