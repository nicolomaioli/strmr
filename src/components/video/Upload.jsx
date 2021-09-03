import React, { useState } from "react";

import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import Grid from "@material-ui/core/Grid";
import Input from "@material-ui/core/Input";
import LinearProgress from "@material-ui/core/LinearProgress";
import Paper from "@material-ui/core/Paper";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";

import { putObject } from "../../lib/storage";

export default function Upload() {
  const [error, setError] = useState(null);
  const [progress, setProgress] = useState(0);

  const progressCallback = (p) => {
    setProgress(Math.trunc((p.loaded * 100) / p.total));
  };

  const handleChange = async (e) => {
    try {
      console.log("path", e.target.value);
      const res = await putObject("uploads/video.mkv", e.target.files[0], {
        progressCallback,
      });
      console.log("res", res);
    } catch (err) {
      setError(err);
    }
  };

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
                  Upload a video
                </Typography>
                {error && (
                  <Typography color="error">{error.message}</Typography>
                )}
                <FormControl component="form" fullWidth>
                  <Box mt={3}>
                    <TextField
                      label="Tile"
                      name="title"
                      placeholder="Title"
                      required
                      fullWidth
                    />
                  </Box>
                  <Box mt={3}>
                    <Input
                      accept="image/*"
                      style={{ display: "none" }}
                      id="raised-button-file"
                      onChange={(e) => handleChange(e)}
                      required
                      type="file"
                    />
                    <FormLabel htmlFor="raised-button-file">
                      <Button variant="contained" component="span">
                        Select File
                      </Button>
                    </FormLabel>
                  </Box>
                  <Box mt={3}>
                    <Button
                      type="submit"
                      color="primary"
                      variant="contained"
                      // onClick={(e) => handleSubmit(e)}
                      fullWidth
                    >
                      Upload
                    </Button>
                  </Box>
                </FormControl>
              </Box>
            </Paper>
            <LinearProgress variant="determinate" value={progress} />
          </Box>
        </Grid>
      </Grid>
    </Box>
  );
}
