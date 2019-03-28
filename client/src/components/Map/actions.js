import axios from 'axios';
export const SET_LAT_LON = 'set_lan_lon';
export const SET_LOCATIONS = 'set_locations';
export const SET_CENTER = 'set_center';
export const LOADING = 'loading';
export const SET_PERMISSION = 'set_permission';


const url = process.env.REACT_APP_BACKEND_URL;

export const getLocations = (lon, lat) => async dispatch => {
    dispatch({ type: LOADING, payload: true })
    const locations = await axios.get(`${url}/locations/${lat}/${lon}`);
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
export const setPermission = (value) => ({ type: SET_PERMISSION, payload: value});
