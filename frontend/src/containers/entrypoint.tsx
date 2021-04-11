import { useEffect } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { EntrypointState, EntrypointActions, EntrypointProps } from './entrypoint.types';
import { DispatchToPropsFn, StateToPropsFn } from '../shared/types';
import { login } from '../store/session/actions';
import { PublicNavigation as PublicEntryPoint } from './public';
import { PrivateEntrypoint } from './private';

const Entrypoint = (props: EntrypointProps) => {
    useEffect(() => props.login(), []);

    return (
        <>
            {
                !!props.userSession?.session?.id
                    ? <PrivateEntrypoint />
                    : <PublicEntryPoint />
            }
        </>
    );
}

const mapStateToProps: StateToPropsFn<EntrypointState> = () => state => ({
    userSession: state.userSession,
});

const mapDispatchToProps: DispatchToPropsFn<EntrypointActions> = () => dispatch => ({
    login: bindActionCreators(login, dispatch),
});

export default connect(mapStateToProps(), mapDispatchToProps())(Entrypoint);
