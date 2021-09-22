import React, { useEffect, useState } from "react";

import Box from "@material-ui/core/Box";
import Card from "@material-ui/core/Card";
import CardContent from "@material-ui/core/CardContent";
import CardMedia from "@material-ui/core/CardMedia";
import Grid from "@material-ui/core/Grid";
import { makeStyles } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import { Link } from "react-router-dom";

const { REACT_APP_STRMR_API_URL } = process.env;

const useStyles = makeStyles({
  card: {
    height: "100%",
  },
});

const Home = () => {
  const classes = useStyles();

  const [videos, setVideos] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getVideos = async () => {
      try {
        const res = await fetch(`${REACT_APP_STRMR_API_URL}/video`);
        const videos = await res.json();

        setVideos(videos);
      } catch (err) {
        setError(err);
      }
    };

    getVideos();
  }, [setVideos, setError]);

  return (
    <Box mt={4}>
      {error && <Typography color="error">{error.message}</Typography>}
      <Grid container spacing={3}>
        {videos.map((video, i) => {
          return (
            <Grid key={i} item xs={12} sm={6} md={4}>
              <Card className={classes.card}>
                <Link to={`/play/${video["ID"]}`}>
                  <CardMedia
                    component="img"
                    height="160"
                    image={video["PosterFrame"]}
                    alt={video["Title"]}
                  />
                </Link>
                <CardContent>
                  <Typography gutterBottom variant="h5" component="div">
                    {video["Title"]}
                  </Typography>
                  <Typography variant="body2" color="text.secondary">
                    {video["Username"]}
                  </Typography>
                </CardContent>
              </Card>
            </Grid>
          );
        })}
      </Grid>
    </Box>
  );
};

export default Home;
