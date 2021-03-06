import React, { useState, useEffect } from "react";

import Avatar from "@material-ui/core/Avatar";
import { deepOrange } from "@material-ui/core/colors";
import IconButton from "@material-ui/core/IconButton";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import Menu from "@material-ui/core/Menu";
import MenuItem from "@material-ui/core/MenuItem";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import CloudUploadIcon from "@material-ui/icons/CloudUpload";
import ExitToAppIcon from "@material-ui/icons/ExitToApp";
import { NavLink } from "react-router-dom";

import { useUser } from "../../contexts/UserCtx";
import { signOut, getUser } from "../../lib/auth";

const useStyles = makeStyles((theme) => ({
  orange: {
    color: theme.palette.getContrastText(deepOrange[500]),
    backgroundColor: deepOrange[500],
  },
}));

const UserMenu = () => {
  const [anchorEl, setAnchorEl] = useState(null);
  const { user, setUser } = useUser();

  const classes = useStyles();

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

  const handleSignOut = async () => {
    try {
      const res = await signOut();
      setUser(res);
    } catch (err) {
      console.error(err);
      setUser(null);
    }
  };

  const openMenu = (e) => {
    setAnchorEl(e.currentTarget);
  };

  const closeMenu = () => {
    setAnchorEl(null);
  };

  return (
    <React.Fragment>
      <IconButton onClick={openMenu}>
        <Avatar className={classes.orange}>
          {user.getUsername()[0].toUpperCase()}
        </Avatar>
      </IconButton>
      <Menu
        getContentAnchorEl={null}
        anchorOrigin={{
          vertical: "bottom",
          horizontal: "left",
        }}
        transformOrigin={{
          vertical: "top",
          horizontal: "left",
        }}
        keepMounted
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={closeMenu}
      >
        <MenuItem component={NavLink} to="/upload" onClick={closeMenu}>
          <ListItemIcon>
            <CloudUploadIcon />
          </ListItemIcon>
          <Typography variant="inherit">Upload video</Typography>
        </MenuItem>
        <MenuItem onClick={handleSignOut}>
          <ListItemIcon>
            <ExitToAppIcon />
          </ListItemIcon>
          <Typography variant="inherit">Sign out</Typography>
        </MenuItem>
      </Menu>
    </React.Fragment>
  );
};

export default UserMenu;
