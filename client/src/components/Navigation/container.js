import Navigation from './index';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux'
import { getLocations, setCompact } from '../Map/actions';

export const mapStateToProps = ({
    map: {
        center = [],
        loading = false,
        permission = false,
        compact = false,
    } = {}
}= {}) => ({
    center,
    loading,
    permission,
    compact,
});

export const mapDispatchToProps = dispatch => bindActionCreators({
    getLocations,
    setCompact,
}, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(Navigation)
