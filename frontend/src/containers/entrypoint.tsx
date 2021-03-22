import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { EntrypointState, EntrypointActions, EntrypointProps } from './entrypoint.types';
import { PrivateNavigation } from './private';
import { PublicNavigation } from './public';
import { DispatchToPropsFn, StateToPropsFn } from '../shared/types';
import { login } from '../store/session/actions';

class Entrypoint extends React.Component<EntrypointProps, any> {
    componentDidMount() {
        this.props.login();
    }

    render() {
        return (
            <>
                {
                    !!this.props?.userSession?.session?.id
                        ? <PrivateNavigation/>
                        : <PublicNavigation/>
                }
            </>
        );
    }
}

const mapStateToProps: StateToPropsFn<EntrypointState> = () => state => ({
    userSession: state.userSession,
});

const mapDispatchToProps: DispatchToPropsFn<EntrypointActions> = () => dispatch => ({
    login: bindActionCreators(login, dispatch),
});

export default connect(mapStateToProps(), mapDispatchToProps())(Entrypoint);
