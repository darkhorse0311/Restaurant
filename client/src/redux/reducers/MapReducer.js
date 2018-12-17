import { UPDATE_SELECTED } from '../actions/MapActions'

const defaultState = {
    selected: null,
}

export default (state = defaultState, action) => {
    switch(action.type) {
        case UPDATE_SELECTED: return {...state, selected: action.payload};
        default: return {...state};
    }
} 