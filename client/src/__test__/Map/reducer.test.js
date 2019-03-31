import reducer, { defaultState } from '../../components/Map/reducer';
import { SET_LAT_LON, SET_LOCATIONS, SET_CENTER, LOADING, SET_PERMISSION } from '../../components/Map/actions'

const locations = [
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
];

describe('Map reducer', () => {
    it('should return the initial state', () => {
        expect(reducer(undefined, {})).toEqual(defaultState)
    })

    it('should handle SET_LAT_LON', () => {
        expect(
          reducer(defaultState, {
            type: SET_LAT_LON,
            payload: {
                lon: -74.0060, 
                lat: 40.7128
            }
          })
        ).toEqual({
            ...defaultState, 
            lon: -74.0060, 
            lat: 40.7128
        })
    });

    it('should handle SET_LOCATIONS', () => {
        expect(
          reducer(defaultState, {
            type: SET_LOCATIONS,
            payload: locations,
          })
        ).toEqual({
            ...defaultState, 
            locations,
        })
    });

    it('should handle SET_CENTER', () => {
        expect(
          reducer(defaultState, {
            type: SET_CENTER,
            payload: [-74.0060, 40.7128],
          })
        ).toEqual({
            ...defaultState, 
            center: [-74.0060, 40.7128],
        })
    });

    it('should handle LOADING', () => {
        expect(
          reducer(defaultState, {
            type: LOADING,
            payload: true,
          })
        ).toEqual({
            ...defaultState, 
            loading: true
        })
    });

    it('should handle SET_PERMISSION', () => {
        expect(
          reducer(defaultState, {
            type: SET_PERMISSION,
            payload: true,
          })
        ).toEqual({
            ...defaultState, 
            permission: true
        })
    });
});


