import axios from 'axios';
import { SET_BUSINESS, SET_ITEMS, URL } from '../types'

export const getItems = place => async dispatch => {
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