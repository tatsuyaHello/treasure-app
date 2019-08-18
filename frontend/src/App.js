import React, { Component } from "react";
import { BrowserRouter, Route, Link } from 'react-router-dom'

import { Form } from './Component/Form';
import { About } from './Component/About'

const App = () => (
  <BrowserRouter>
    <div>
      <Route exact path='/' component={Form} />
      <Route path='/about/:id' component={About} />
    </div>
  </BrowserRouter>
)

export default App;
