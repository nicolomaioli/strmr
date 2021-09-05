import React, { useState } from "react";

import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import Grid from "@material-ui/core/Grid";
import Input from "@material-ui/core/Input";
import LinearProgress from "@material-ui/core/LinearProgress";
import Paper from "@material-ui/core/Paper";
import Snackbar from "@material-ui/core/Snackbar";
import TextField from "@material-ui/core/TextField";
import Typography from "@material-ui/core/Typography";
import MuiAlert from "@material-ui/lab/Alert";
import { v4 as uuid } from "uuid";

import { useUser } from "../../contexts/UserCtx";
import { putObject } from "../../lib/storage";

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
}

export default function Upload() {
  const { user } = useUser();
  const [title, setTitle] = useState("");
  const [file, setFile] = useState(null);
  const [data, setData] = useState(null);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(null);
  const [progress, setProgress] = useState(0);

  // Loads the file f into a throw-away HTML5 video element
  // to extract mime type and duration.
  // Throws an error if the file is not a video file.
  const getVideoMetadata = async (f) =>
    new Promise((resolve, reject) => {
      const video = document.createElement("video");
      video.preload = "metadata";

      video.onloadeddata = () =>
        resolve({
          duration: video.duration.toString(),
          width: video.videoWidth.toString(),
          height: video.videoHeight.toString(),
        });

      video.onerror = () => reject(new Error("Invalid video file"));

      video.src = window.URL.createObjectURL(f);
    });

  const progressCallback = (p) => {
    const currentProgress = Math.trunc((p.loaded * 100) / p.total);
    setProgress(Math.trunc((p.loaded * 100) / p.total));

    if (currentProgress === 100) {
      setSuccess(true);
      setProgress(0);
    }
  };

  const handleClose = () => {
    setSuccess(false);
  };

  const handleTitleChange = (e) => {
    setTitle(e.target.value);
  };

  const handleFileChange = async (e) => {
    try {
      const blob = e.target.files[0];
      const data = await getVideoMetadata(blob);
      setData(data);
      setFile(blob);
    } catch (err) {
      setFile(null);
      setError(err);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const username = user.getUsername();
      const id = uuid();
      const path = `uploads/${username}/${id}`;
      const res = await putObject(path, file, {
        progressCallback,
        metadata: {
          title,
          username,
          id,
          ...data,
        },
      });
      console.log("res", res);
    } catch (err) {
      setError(err);
    }
  };

  return (
    <React.Fragment>
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
                        onChange={(e) => handleTitleChange(e)}
                        required
                        fullWidth
                      />
                    </Box>
                    <Box mt={3}>
                      <Input
                        accept="video/*"
                        style={{ display: "none" }}
                        id="raised-button-file"
                        onChange={(e) => handleFileChange(e)}
                        required
                        type="file"
                      />
                      <FormLabel htmlFor="raised-button-file">
                        <Button variant="contained" component="span">
                          Select File
                        </Button>
                      </FormLabel>
                    </Box>
                    {file && (
                      <Box mt={3}>
                        <Typography>{file.name}</Typography>
                      </Box>
                    )}
                    <Box mt={3}>
                      <Button
                        type="submit"
                        color="primary"
                        variant="contained"
                        onClick={(e) => handleSubmit(e)}
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
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          File successfully uploaded
        </Alert>
      </Snackbar>
    </React.Fragment>
  );
}
