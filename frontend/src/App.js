import React, { Component } from 'react';
import './App.css';
import { Admin, Resource, fetchUtils } from 'react-admin';
import jsonServerRestClient from 'ra-data-json-server';
import { ItemList } from './items';
import CatchAll from './CatchAll';
import AuthProvider from './authProvider';
import Dashboard from './Dashboard';

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }

    const token = localStorage.getItem('authKey');
    options.headers.set('Authorization', `Bearer ${token}`);
    return fetchUtils.fetchJson(url, options);
}
const dataProvider = jsonServerRestClient(process.env.REACT_APP_RESTPOS_HOST, httpClient);
class App extends Component {
    render() {
        return (
            <Admin dataProvider={dataProvider} authProvider={AuthProvider} title="Restaurant POS" dashboard={Dashboard} catchAll={CatchAll} >
                <Resource name="items" list={ItemList} />
                {/* <Resource name="orders" /> */}
            </Admin>
        );
    }
}

export default App;
