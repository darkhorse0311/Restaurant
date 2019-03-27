import { SET_BUSINESS, SET_ITEMS } from '../types'

const defaultState = {
    id: null,
    name: "",
    coordinates: {},
    photos: [],
    distance: null,
    r_id: null,
    items: [],
};

export default (state = defaultState, action) => {
    const { payload, type } = action;
    switch(type) {
        case SET_BUSINESS: 
            return {
                ...state, 
                ...payload
            };
        case SET_ITEMS:
            return {
                ...state, 
                items: payload
            };
        default: return {...state};
    }
}
