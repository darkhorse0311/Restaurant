import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'

import {
    setCenter,
    getLocations,
    setLoading,
    setPermission,
    getAllBusinesses,
} from './actions'

import {
    setShowModal,
    getItems,
    setSortMode,
} from '../Items/actions'


import Map from './index';


export const mapStateToProps = ({
    map: {
        locations = [],
        center = [],
        zoom = [],
        mapStyle = {},
        permission = false,
        allBusinesses = [],
        compact = false,
    } = {},
    business: {
        name = "",
    } = {},
} = {}) => ({
    locations,
    center,
    zoom,
    mapStyle,
    name,
    permission,
    allBusinesses,
    compact,
})

export const mapDispatchToProps = dispatch => (
    bindActionCreators(
        {
            getLocations,
            getItems,
            setCenter,
            setShowModal,
            setLoading,
            setPermission,
            getAllBusinesses,
            setSortMode,
        },
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Map);
