import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import Jogos from './components/Jogos';


ReactDOM.render(
  <React.StrictMode>
    <App />
    <Jogos />
  </React.StrictMode>,
  document.getElementById('root')
);


reportWebVitals();
