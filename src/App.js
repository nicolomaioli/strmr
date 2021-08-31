import React from 'react'
import Container from '@material-ui/core/Container'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { UserProvider } from './contexts/UserCtx'
import Nav from './components/Nav'
import Home from './components/Home'

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
