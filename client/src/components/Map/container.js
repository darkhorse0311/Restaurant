import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'

import {
    setCordinates,
    getLocations,
    setCenter
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
        lat = null,
        lon = null,
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
    lat,
    lon,
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
            setCordinates,
            getLocations,
            getItems,
            setCenter,
            setShowModal,
        },
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Map);
