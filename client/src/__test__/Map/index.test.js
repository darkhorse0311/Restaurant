import React from 'react';
import { shallow } from 'enzyme';
import Map, { Mapbox } from '../../components/Map/index';
import { Feature, Image } from "react-mapbox-gl";

const props = {
    locations: [
      {
        "id": "r4jG613xCL4ispS14rEBnA",
        "name": "Subway",
        "coordinates": {
          "latitude": 40.7163137197495,
          "longitude": -74.004714936018
        },
        "photos": [
          "https://s3-media1.fl.yelpcdn.com/bphoto/TWRxx3GueDGljuZWExE5eA/o.jpg"
        ],
        "distance": 414.39746949756113,
        "r_id": 26
      }
    ],
    allBusinesses: [
      {
        id: 26,
        name: "Subway",
        logo: "https://i.imgur.com/iPihkvD.png",
      }
    ],
    center: [-74.0060, 40.7128],
    zoom: [14],
    mapStyle: {
        flex: 1,
        width: "100vw",
        height: "100vh",
        position: "fixed",
        top: 0,
        left: 0
    },
    permission: false,
    name: "",
    getLocations: jest.fn(),
    getItems: jest.fn(),
    setCenter: jest.fn(),
    setShowModal: jest.fn(),
    setLoading: jest.fn(),
    setPermission: jest.fn(),
    getAllBusinesses: jest.fn(),
};

describe('Map Component', () => {
    it('renders without crashing', () => {
        const wrapper = shallow(<Map {...props}/>);
        expect(wrapper.exists()).toBe(true)
    });

    it('componentDidMount calls setLoading and getCurrentCoord, setDefaultValues', () => {
        const wrapper = shallow(<Map {...props}/>);
        const instance = wrapper.instance();
        const getCC = jest.spyOn(instance, 'getCurrentCoord');
        const setDV = jest.spyOn(instance, 'setDefaultValues');
        instance.componentDidMount();

        expect(props.setLoading).toHaveBeenCalled();
        expect(setDV).toHaveBeenCalled();
        expect(getCC).toHaveBeenCalled();
    });

    it('componentDidMount calls setLoading, getLocations, setCenter and setPermission when geo location is on', () => {
      const mockGeolocation = {
        getCurrentPosition: jest.fn()
        .mockImplementationOnce((success) => Promise.resolve(
            success({
                coords: {
                    longitude: -74.0060, 
                    latitude: 40.7128
                }
            })
        ))
      };
        
      global.navigator.geolocation = mockGeolocation;

      const wrapper = shallow(<Map {...props}/>);
      const instance = wrapper.instance();
      const getCC = jest.spyOn(instance, 'getCurrentCoord');
      instance.componentDidMount();

      expect(props.setLoading).toHaveBeenCalled();
      expect(props.getLocations).toHaveBeenCalled();
      expect(props.setCenter).toHaveBeenCalled();
      expect(props.setPermission).toHaveBeenCalled();
      expect(getCC).toHaveBeenCalled();
    });

    it('componentDidMount calls setDefaultValues when geoLocation is denied', () => {
      const wrapper = shallow(<Map {...props}/>);
      const instance = wrapper.instance();
      const setDV = jest.spyOn(instance, 'setDefaultValues');
      const err = {
        code: 1,
        message: "User denied Geolocation"
      };
      const mockGeolocation = {
        getCurrentPosition: jest.fn()
        .mockImplementationOnce((success, error) => error(err))
      };
      global.navigator.geolocation = mockGeolocation;
      instance.componentDidMount();
      expect(setDV).toHaveBeenCalled();
    });

    it('one feature found', () => {
        const wrapper = shallow(<Map {...props}/>);
        wrapper.find(Feature)
        expect(wrapper.find(Feature).length).toEqual(1);
    });

    it('feature clicked calls markerClick, setCenter and setShowModal', () => {
        const wrapper = shallow(<Map {...props}/>);
        wrapper.find(Feature).at(0).simulate('click');

        expect(props.setCenter).toHaveBeenCalled();
        expect(props.setShowModal).toHaveBeenCalled();
    });

    it('calls setCenter on onMoveEnd event from Mapbox component', () => {
      const wrapper = shallow(<Map {...props}/>);
      const map = wrapper.find(Mapbox);
      expect(map.length).toEqual(1);

      map.getElement().props.onMoveEnd();
      expect(props.setCenter).toHaveBeenCalled();
    })

    it('calls onLoad event from Image and increments imagesLoaded', () => {
      const wrapper = shallow(<Map {...props}/>);
      const img = wrapper.find(Image);
      expect(Image.length).toEqual(1);

      const instance = wrapper.instance();
      const imgFunc = jest.spyOn(instance, 'imageLoaded');

      img.getElement().props.onLoaded();
      expect(imgFunc).toHaveBeenCalled();
    })
});
