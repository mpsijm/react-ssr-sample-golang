import React from 'react';
import ReactDOM from 'react-dom';
import {BrowserRouter} from 'react-router-dom'

import App from './components/app'
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';

const initialState = window.__PRELOADED_STATE__ ? window.__PRELOADED_STATE__ : {};

ReactDOM.hydrate((
    <BrowserRouter>
        <App store={initialState}/>
    </BrowserRouter>
), document.getElementById('app'));
