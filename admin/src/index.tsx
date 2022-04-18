import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import {Authentication} from './pages/authentication';
import { store } from './state';
import "./assets/css/phoenix.min.css";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
const App = () => {
  return (
    <Provider store={store}>
      <Authentication />
    </Provider>
  );
};
root.render(
  <App />
);