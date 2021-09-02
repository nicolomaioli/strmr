const {
  REACT_APP_AWS_REGION,
  REACT_APP_AWS_COGNITO_USER_POOL_ID,
  REACT_APP_AWS_COGNITO_USER_POOL_WEB_CLIENT_ID,
  REACT_APP_AWS_COGNITO_IDENTITY_POOL_ID,
  REACT_APP_AWS_COGNITO_REDIRECT_SIGN_IN,
  REACT_APP_AWS_COGNITO_REDIRECT_SIGN_OUT,
  REACT_APP_AWS_COGNITO_DOMAIN,
  REACT_APP_AWS_VIDEO_BUCKET_NAME,
} = process.env;

const awsconfig = {
  Auth: {
    region: REACT_APP_AWS_REGION,
    userPoolId: REACT_APP_AWS_COGNITO_USER_POOL_ID,
    userPoolWebClientId: REACT_APP_AWS_COGNITO_USER_POOL_WEB_CLIENT_ID,
    identityPoolId: REACT_APP_AWS_COGNITO_IDENTITY_POOL_ID,
    mandatorySignIn: true,
    authenticationFlowType: "USER_PASSWORD_AUTH",
    oauth: {
      domain: REACT_APP_AWS_COGNITO_DOMAIN,
      scope: ["email", "profile"],
      redirectSignIn: REACT_APP_AWS_COGNITO_REDIRECT_SIGN_IN,
      redirectSignOut: REACT_APP_AWS_COGNITO_REDIRECT_SIGN_OUT,
      responseType: "code",
    },
  },
  Storage: {
    AWSS3: {
      bucket: REACT_APP_AWS_VIDEO_BUCKET_NAME,
      region: REACT_APP_AWS_REGION,
    },
  },
};

export default awsconfig;
