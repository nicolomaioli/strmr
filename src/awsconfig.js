const {
  REACT_APP_COGNITO_USER_POOL_ID,
  REACT_APP_COGNITO_USER_POOL_WEB_CLIENT_ID,
  REACT_APP_COGNITO_REDIRECT_SIGN_IN,
  REACT_APP_COGNITO_REDIRECT_SIGN_OUT,
  REACT_APP_COGNITO_DOMAIN
} = process.env

const awsconfig = {
  region: 'eu-west-1',
  userPoolId: REACT_APP_COGNITO_USER_POOL_ID,
  userPoolWebClientId: REACT_APP_COGNITO_USER_POOL_WEB_CLIENT_ID,
  authenticationFlowType: 'USER_PASSWORD_AUTH',
  oauth: {
    domain: REACT_APP_COGNITO_DOMAIN,
    scope: ['email', 'profile'],
    redirectSignIn: REACT_APP_COGNITO_REDIRECT_SIGN_IN,
    redirectSignOut: REACT_APP_COGNITO_REDIRECT_SIGN_OUT,
    responseType: 'code'
  }
}

export default awsconfig
