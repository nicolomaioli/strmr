# Strmr Infrastructure

This repository contains the necessary infrastructure to deploy the Strmr
application to AWS.

## Terraform

### Provision

Create an S3 bucket to initialise the Terraform backend. Create a file named
`terraform/terraform.backend.tfvars` with the content:

```terraform
bucket = "<YOUR_BUCKET_NAME>"
region = "<YOUR_REGION>"
```

Optionally create a `terraform/terraform.tfvars` file to personalise some of
the values in `terraform/variables.tf`.

Initialise Terraform:

```sh
cd terraform
terraform init -backend-config=terraform.backend.tfvars
terraform workspace new dev
```

Create and apply a plan:

```sh
terraform plan -out plan.o
terraform apply plan.o
```

After the infrastructure has been provisioned, you can collect the outputs by
running:

```sh
terraform output
```
