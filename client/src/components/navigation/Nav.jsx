import React, { useEffect } from "react";

import AppBar from "@material-ui/core/AppBar";
import { makeStyles } from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { NavLink } from "react-router-dom";

import { useUser } from "../../contexts/UserCtx";
import { getUser } from "../../lib/auth";
import SignInButton from "../authentication/SignInButton";
import UserMenu from "./UserMenu";

const useStyles = makeStyles(() => ({
  title: {
    flexGrow: 1,
    textDecoration: "none",
  },
}));

const Nav = () => {
  const { user, setUser } = useUser();

  useEffect(() => {
    const handleUser = async () => {
      try {
        const user = await getUser();
        setUser(user);
      } catch (err) {
        setUser(null);
      }
    };

    handleUser();
  }, [setUser]);

  const classes = useStyles();

  return (
    <React.Fragment>
      <AppBar elevation={0}>
        <Toolbar>
          <Typography
            component={NavLink}
            to="/"
            variant="h6"
            color="inherit"
            className={classes.title}
          >
            Strmr
          </Typography>
          {user ? <UserMenu /> : <SignInButton />}
        </Toolbar>
      </AppBar>
      {/*
        Would you believe me if I told you this is in the actual docs?
        https://material-ui.com/components/app-bar/#fixed-placement
      */}
      <Toolbar />
    </React.Fragment>
  );
};

export default Nav;
