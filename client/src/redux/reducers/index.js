import { combineReducers } from 'redux';
import mapReducer from './mapReducer';
import businessReducer from './businessReducer';

export default combineReducers({
    map : mapReducer,
    business: businessReducer,
})
