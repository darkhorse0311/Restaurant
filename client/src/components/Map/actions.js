import axios from 'axios';
export const SET_LAT_LON = 'set_lan_lon';
export const SET_LOCATIONS = 'set_locations';
export const SET_CENTER = 'set_center';
export const URL = 'https://carbtographer.herokuapp.com/';


export const getLocations = (lon, lat) => async dispatch => {
    const locations = await axios.get(`${URL}/locations/${lat}/${lon}`);
    console.log(locations)
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
