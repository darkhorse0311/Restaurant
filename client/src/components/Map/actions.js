import axios from 'axios';
export const SET_LOCATIONS = 'set_locations';
export const SET_CENTER = 'set_center';
export const LOADING = 'loading';
export const SET_PERMISSION = 'set_permission';
export const SET_BUSINESSES = 'set_businesses';
export const SET_COMPACT = 'set_compact';


const url = process.env.REACT_APP_BACKEND_URL;

export const getLocations = (lon, lat, loc = []) => async dispatch => {
    dispatch({ type: LOADING, payload: true })
    const locations = await axios.get(`${url}/locations/${lat}/${lon}`);

    let check = {};

    let filterdLoc = [];

    loc.forEach(l => {
        check[l.id] = true;
        filterdLoc.push(l)
    })

    locations.data.forEach(l => {
        if (!(l.id in check)) {
            check[l.id] = true;
            filterdLoc.push(l);
        }
    })

    dispatch({
        type: SET_LOCATIONS,
        payload: filterdLoc
    })
    dispatch({ type: LOADING, payload: false })
}

export const getAllBusinesses = () => async dispatch => {
    dispatch({ type: LOADING, payload: true })
    const names = await axios.get(`${url}/names`);
    dispatch({
        type: SET_BUSINESSES,
        payload: names.data
    })
    dispatch({ type: LOADING, payload: false })
}

export const setCenter = (coords) => ({type: SET_CENTER, payload: coords });
export const setLoading = (value) => ({ type: LOADING, payload: value });
export const setPermission = (value) => ({ type: SET_PERMISSION, payload: value});
export const setCompact = (value) => ({ type: SET_COMPACT, payload: value});
