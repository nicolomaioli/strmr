import React from "react";

import Container from "@material-ui/core/Container";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import SignInForm from "./components/authentication/SignInForm";
import Home from "./components/Home";
import Nav from "./components/navigation/Nav";
import { UserProvider } from "./contexts/UserCtx";

function App() {
  return (
    <UserProvider>
      <Router>
        <Nav />
        <Container>
          <Switch>
            <Route exact path="/">
              <Home />
            </Route>
            <Route exact path="/signin">
              <SignInForm />
            </Route>
          </Switch>
        </Container>
      </Router>
    </UserProvider>
  );
}

export default App;
