import { mapStateToProps, mapDispatchToProps } from '../../components/Navigation/container';
// import { getLocations } from '../../components/Map/actions'
// import { bindActionCreators } from 'redux';

const initialState = {
    map: {
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
        compact: false,
    }
}

describe('Navigator Contianer', () => {

    it('should return expected state when passed in expectedState' , () => {
        const expectedState = {
            center: [],
            permission: false,
            loading: false,
            compact: false,
            locations: [],
        }
        expect(mapStateToProps(initialState)).toEqual(expectedState);
    });

    it('should return expected state when no values passed in' , () => {
        const expectedState = {
            center: [],
            permission: false,
            loading: false,
            compact: false,
            locations: [],
        }
        expect(mapStateToProps()).toEqual(expectedState);
    });

    it('should return getLocations prop function', () => {
        const dispatch = jest.fn();
        const response = mapDispatchToProps(dispatch)
        expect("getLocations" in response).toBe(true);
    });
});