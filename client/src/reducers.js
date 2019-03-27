import { combineReducers } from 'redux';
import map from './components/Map/reducer';
import business from './components/Items/reducer';

export default combineReducers({
    map,
    business,
})
