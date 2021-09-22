import React, { useEffect, useState } from "react";

import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import { useParams } from "react-router-dom";
import ShakaPlayer from "shaka-player-react";
import "shaka-player/dist/controls.css";

const { REACT_APP_STRMR_API_URL } = process.env;

const Play = () => {
  const [video, setVideo] = useState("");
  const [error, setError] = useState(null);
  const { id } = useParams();

  useEffect(() => {
    const getVideo = async (id) => {
      try {
        const res = await fetch(`${REACT_APP_STRMR_API_URL}/video/${id}`);
        const video = await res.json();

        setVideo(video);
      } catch (err) {
        setError(err);
      }
    };

    getVideo(id);
  }, [id, setVideo, setError]);

  return (
    <Box mt={4}>
      {error && <Typography color="error">{error.message}</Typography>}
      <ShakaPlayer autoPlay src={video["Path"]} />
      <Box mt={2}>
        <Typography variant="h5">{video["Title"]}</Typography>
      </Box>
    </Box>
  );
};

export default Play;
