import { SET_LAT_LON, SET_LOCATIONS, SET_CENTER, LOADING } from './actions'
import mockPlaces from '../shared/mockPlaces';

const defaultState = {
    locations: mockPlaces,
    center: [-74.0060, 40.7128],
    // locations: [],
    // center: [],
    zoom: [14],
    mapStyle: {
        flex: 1,
        width: "100vw",
        height: "100vh",
        position: "fixed",
        top: 0,
        left: 0
    },
    loading: false,
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
        
        case LOADING:
            return {
                ...state,
                loading: payload
            }
        default: return {...state};
    }
} 