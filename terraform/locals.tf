locals {
  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
