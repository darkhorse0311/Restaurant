import React from 'react';
import { shallow } from 'enzyme';
import Map, { Mapbox } from '../../components/Map/index';
import { Feature } from "react-mapbox-gl";

const props = {
    locations: [
      {
        "id": "p24aXNna282LwdSk0n-4UQ",
        "name": "Subway",
        "coordinates": {
          "latitude": 40.71289,
          "longitude": -74.00758
        },
        "photos": [
          "https://s3-media2.fl.yelpcdn.com/bphoto/y5kSjjmza9da5Y4lGyNHxg/o.jpg"
        ],
        "distance": 154.94281540312298,
        "r_id": 1
      },
      {
        "id": "jPIZ3FR5LNcwPuUHi2Fe4g",
        "name": "McDonald's",
        "coordinates": {
          "latitude": 40.70944,
          "longitude": -74.01012
        },
        "photos": [
          "https://s3-media1.fl.yelpcdn.com/bphoto/uCgd-65G_jhB9PiGxY8uEA/o.jpg"
        ],
        "distance": 507.50127180505456,
        "r_id": 4
      },
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



    it('two features found', () => {
        const wrapper = shallow(<Map {...props}/>);
        wrapper.find(Feature)
        expect(wrapper.find(Feature).length).toEqual(2);
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

});
