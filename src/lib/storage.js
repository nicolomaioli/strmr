import Amplify, { Storage } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

export async function putObject(key, object) {
  return Storage.put(key, object);
}
