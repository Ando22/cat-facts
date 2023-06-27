import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';

import { initDB } from './lib/db';

const root = document.createElement("div")
root.className = "container"
document.body.appendChild(root)
const rootDiv = ReactDOM.createRoot(root);

// init indexedDB
initDB();

rootDiv.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);