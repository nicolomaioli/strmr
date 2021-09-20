import Amplify, { Storage } from "aws-amplify";

import awsconfig from "./awsconfig";

Amplify.configure(awsconfig);

const putObject = async (key, object, options) => {
  return Storage.put(key, object, options);
};

export { putObject };
