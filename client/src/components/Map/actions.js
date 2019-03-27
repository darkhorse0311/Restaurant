import axios from 'axios';
export const SET_LAT_LON = 'set_lan_lon';
export const SET_LOCATIONS = 'set_locations';
export const SET_CENTER = 'set_center';
export const URL = 'https://carbtographer.herokuapp.com/';

export const setCordinates = (lat, lon) => {
    return {
        type: SET_LAT_LON,
        payload: {lat, lon}
    }
}

export const getLocations = (lat, lon) => async dispatch => {
    const locations = await axios.get(`${URL}/locations/${lat}/${lon}`);
    dispatch({
        type: SET_LOCATIONS,
        payload: locations.data
    })
}

export const setCenter = (coords) => {
    return {
        type: SET_CENTER,
        payload: coords
    }
}