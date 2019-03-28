import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature } from "react-mapbox-gl";

const token = process.env.REACT_APP_MAP_BOX_KEY;
const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token,
});
class Map extends Component {

  componentDidMount() {
    const { setLoading } = this.props;
    setLoading(true)
    this.getCurrentCoord();
  }

  getCurrentCoord = () => {
    const { setCenter, getLocations, setPermission } = this.props;

    navigator.permissions.query({name:'geolocation'}).then(result => {
      const { state } = result;
      if (state === "denied") {
          setCenter([-74.0060, 40.7128]);
          setPermission(false)
      } else {
        navigator.geolocation.getCurrentPosition((res) => {
          const { longitude, latitude } = res.coords;
          getLocations(longitude, latitude)
          setCenter([longitude, latitude]);
          setPermission(true)
        })
      }
    })
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
    const { 
      center, 
      zoom, 
      mapStyle, 
      locations, 
      setCenter, 
      // setLoading,
    } = this.props;

    const flyToOptions = { speed: 0.8 };

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
        // onStyleLoad={() => setLoading(false)}
      >
        <Layer
          type="symbol"
          id="marker"
          layout={{ "icon-image": "marker-15" }}
          style={{
            zIndex: 5
          }}
        >
          {locations.length && locations.map((place, i) => {
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
