import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature, Image } from "react-mapbox-gl";

const token = process.env.REACT_APP_MAP_BOX_KEY;
export const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token,
});


class Map extends Component {

  componentDidMount() {
    const { setLoading, getAllBusinesses } = this.props;
    getAllBusinesses();
    setLoading(true);
    this.getCurrentCoord();
  }

  setDefaultValues = () => {
    const { setCenter, getLocations, setPermission } = this.props;
    setCenter([-74.0060, 40.7128]);
    getLocations(-74.0060, 40.7128)
    setPermission(false)
  }

  getCurrentCoord = () => {
    const { setCenter, getLocations, setPermission } = this.props;

    if ("geolocation" in navigator) {
      navigator.geolocation.getCurrentPosition(res => {
        const { longitude, latitude } = res.coords;
        getLocations(longitude, latitude)
        setCenter([longitude, latitude]);
        setPermission(true)
      }, err => {
        this.setDefaultValues()
      });
    } else {this.setDefaultValues()}
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
      locations, 
      setCenter,
    } = this.props;

    const flyToOptions = { speed: 0.8 };
    
    const mapStyle = {
      flex: 1,
      width: '100vw',
      height: "100vh",
      position: "fixed",
      top: 0,
      left: 0
    };

    return center.length === 2 ? (
      <Mapbox
        // eslint-disable-next-line react/style-prop-object
        style="mapbox://styles/mapbox/dark-v9"
        containerStyle={mapStyle}
        center={center}
        zoom={zoom}
        flyToOptions={flyToOptions}
        onMoveEnd={e => {
          const { transform = {} } = e || {};
          const { lng = -74.0060, lat = 40.7128 } = transform._center || {};
          setCenter([lng, lat]);
        }}
      >
        <Image id={'starbucks'} url={'https://i.imgur.com/Lt5dsl6.png'} />
        <Layer
          type="symbol"
          id="marker"
          layout={{ "icon-image": "starbucks" }}
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
