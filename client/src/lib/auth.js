import Amplify, { Auth } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

const signOut = async () => {
  return Auth.signOut();
};

const signIn = async (username, password) => {
  return Auth.signIn(username, password);
};

const getUser = async () => {
  return Auth.currentAuthenticatedUser();
};

const getCurrentSession = async () => {
  return Auth.currentSession();
};

export { signOut, signIn, getUser, getCurrentSession };
