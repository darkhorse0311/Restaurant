import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'
import { setShowModal, setSortMode } from './actions';

import Items from './index';


const mapStateToProps = ({
    business: {
        id = null,
        name = "",
        coordinates = {},
        photos = [],
        distance = null,
        r_id = null,
        items = [],
        showModal = false,
        sortMode = 'N',
        order = 'A',
    } = {},
} = {}) => ({
    id,
    name,
    coordinates,
    photos,
    distance,
    r_id,
    items,
    showModal,
    sortMode,
    order,
})


const mapDispatchToProps = dispatch => (
    bindActionCreators(
        {
            setShowModal,
            setSortMode
        },
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Items);
