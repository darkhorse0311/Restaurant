import { SET_LAT_LON, SET_LOCATIONS, SET_CENTER } from './actions'
import mockPlaces from '../../mockPlaces';

const defaultState = {
    lat: null,
    lon: null,
    locations: [],
    center: [-74.0060, 40.7128],
    zoom: [14],
    places: mockPlaces,
    mapStyle: {
        flex: 1,
        width: "100vw",
        height: "100vh",
        position: "fixed",
        top: 0,
        left: 0
    }
}

export default (state = defaultState, action) => {
    const { payload, type } = action;
    switch(type) {
        case SET_LAT_LON:
            return {
                ...state,
                lat: payload.lat,
                lon: payload.lon,
            };
        case SET_LOCATIONS:
            return {
                ...state,
                locations: payload,
            };
        case SET_CENTER:
            return {
                ...state,
                center: payload
            };
        default: return {...state};
    }
} 