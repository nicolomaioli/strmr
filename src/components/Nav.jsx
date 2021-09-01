import React, { useEffect } from "react";

import AppBar from "@material-ui/core/AppBar";
import Avatar from "@material-ui/core/Avatar";
import { deepOrange } from "@material-ui/core/colors";
import { makeStyles } from "@material-ui/core/styles";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import { NavLink } from "react-router-dom";

import { useUser } from "../contexts/UserCtx";
import { getUser } from "../lib/auth";
import SignInButton from "./authentication/SignInButton";

const useStyles = makeStyles((theme) => ({
  title: {
    flexGrow: 1,
    textDecoration: "none",
  },
  orange: {
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
  },
}));

export default function Nav() {
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
          {user ? (
            <Avatar className={classes.orange}>
              {user.getUsername()[0].toUpperCase()}
            </Avatar>
          ) : (
            <SignInButton />
          )}
        </Toolbar>
      </AppBar>
      {/*
        Would you believe me if I told you this is in the actual docs?
        https://material-ui.com/components/app-bar/#fixed-placement
      */}
      <Toolbar />
    </React.Fragment>
  );
}
