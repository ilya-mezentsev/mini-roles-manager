import { AccountCredentials } from '../../../services/api';
import { RegistrationResult } from '../../../store/registration/registration.types';

export interface SignUpActions {
    signUpAction: (c: AccountCredentials) => void;
    cleanSignUpAction: () => void;
}

export interface SignUpState {
    registrationResult: RegistrationResult | null;
}

export type SignUpProps = SignUpActions & SignUpState;
