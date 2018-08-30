import React, { Component } from 'react';
import './App.css';
import { Admin, Resource, fetchUtils } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import { ItemList } from './items';
import CatchAll from './CatchAll';
import AuthProvider from './authProvider';

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }

    const token = localStorage.getItem('authKey');
    options.headers.set('Authorization', `Bearer ${token}`);
    return fetchUtils.fetchJson(url, options);
}

const dataProvider = simpleRestProvider('http://localhost:81', httpClient);

class App extends Component {
    render() {
        return (
            <Admin dataProvider={dataProvider} authProvider={AuthProvider} title="Restaurant POS" catchAll={CatchAll} >
                <Resource name="items" list={ItemList} />
                <Resource name="orders" />
            </Admin>
        );
    }
}

export default App;
