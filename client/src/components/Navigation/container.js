import Navigation from './index';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'
import { getLocations } from '../Map/actions';

const mapStateToProps = ({
    map: {
        center =[],
        loading,
    }
}) => ({
    center,
    loading,
});

const mapDispatchToProps = dispatch => bindActionCreators({
    getLocations,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
