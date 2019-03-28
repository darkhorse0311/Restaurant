import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'

import {
    setCenter,
    getLocations,
    setLoading,
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
})

const mapDispatchToProps = dispatch => (
    bindActionCreators(
        {
            getLocations,
            getItems,
            setCenter,
            setShowModal,
            setLoading,
        },
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Map);
