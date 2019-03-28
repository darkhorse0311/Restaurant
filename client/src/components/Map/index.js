import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature } from "react-mapbox-gl";
import Items from '../Items/container';

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token
});

class Map extends Component {
  markerClick = async (place, coord) => {
    const { setCenter, getItems, setShowModal } = this.props;
    setCenter(coord);
    getItems(place);
    setShowModal(true)
  };

  render() {
    const { center, zoom, places, mapStyle, showModal } = this.props;
    const flyToOptions = { speed: 0.8 };
    return (
      <>
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
        {
          showModal ? <Items/> : null
        }
      </>
    );
  }
}

export default Map;
