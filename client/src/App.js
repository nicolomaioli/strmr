import React from "react";

import Container from "@material-ui/core/Container";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import SignInForm from "./components/authentication/SignInForm";
import Home from "./components/Home";
import Nav from "./components/navigation/Nav";
import Play from "./components/video/Play";
import Upload from "./components/video/Upload";
import { UserProvider } from "./contexts/UserCtx";

const App = () => {
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
            <Route exact path="/upload">
              <Upload />
            </Route>
            <Route path="/play/:id">
              <Play />
            </Route>
          </Switch>
        </Container>
      </Router>
    </UserProvider>
  );
};

export default App;
