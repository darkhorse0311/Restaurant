import Navigation from './index';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'
import { getLocations } from '../Map/actions';

export const mapStateToProps = ({
    map: {
        center = [],
        loading = false,
        permission = false
    } = {}
}= {}) => ({
    center,
    loading,
    permission,
});

export const mapDispatchToProps = dispatch => bindActionCreators({
    getLocations,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
