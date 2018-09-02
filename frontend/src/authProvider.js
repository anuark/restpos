import { AUTH_LOGIN, AUTH_LOGOUT, AUTH_ERROR, AUTH_CHECK } from 'react-admin';

export default (type, params) => {
    if (type === AUTH_LOGIN) {
        let { username, password } = params;
        console.log(process.env.REACT_APP_RESTPOS_HOST+"/auth");
        const request = new Request(process.env.REACT_APP_RESTPOS_HOST+"/auth", {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            // headers: new Headers({ 'Content-Type': 'application/x-www-form-urlencoded' })
            headers: new Headers({ 'Content-Type': 'application/json' })
        });

        return fetch(request)
        .then(res => res.ok ? res.json() : Promise.reject())
        .then(({key}) => {
            localStorage.setItem("authKey", key);
        });
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
        return localStorage.getItem('authKey') ? Promise.resolve() : Promise.reject();
        // return localStorage.getItem('authKey') ? Promise.resolve() : Promise.reject({redirectTo: '/no-access'});
    }
}
