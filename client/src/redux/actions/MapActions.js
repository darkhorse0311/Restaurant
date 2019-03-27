import axios from 'axios';
import { SET_LAT_LON, SET_LOCATIONS, URL } from '../actions/MapActions'

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