import { SET_BUSINESS, SET_ITEMS, SHOW_MODAL } from './actions'

const defaultState = {
    id: null,
    name: "",
    coordinates: {},
    photos: [],
    distance: null,
    r_id: null,
    items: [],
    showModal: false
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
        case SHOW_MODAL:
            return {
                ...state,
                showModal: payload
            }
        default: return {...state};
    }
}
