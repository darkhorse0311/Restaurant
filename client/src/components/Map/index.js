import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature } from "react-mapbox-gl";

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token
});

class Map extends Component {
  markerClick = async (place, coord) => {
    const { setCenter, getItems } = this.props;
    setCenter(coord);
    getItems(place)
  };

  render() {
    const { center, zoom, places, mapStyle } = this.props;
    const flyToOptions = { speed: 0.8 };
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
      </Mapbox>
    );
  }
}

export default Map;
