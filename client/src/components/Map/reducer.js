import { SET_LOCATIONS, SET_CENTER, LOADING, SET_PERMISSION } from './actions'

export const defaultState = {
    locations: [],
    center: [],
    zoom: [14],
    mapStyle: {
        flex: 1,
        width: "100vw",
        height: "100vh",
        position: "fixed",
        top: 0,
        left: 0
    },
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