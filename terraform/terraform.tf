provider "aws" {
  region = var.region
}

terraform {
  backend "s3" {
    key = "state"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
