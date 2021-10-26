import { RegistrationStore } from './registration.store';

import * as api from '../../services/api';
import { ErrorResponse, SuccessResponse } from '../../services/api/shared';
import { UnknownErrorCode, UnknownErrorDescription } from '../shared/const';

jest.mock('../../services/api');

describe('registration tests', () => {
    let registrationStore = new RegistrationStore();
    const d = {
        login: 'some-login',
        password: 'some-password',
    };

    beforeEach(() => {
        registrationStore = new RegistrationStore();
    });

    it('registration success', async () => {
        // @ts-ignore
        api.signUp = jest.fn().mockResolvedValue(new SuccessResponse(null));

        await registrationStore.signUp(d);

        expect(api.signUp).toBeCalledWith(d);
        expect(registrationStore.registeredOk).toBeTruthy();
    });

    it('registration parsed error', async () => {
        // @ts-ignore
        api.signUp = jest.fn().mockResolvedValue(new ErrorResponse('some-error'));

        await registrationStore.signUp(d);

        expect(api.signUp).toBeCalledWith(d);
        expect(registrationStore.registeredOk).toBeFalsy();
        expect(registrationStore.registrationError).toEqual('some-error');
    });

    it('registration unknown error', async () => {
        // @ts-ignore
        api.signUp = jest.fn().mockRejectedValue('some-error');

        await registrationStore.signUp(d);

        expect(api.signUp).toBeCalledWith(d);
        expect(registrationStore.registeredOk).toBeFalsy();
        expect(registrationStore.registrationError).toEqual({
            code: UnknownErrorCode,
            description: UnknownErrorDescription,
        });
    });

    it('clean registration action', () => {
        registrationStore.registeredOk = true;
        registrationStore.registrationError = 'some-error' as any;

        registrationStore.cleanRegistrationAction();

        expect(registrationStore.registeredOk).toBeFalsy();
        expect(registrationStore.registrationError).toBeNull();
    });
});
