import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'

import {
    setCenter,
    getLocations,
    setLoading,
    setPermission,
} from './actions'

import {
    setShowModal,
    getItems,
} from '../Items/actions'


import Map from './index';


export const mapStateToProps = ({
    map: {
        locations = [],
        center = [],
        zoom = [],
        mapStyle = {},
        permission = false,
    },
    business: {
        name = "",
    }
}) => ({
    locations,
    center,
    zoom,
    mapStyle,
    name,
    permission,
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
        },
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Map);
