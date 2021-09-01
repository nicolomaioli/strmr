import React from 'react'

import Container from '@material-ui/core/Container'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'

import Home from './components/Home'
import Nav from './components/Nav'
import { UserProvider } from './contexts/UserCtx'

function App () {
  return (
    <UserProvider>
      <Nav />
      <Container>
        <Router>
          <Switch>
            <Route exact path="/">
              <Home />
            </Route>
          </Switch>
        </Router>
      </Container>
    </UserProvider>
  )
}

export default App
