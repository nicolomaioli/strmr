import React, { useEffect, useState } from "react";

import Typography from "@material-ui/core/Typography";
import { useParams } from "react-router-dom";
import ShakaPlayer from "shaka-player-react";
import "shaka-player/dist/controls.css";

const { REACT_APP_STRMR_API_URL } = process.env;

export default function Play() {
  const [videoSrc, setVideoSrc] = useState("");
  const [error, setError] = useState(null);
  const { id } = useParams();

  useEffect(() => {
    const getVideo = async (id) => {
      try {
        const res = await fetch(`${REACT_APP_STRMR_API_URL}/video/${id}`);
        const video = await res.json();

        setVideoSrc(video["Path"]);
      } catch (err) {
        console.error(err);
        setError(err);
      }
    };

    getVideo(id);
  }, [id, setVideoSrc, setError]);

  return (
    <React.Fragment>
      {error ? (
        <Typography color="error">{error.message}</Typography>
      ) : (
        <ShakaPlayer autoPlay src={videoSrc} />
      )}
    </React.Fragment>
  );
}
