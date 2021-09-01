import React, { useEffect } from "react";

import Button from "@material-ui/core/Button";
import { NavLink } from "react-router-dom";

import { useUser } from "../../contexts/UserCtx";
import { getUser } from "../../lib/auth";

export default function SignInButton() {
  const { setUser } = useUser();

  useEffect(() => {
    const handleUser = async () => {
      const user = await getUser();
      setUser(user);
    };

    handleUser();
  }, [setUser]);

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
