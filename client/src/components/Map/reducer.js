import { 
    SET_LOCATIONS, 
    SET_CENTER, LOADING, 
    SET_PERMISSION, 
    SET_BUSINESSES,
    SET_COMPACT,
} from './actions'

export const defaultState = {
    locations: [],
    center: [],
    zoom: [14],
    permission: false,
    loading: false,
    allBusinesses: [],
    compact: false,
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
        case SET_BUSINESSES:
            return {
                ...state,
                allBusinesses: payload
            }
        case SET_COMPACT:
            return {
                ...state,
                compact: payload
            }
        default: return {...state};
    }
} 