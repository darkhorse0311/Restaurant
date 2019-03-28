import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature } from "react-mapbox-gl";

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token
});

class Map extends Component {

  componentDidMount() {
    // this.getCurrentCoord();
  }

  getCurrentCoord = () => {
    const { setCenter, getLocations } = this.props;
    if ("geolocation" in navigator) {
      navigator.geolocation.getCurrentPosition(({coords}) => {
        const { longitude, latitude } = coords;
        console.log(coords)
        getLocations(longitude, latitude)
        setCenter([longitude, latitude]);
      })
    } else {
        console.log("error")
    }
  }


  markerClick = (place, coord) => {
    const { setCenter, getItems, setShowModal, name } = this.props;
    setCenter(coord);
    if (place.name !== name) {
      getItems(place);
    }
    setShowModal(true)
  };

  render() {
    const { center, zoom, mapStyle, locations, setCenter } = this.props;
    const flyToOptions = { speed: 0.8 };
    console.log("center: ", center)
    return center.length === 2 ? (
      <Mapbox
        // eslint-disable-next-line react/style-prop-object
        style="mapbox://styles/mapbox/dark-v9"
        containerStyle={mapStyle}
        center={center}
        zoom={zoom}
        flyToOptions={flyToOptions}
        onMoveEnd={({transform}) => {
          const { lng, lat } = transform._center;
          setCenter([lng, lat]);
        }}
      >
        <Layer
          type="symbol"
          id="marker"
          layout={{ "icon-image": "marker-15" }}
          style={{
            zIndex: 5
          }}
        >
          {locations.map((place, i) => {
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
    ): null;
  }
}

export default Map;
