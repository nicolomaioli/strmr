import React, { useEffect, useState } from "react";

import Typography from "@material-ui/core/Typography";
import { Link } from "react-router-dom";

const { REACT_APP_STRMR_API_URL } = process.env;

const Home = () => {
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
    <React.Fragment>
      {error && <Typography color="error">{error.message}</Typography>}
      <ul>
        {videos.map((video, i) => {
          return (
            <li key={i}>
              <Link to={`/play/${video["ID"]}`}>{video["Title"]}</Link>
            </li>
          );
        })}
      </ul>
    </React.Fragment>
  );
};

export default Home;
