import React, { Component } from "react";
import ReactMapBoxGl, { Layer, Feature, Image } from "react-mapbox-gl";

const token = process.env.REACT_APP_MAP_BOX_KEY;
export const Mapbox = ReactMapBoxGl({
  minZoom: 10,
  accessToken: token,
});


class Map extends Component {

  state = {
    imagesLoaded: 0,
  }

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

  imageLoaded = () => {
    this.setState({imagesLoaded: this.state.imagesLoaded + 1})
  }

  renderFeature = (place) => {
    const { longitude, latitude } = place.coordinates;
    const coord = [longitude, latitude];
    return (
      <Feature
        key={`feature-${place.id}`}
        coordinates={coord}
        onClick={() => {
          this.markerClick(place, coord)
        }}
      />
    );
  }

  renderLayer = (layer) => {
    const { coords, icon, id, logo } = layer;

    // let img = new Image();
    // img.src = logo;
    // img.alt = icon;
    // img.width = 60;
    // img.height = 60;
    return (
      <React.Fragment key={id}>
        <Image 
          id={icon} 
          url={logo}
          onLoaded={() => this.imageLoaded()}
        />
        <Layer
          type="symbol"
          id={id}
          layout={{"icon-image": icon}}
          style={{zIndex: 5}}
          // images={{imageKey: icon, image: img }}
        >
          {coords.map(place => this.renderFeature(place))}
        </Layer>
      </React.Fragment>
    )
  }

  filterLayers = () => {
    const { 
      allBusinesses,
      locations,
    } = this.props;

    let layers = [];
    allBusinesses.forEach((bus, i) => {
      let loc = []
      locations.forEach(curb => {
        if (curb.r_id === bus.id) {
          loc.push(curb)
        }
      })

      if (loc.length) {
        layers.push({
          coords: loc,
          icon: `place-${bus.id}`,
          id: `bus-${bus.id}`,
          logo: bus.logo,
        })
      }
    })

    return layers;
  }

  render() {
    const { 
      center, 
      zoom,
      locations, 
      setCenter,
      allBusinesses,
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
        {
          allBusinesses.length 
            && locations.length
              && this.filterLayers().map(layer => this.renderLayer(layer))
        }
      </Mapbox>
    
    ): null;
  }
}

export default Map;
