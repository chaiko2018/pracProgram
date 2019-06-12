import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

import createStore from 'redux';

import { createMuiTheme, MuiThemeProvider  } from '@material-ui/core/styles';
import red from '@material-ui/core/colors/red';
import blue from '@material-ui/core/colors/blue';

import Provider from 'react-redux';

import Router from 'react-router-dom';


const store = createStore(
  
);

// テーマカスタマイズ
const theme = createMuiTheme({
  palette: {
    type: 'light',
    primary: red,
    secondary: blue,
  },
});

ReactDOM.render(
  <Provider store={store}>
    <MuiThemeProvider theme={theme}>
      <Router>
        <App />
      </Router>
    </MuiThemeProvider>
  </Provicer>
  , document.getElementById('root')
);

