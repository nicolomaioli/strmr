import Amplify, { Auth } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

export function signOut() {
  return Auth.signOut();
}

export function getUser() {
  return Auth.currentAuthenticatedUser();
}

export function signIn(username, password) {
  return Auth.signIn(username, password);
}
