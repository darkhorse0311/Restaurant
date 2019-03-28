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
} from '../Items/actions'

import {
    getItems,
} from '../Items/actions'

import Map from './index';


const mapStateToProps = ({
    map: {
        locations = [],
        center = [],
        zoom = [],
        places = [],
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
    places,
    mapStyle,
    name,
    permission,
})

const mapDispatchToProps = dispatch => (
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
