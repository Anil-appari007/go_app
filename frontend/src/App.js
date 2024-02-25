import './App.css';
import Create from './components/create';
import GetItems from './components/getAll';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import HomePage from './components/homePage';
import { API_URL } from './config';
import { useEffect } from 'react';

function App() {
  return (
      <div className="main">
        <BrowserRouter>
          <Routes>
            <Route exact path='/' Component={HomePage}></Route>
            <Route exact path='/getAll' Component={GetItems}></Route>
            <Route exact path='/addItem' Component={Create}></Route>
          </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;
