import mockAxios from 'axios';
import configureMockStore from 'redux-mock-store'
import thunk from 'redux-thunk'
// import fetchMock from 'fetch-mock'

import { 
    SET_LOCATIONS, 
    SET_CENTER, 
    LOADING, 
    SET_PERMISSION,
    getLocations,
    setCenter,
    setLoading,
    setPermission,
} from '../../components/Map/actions'

const middlewares = [thunk]
const mockStore = configureMockStore(middlewares)


describe('Map Actions', () => {
    // afterEach(() => {
    //     fetchMock.restore()
    // })

    it('should create actions LOADING twice and SET_LOCATIONS once', async () => {
        const expectedActions = [
            {type: LOADING, payload: true},
            {type: SET_LOCATIONS, payload: []},
            {type: LOADING, payload: false},
        ]
        const store = mockStore({})

        // setup axios
        mockAxios.get.mockImplementationOnce(() =>
            Promise.resolve({
                data: []
            })
        );

        return store.dispatch(getLocations()).then(() => {
            expect(store.getActions()).toEqual(expectedActions)
        })
    })

    it('setCenter should call SET_CENTER actions', () => {
        const payload = [-74.0060, 40.7128];
        const expectedAction = {
            type: SET_CENTER,
            payload,
        }
        expect(setCenter(payload)).toEqual(expectedAction)
    })

    it('setLoading should call LOADING actions', () => {
        const payload = true;
        const expectedAction = {
            type: LOADING,
            payload,
        }
        expect(setLoading(payload)).toEqual(expectedAction)
    })

    it('setPermission should call SET_PERMISSION actions', () => {
        const payload = true;
        const expectedAction = {
            type: SET_PERMISSION,
            payload,
        }
        expect(setPermission(payload)).toEqual(expectedAction)
    })
  })
