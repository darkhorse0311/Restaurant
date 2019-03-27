import { SET_LAT_LON, SET_LOCATIONS } from '../actions/MapActions'

const defaultState = {
    lat: null,
    lon: null,
    locations: [],
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
        default: return {...state};
    }
} 