import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'
import { setShowModal } from './actions';

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
    },
}) => ({
    id,
    name,
    coordinates,
    photos,
    distance,
    r_id,
    items,
    showModal
})


const mapDispatchToProps = dispatch => (
    bindActionCreators(
        {setShowModal},
        dispatch
    )
);

export default connect(mapStateToProps, mapDispatchToProps)(Items);
