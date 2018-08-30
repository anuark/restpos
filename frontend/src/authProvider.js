import { AUTH_LOGIN, AUTH_LOGOUT, AUTH_ERROR, AUTH_CHECK } from 'react-admin';

export default (type, params) => {
    if (type === AUTH_LOGIN) {
        
    }
    if (type === AUTH_LOGOUT) {
        localStorage.removeItem('authKey');
    }
    if (type === AUTH_ERROR) {
        const status = params.status;
        if (status === 401 || status === 403) {
            localStorage.removeItem('authKey');
            return Promise.reject();
        }
        return Promise.resolve();
    }
    if (type === AUTH_CHECK) {
        return localStorage.getItem('token') ? Promise.resolve() : Promise.reject({redirectTo: '/no-access'});
    }
}
