import React from "react";

import Button from "@material-ui/core/Button";
import { NavLink } from "react-router-dom";

export default function SignInButton() {
  return (
    <Button
      component={NavLink}
      to="/signin"
      variant="contained"
      color="primary"
    >
      Sign In
    </Button>
  );
}
