import React, { Component } from "react";
import axios from "axios";
import ReactMapBoxGl, { Layer, Feature, Popup } from "react-mapbox-gl";
import StyledPopup from "./StyledPopup";
import BottomPopup from './BottomPopup';

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
    places: [
      {
        id: "r4jG613xCL4ispS14rEBnA",
        name: "Subway",
        coordinates: {
          latitude: 40.7163137197495,
          longitude: -74.004714936018
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/Gvi5nfplLSzbcSFYKivFVg/o.jpg"
        ],
        distance: null
      },
      {
        id: "p24aXNna282LwdSk0n-4UQ",
        name: "Subway",
        coordinates: {
          latitude: 40.71289,
          longitude: -74.00758
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/y5kSjjmza9da5Y4lGyNHxg/o.jpg"
        ],
        distance: null
      },
      {
        id: "pymPEcd_X1A1pvXHQeg0_w",
        name: "Wendy's",
        coordinates: {
          latitude: 40.71013,
          longitude: -74.00816
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/v30Yqk15dCCaeg8RJpOU8A/o.jpg"
        ],
        distance: null
      },
      {
        id: "nT3rCBXEeORDUEj0rlWDUA",
        name: "Subway",
        coordinates: {
          latitude: 40.71596,
          longitude: -74.00996
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/fAOvbeERGjiW1vKr2vFf3A/o.jpg"
        ],
        distance: null
      },
      {
        id: "mnEMipzEekqLWIogjAB_zA",
        name: "Subway",
        coordinates: {
          latitude: 40.707682,
          longitude: -74.005945
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/6_6U9kCtr-Kjs8dYrNVZ6A/o.jpg"
        ],
        distance: null
      },
      {
        id: "LAo4PdkUMwRF88x5TnG3wA",
        name: "Subway",
        coordinates: {
          latitude: 40.7151328772306,
          longitude: -74.0077042579651
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/K81z15CGfnNMqRO0onPB-Q/o.jpg"
        ],
        distance: null
      },
      {
        id: "xVoy_j2M7KulKHs3NZLqLg",
        name: "Subway",
        coordinates: {
          latitude: 40.70933,
          longitude: -74.00903
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/XBrYo0HowL0frL9zaRbJDA/o.jpg"
        ],
        distance: null
      },
      {
        id: "uL9Pz4XTLmv9Rial9BKggw",
        name: "Subway",
        coordinates: {
          latitude: 40.71012,
          longitude: -74.0054399
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/SRqV-Xr4vF2IpbeJvglHYw/o.jpg"
        ],
        distance: null
      },
      {
        id: "apD1JSpVX_q4PJh90xJPoQ",
        name: "McDonald's",
        coordinates: {
          latitude: 40.715771,
          longitude: -74.005234
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/cQ6rCEKzwfXryupYe5VS8g/o.jpg"
        ],
        distance: null
      },
      {
        id: "jPIZ3FR5LNcwPuUHi2Fe4g",
        name: "McDonald's",
        coordinates: {
          latitude: 40.70944,
          longitude: -74.01012
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/nt-MU_AkFIC299Xsr82OyA/o.jpg"
        ],
        distance: null
      },
      {
        id: "hT229QtA--o0F591IUE7SQ",
        name: "McDonald's",
        coordinates: {
          latitude: 40.708279,
          longitude: -74.004868
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/t5Mpt-eGJFJDdvrr1y9Jpg/o.jpg"
        ],
        distance: null
      },
      {
        id: "in7QM4vNUTHVZIi7Outetg",
        name: "McDonald's",
        coordinates: {
          latitude: 40.7163213865248,
          longitude: -74.0106988300204
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/2Yr8BIP_Q-LEwhtKPSna3g/o.jpg"
        ],
        distance: null
      },
      {
        id: "YV0VaJgV2zotzI7XNUaMCQ",
        name: "McDonald's",
        coordinates: {
          latitude: 40.71852,
          longitude: -74.00115
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/U6azm8z8Og9Jt19Dm38rbg/o.jpg"
        ],
        distance: null
      },
      {
        id: "2JAoFi1nu3xlrdNzW5FLMQ",
        name: "Quiznos",
        coordinates: {
          latitude: 40.7143517,
          longitude: -74.0070343
        },
        photos: ["https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"],
        distance: null
      },
      {
        id: "pTv4RDll3N_CGG3BwtEctg",
        name: "Starbucks",
        coordinates: {
          latitude: 40.711619,
          longitude: -74.006730001524
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/4FTYS3muFZKyO96BrY5oKw/o.jpg"
        ],
        distance: null
      },
      {
        id: "4jWEL2gMhfwfzej0rfgJiQ",
        name: "Starbucks",
        coordinates: {
          latitude: 40.712214,
          longitude: -74.008202
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/T4GTFPs4plbOevCLS5J32g/o.jpg"
        ],
        distance: null
      },
      {
        id: "V05D4jAQb3rT3WqEBGq-og",
        name: "Starbucks",
        coordinates: {
          latitude: 40.714843,
          longitude: -74.005956
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/1vbi7VxH8sjE8C79eLdh3w/o.jpg"
        ],
        distance: null
      },
      {
        id: "8uTTEIDkYuEXywiWcP5w-w",
        name: "Starbucks",
        coordinates: {
          latitude: 40.713782,
          longitude: -74.009077
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/z2nkikSCdvrqlIaY3dv3Gg/o.jpg"
        ],
        distance: null
      },
      {
        id: "12xF41OEVDoqsSwnDTEH1g",
        name: "Starbucks",
        coordinates: {
          latitude: 40.710792,
          longitude: -74.008533
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/ZgZ1A3T0Ulf7VIT7QkIKtQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "j6nTfbYavGVsJKxz6bte6w",
        name: "Starbucks",
        coordinates: {
          latitude: 40.710194,
          longitude: -74.007894
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/e21NmNVO8CSTjmIqHhnrHg/o.jpg"
        ],
        distance: null
      },
      {
        id: "rveP6oXixE2a2sKwMqS-ww",
        name: "Starbucks",
        coordinates: {
          latitude: 40.715532,
          longitude: -74.009003
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/DTTJRPGnlyvMPH-jNVUSnw/o.jpg"
        ],
        distance: null
      },
      {
        id: "XzAZVxmGPcysK8hYdcAJKA",
        name: "Starbucks",
        coordinates: {
          latitude: 40.715754,
          longitude: -74.003125
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/zzyiNhtV-LYrBlObVjDM9g/o.jpg"
        ],
        distance: null
      },
      {
        id: "CszSAlKolUeJ0-IaRDL0YQ",
        name: "Starbucks",
        coordinates: {
          latitude: 40.71091,
          longitude: -74.010363
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/0FUeyogFE3caeNMELwyUCg/o.jpg"
        ],
        distance: null
      },
      {
        id: "hh2T5aLJsSjpHvPRcQxIJw",
        name: "Starbucks",
        coordinates: {
          latitude: 40.712117,
          longitude: -74.01087
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/2iB1RbxLAhUkBOlVtFQzVA/o.jpg"
        ],
        distance: null
      },
      {
        id: "60agfQbky4cX8BEApyltIA",
        name: "Starbucks",
        coordinates: {
          latitude: 40.711018,
          longitude: -74.000916
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/biVl1w3q-gf4_CELfJJCQQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "OuNYQaqEJjkBHRdnC-4HOw",
        name: "Starbucks",
        coordinates: {
          latitude: 40.708413,
          longitude: -74.007387
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/xZM4c99d2aLoZcG0CdnjFg/o.jpg"
        ],
        distance: null
      },
      {
        id: "p8Fygzdm1uONUTu4q-SfJA",
        name: "Starbucks",
        coordinates: {
          latitude: 40.709590668884,
          longitude: -74.010550737875
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/6QYI1E8yqtvguCxar6gj8A/o.jpg"
        ],
        distance: null
      },
      {
        id: "9SA6MazvA6lXqTpM-ZlYIA",
        name: "Starbucks",
        coordinates: {
          latitude: 40.708771,
          longitude: -74.009313
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/gC70m0IL4UEGb_cAxLrLXw/o.jpg"
        ],
        distance: null
      },
      {
        id: "OQEU7KPvgDX9kX2xlHGm5g",
        name: "Starbucks",
        coordinates: {
          latitude: 40.710372,
          longitude: -74.011793
        },
        photos: [
          "https://s3-media4.fl.yelpcdn.com/bphoto/BmELR_JNxYfCPfELix3itQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "wE3_AwZoZ0ZXbngBH1xRrg",
        name: "Starbucks",
        coordinates: {
          latitude: 40.718247,
          longitude: -74.007305
        },
        photos: [
          "https://s3-media4.fl.yelpcdn.com/bphoto/bX2_feJRAbOaeHUil025MQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "FrigBgxfxMxwPJNq-hAr0Q",
        name: "Starbucks",
        coordinates: {
          latitude: 40.708758,
          longitude: -74.011055
        },
        photos: [
          "https://s3-media4.fl.yelpcdn.com/bphoto/IJ_KmH7WIF--a9V4a-I48w/o.jpg"
        ],
        distance: null
      },
      {
        id: "VGMdcAn34GeasB4arA8XUA",
        name: "Starbucks",
        coordinates: {
          latitude: 40.707122,
          longitude: -74.004971
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/Z2pqqE7e6_K4CLiWxT0ZGQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "_H9AnQkr5VtOiykNnxcQlQ",
        name: "Starbucks",
        coordinates: {
          latitude: 40.706487,
          longitude: -74.006685
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/ekACuqH5V_ioAnzPfEMwhQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "_rgXsb2X_NVwbt34qyHa4A",
        name: "Starbucks",
        coordinates: {
          latitude: 40.707437,
          longitude: -74.010976
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/LP1uGWf9bloEzVqutmwIWw/o.jpg"
        ],
        distance: null
      },
      {
        id: "6XAJEUfXDkWCBxAJTQSZxQ",
        name: "Subway",
        coordinates: {
          latitude: 40.7164726406336,
          longitude: -74.0048296004534
        },
        photos: ["https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"],
        distance: null
      },
      {
        id: "JYZIJRhWc6KA-SWuFJPAXw",
        name: "Subway",
        coordinates: {
          latitude: 40.70871,
          longitude: -74.01164
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/LNIiBKiwF0W_wP5DUeVXYA/o.jpg"
        ],
        distance: null
      },
      {
        id: "KMt3QFk-Bd5FMxxMW2mpnQ",
        name: "Subway",
        coordinates: {
          latitude: 40.719177,
          longitude: -74.004311
        },
        photos: [
          "https://s3-media1.fl.yelpcdn.com/bphoto/Fydkht8NaqZR9Gt2IhR7dg/o.jpg"
        ],
        distance: null
      },
      {
        id: "-3Gj4IrIu5Dn0LIZxidwkA",
        name: "Subway",
        coordinates: {
          latitude: 40.7181881373376,
          longitude: -74.000615833625
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/QQIxYB3P43IXAugDw1ETwg/o.jpg"
        ],
        distance: null
      },
      {
        id: "6eId3eiOMQf5iEhypNBV0g",
        name: "Subway",
        coordinates: {
          latitude: 40.70854,
          longitude: -74.01367
        },
        photos: ["https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"],
        distance: null
      },
      {
        id: "Mqm-66RFAvxfpL2qIyDwXw",
        name: "Subway",
        coordinates: {
          latitude: 40.71646,
          longitude: -73.99761
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/wMn0gcGhENEuQAL34_dmnQ/o.jpg"
        ],
        distance: null
      },
      {
        id: "RGK1ASygW8Q4JeJJ-Az9Xw",
        name: "Subway",
        coordinates: {
          latitude: 40.7062784,
          longitude: -74.011724
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/W7H0GugJX_GeGRuN_Gndnw/o.jpg"
        ],
        distance: null
      },
      {
        id: "XfugTKWi1sI_5Ai1Dr8t9A",
        name: "Subway",
        coordinates: {
          latitude: 40.70634,
          longitude: -74.0038299
        },
        photos: ["https://s3-media3.fl.yelpcdn.com/bphoto/None/o.jpg"],
        distance: null
      },
      {
        id: "OV66AsjGZZMvugzrxu4aig",
        name: "Subway",
        coordinates: {
          latitude: 40.70458,
          longitude: -74.01048
        },
        photos: [
          "https://s3-media3.fl.yelpcdn.com/bphoto/Iig_LbXdc0lrSM0fQqSN2Q/o.jpg"
        ],
        distance: null
      },
      {
        id: "z3n_hpAFLyDHfMrJKV6fkA",
        name: "Subway",
        coordinates: {
          latitude: 40.7052841186523,
          longitude: -74.0142669677734
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/TytLd4qUIKW8a_77vQD24w/o.jpg"
        ],
        distance: null
      },
      {
        id: "WdfaJmWIiK80aQd2o2RR-Q",
        name: "Subway",
        coordinates: {
          latitude: 40.703255,
          longitude: -74.01107
        },
        photos: [
          "https://s3-media4.fl.yelpcdn.com/bphoto/eotbiv9b7Pk0sf-jip_e4A/o.jpg"
        ],
        distance: null
      },
      {
        id: "IWZqUl46ogITQyDz3UAXPQ",
        name: "Subway",
        coordinates: {
          latitude: 40.7023139,
          longitude: -74.0114968
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/DajepoLqEE50EsInb7BPFA/o.jpg"
        ],
        distance: null
      },
      {
        id: "FpYVG2jW7SmLhsHBNEGDtg",
        name: "Wendy's",
        coordinates: {
          latitude: 40.701009,
          longitude: -74.0130246
        },
        photos: [
          "https://s3-media2.fl.yelpcdn.com/bphoto/sjZoHjq9YFlR4L-Uz1ekrA/o.jpg"
        ],
        distance: null
      }
    ]
  };

  componentDidMount() {
    this.getPlaces();
  }

  getPlaces = async () => {};

  clearPlace = () => {
    console.log("here");
    this.setState({ place: null, zoom: [14] });
  };

  markerClick = async (place, coord) => {
    const items = await axios.get(
      `http://localhost:9001/api/res/items?name=${place.name}`
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
          // <Popup key={place.id} coordinates={place.coord} anchor="bottom">
          //   <StyledPopup place={place} clearPlace={this.clearPlace} />
          // </Popup>
          <BottomPopup place={place} clearPlace={this.clearPlace} />
        )}
      </Mapbox>
    );
  }
}

export default MapContainer;
