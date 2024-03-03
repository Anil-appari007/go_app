import './App.css';
// import Create from './components/create';
// import GetItems from './components/getAll';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
// import HomePage from './components/homePage';
// import { API_URL } from './config';
// import { useEffect } from 'react';
import Table from './components/Test';

function App() {
  return (
      <div>
        <h1>Sample InventoryList Application</h1>
        <em>Written with React, Go lang & Postgresql </em>
        <BrowserRouter>
          <Routes>
            <Route exact path='/' Component={Table}></Route>
            {/* <Route exact path='/' Component={HomePage}></Route>
            <Route exact path='/getAll' Component={GetItems}></Route>
            <Route exact path='/addItem' Component={Create}></Route>
            <Route exact path='/test' Component={Table}></Route> */}
          </Routes>
        </BrowserRouter>
        <footer>
          <p>v1.00</p>
        </footer>
      </div>
  );
}

export default App;
