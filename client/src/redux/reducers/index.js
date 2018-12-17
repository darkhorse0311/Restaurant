import { combineReducers } from 'redux';
import mapReducer from './MapReducer';

export default combineReducers({
    info: mapReducer
})
