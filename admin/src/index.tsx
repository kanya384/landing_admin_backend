import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import { store } from './state';
import "./assets/css/phoenix.min.css";
import Router from './feature/router/router';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
const App = () => {
  return (
    <Provider store={store}>
      <Router />
    </Provider>
  );
};
root.render(
  <App />
);