import axios from 'axios';
export const SET_BUSINESS = 'set_business';
export const SET_ITEMS = 'set_items';
export const URL = 'https://carbtographer.herokuapp.com/';

export const getItems = place => async dispatch => {
    dispatch({
        type: SET_BUSINESS,
        payload: place
    })
    const items = await axios.get(`${URL}/items/${place.r_id}`);
    dispatch({
        type: SET_ITEMS,
        payload: items.data
    })
}

export const setBusiness = business => {
    return {
        type: SET_BUSINESS,
        payload: business
    }
}