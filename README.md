# Strmr

Strmr is a Video On Demand platform built on AWS. It is fronted by a React
client (my first React project, be kind) and powered by a variety of services
including but not limited to Cognito, S3, Lambda, DynamoDB and MediaConvert.
The architecture looks something like this:

![](architecture.png "Architecture diagram")

## Repository Overview

`./terraform`: AWS resources and configuration for Strmr.
`./serverless`: Go Lambda resources for Strmr.
`./client`: A React application (`create-react-app`).
`./presentation`: A presentation I gave on Strmr at AND Digital.

## Run (Somewhat) Locally

- Start with deploying the infrastructure in `./terraform`. I have not included
    configuration for the Route53 Hosting Zone or ACM on account of I was using
    an already configured domain, so you might have to do that manually. You
    will need a TLS Certificate in your Region (for API Gateway), and one at
    Edge in `us-east-1` (for CloudFront).
- Reference `./serverless/video/serverless.yml` to create a `.env.dev` file in
    `./serverless/video` (I am planning on creating a script to automate this
    step). You can now deploy by running `make deploy`.
- Reference the codebase (oh boy I know, I know I am planning on automating
    this too) to create a `.env.development.local` file in `./client`. Then
    simply run `npm i` and `npm start` to start a local dev server.
- Profit?

## Planned Improvements

It is unlikely that I will have time to go back to this project anytime soon,
that said there are a few things I would like to address:

- Write some tests;
- Automate the creation of the various `.env*` files;
- Add support for `HLS`;
- Refactor the Terraform code to leverage Terraform remote state;
- Add resources and configuration to deploy the React client.
