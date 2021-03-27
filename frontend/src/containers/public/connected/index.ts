import { connect } from 'react-redux';

import {
    SignUp as SignUpContainer,
    mapDispatchToProps as SignUpMapDispatchToProps,
    mapStateToProps as SignUpMapStateToProps,
} from '../sign_up/sign_up';

import {
    SignIn as SignInContainer,
    mapDispatchToProps as SignInMapDispatchToProps,
    mapStateToProps as SignInMapStateToProps,
} from '../sign_in/sign_in';

export const SignUp = connect(SignUpMapStateToProps(), SignUpMapDispatchToProps())(SignUpContainer);
export const SignIn = connect(SignInMapStateToProps(), (SignInMapDispatchToProps))(SignInContainer);
