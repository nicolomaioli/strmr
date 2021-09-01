import Amplify, { Auth } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

export function signOut() {
  return Auth.signOut();
}

export function getUser() {
  return Auth.currentAuthenticatedUser();
}

export function signIn() {
  return Auth.signIn(
    // nothing to see here, move along!
    process.env.REACT_APP_TEST_USER_USERNAME,
    process.env.REACT_APP_TEST_USER_PASSWORD
  );
}
