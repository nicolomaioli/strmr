import Amplify, { Auth } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

export async function signOut() {
  return Auth.signOut();
}

export async function signIn(username, password) {
  return Auth.signIn(username, password);
}

export async function getUser() {
  return Auth.currentAuthenticatedUser();
}

export async function getCurrentSession() {
  return Auth.currentSession();
}
