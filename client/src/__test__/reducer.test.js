import { createStore } from 'redux';
import rootReducer from '../reducers';

import { defaultState as mapState } from '../components/Map/reducer';
import { defaultState as businessState } from '../components/Items/reducer';


let store = createStore(rootReducer);

describe('rootReducer', () => {
    it('map state should return correct state', () => {
        expect(store.getState().map).toEqual(mapState)
    });
    it('map state should return correct state', () => {
        expect(store.getState().business).toEqual(businessState)
    });
});
