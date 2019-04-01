import { SET_LOCATIONS, SET_CENTER, LOADING, SET_PERMISSION } from './actions'

export const defaultState = {
    locations: [],
    center: [],
    zoom: [14],
    permission: false,
    loading: false,
}

export default (state = defaultState, action) => {
    const { payload, type } = action;
    switch(type) {
        case SET_LOCATIONS:
            return {
                ...state,
                locations: payload,
            };
        case SET_CENTER:
            return {
                ...state,
                center: payload
            };
        
        case LOADING:
            return {
                ...state,
                loading: payload
            }

        case SET_PERMISSION:
            return {
                ...state,
                permission: payload
            }
        default: return {...state};
    }
} 