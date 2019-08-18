import React, { Component } from "react";
import { BrowserRouter, Route, Link } from 'react-router-dom'

import { Form } from './Component/Form';

const App = () => (
  <BrowserRouter>
    <div>
      <ul>
       <li><Link to='/'>Home</Link></li>
       <li><Link to='/about'>About</Link></li>
     </ul>

      <Route exact path='/' component={Form} />
      <Route path='/about' component={About} />
    </div>
  </BrowserRouter>
)

const About = () => (
  <div>
    <h2>About</h2>
    <p>フレンズに投票するページです</p>
  </div>
)


export default App;
