
import { AccountInfoStore } from './account_info/account_info.store';
import { RolesVersionStore } from './roles_version/roles_version.store';
import { ResourceStore } from './resource/resource.store';
import { AppDataStore } from './app_data/app_data.store';
import { RoleStore } from './role/role.store';
import { SessionStore } from './session/session.store';
import { RegistrationStore } from './registration/registration.store';
import { PermissionStore } from './permission/permission.store';

export const accountInfoStore = new AccountInfoStore();
export const roleStore = new RoleStore();
export const resourceStore = new ResourceStore(roleStore);
export const rolesVersionStore = new RolesVersionStore(roleStore);
export const appDataStore = new AppDataStore(
    roleStore,
    resourceStore,
    rolesVersionStore,
);
export const sessionStore = new SessionStore();
export const registrationStore = new RegistrationStore();
export const permissionStore = new PermissionStore();
