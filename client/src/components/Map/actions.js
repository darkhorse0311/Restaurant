import axios from 'axios';
export const SET_LAT_LON = 'set_lan_lon';
export const SET_LOCATIONS = 'set_locations';
export const SET_CENTER = 'set_center';
export const LOADING = 'loading';

export const URL = 'https://carbtographer.herokuapp.com/';


export const getLocations = (lon, lat) => async dispatch => {
    dispatch({ type: LOADING, payload: true })
    const locations = await axios.get(`${URL}/locations/${lat}/${lon}`);
    dispatch({
        type: SET_LOCATIONS,
        payload: locations.data
    })
    dispatch({ type: LOADING, payload: false })
}

export const setCenter = (coords) => {
    return {
        type: SET_CENTER,
        payload: coords
    }
}


export const setLoading = (value) => ({ type: LOADING, payload: value });
