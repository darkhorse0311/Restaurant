import axios from 'axios';

export const UPDATE_SELECTED = 'update_selected';

const url = 'http://localhost:9001';

export const updateSelected = place => async dispatch => {
    const items = await axios.get(`${url}/items/${place.r_id}`);
    dispatch({
        type: UPDATE_SELECTED,
        payload: {...place, items: items.data}
    })
}