import React, { useState, useEffect, useCallback } from "react";

import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";
import { useHistory } from "react-router-dom";

import { useUser } from "../../contexts/UserCtx";
import { signIn } from "../../lib/auth";

export default function SignInForm() {
  const { user, setUser } = useUser();
  const history = useHistory();
  const [hasError, setHasError] = useState(false);
  const [input, setInput] = useState({
    username: "",
    password: "",
  });

  const redirect = useCallback(
    (path) => {
      history.push(path);
    },
    [history]
  );

  const handleChange = (e) => {
    const { name, value } = e.target;
    setInput({ ...input, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const { username, password } = input;
      const res = await signIn(username, password);
      setUser(res);
      redirect("/");
    } catch (err) {
      setUser(null);
      setHasError(true);
    }
  };

  useEffect(() => {
    if (user) redirect("/");
  }, [user, redirect]);

  return (
    <Box mt={4}>
      <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
      >
        <Grid item xs={12} sm={9} md={6}>
          <Box m={2}>
            <Paper variant="outlined" square>
              <Box m={2}>
                <Typography component="h1" variant="h5">
                  Sign In
                </Typography>
                {hasError && (
                  <Typography color="error">
                    Incorrect username or password.
                  </Typography>
                )}
                <FormControl component="form" fullWidth>
                  <Box mt={3}>
                    <TextField
                      label="Username"
                      name="username"
                      placeholder="Enter username"
                      error={hasError}
                      onChange={(e) => handleChange(e)}
                      fullWidth
                      required
                    />
                  </Box>
                  <Box mt={3}>
                    <TextField
                      label="Password"
                      name="password"
                      placeholder="Enter password"
                      type="password"
                      error={hasError}
                      onChange={(e) => handleChange(e)}
                      fullWidth
                      required
                    />
                  </Box>
                  <Box mt={3}>
                    <Button
                      type="submit"
                      color="primary"
                      variant="contained"
                      onClick={(e) => handleSubmit(e)}
                      fullWidth
                    >
                      Sign in
                    </Button>
                  </Box>
                </FormControl>
              </Box>
            </Paper>
            <Box mt={1}>
              <Typography variant="caption" color="textSecondary">
                This application does not currently allow sign ups.
              </Typography>
            </Box>
          </Box>
        </Grid>
      </Grid>
    </Box>
  );
}
