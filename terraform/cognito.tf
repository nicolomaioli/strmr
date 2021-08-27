resource "aws_cognito_user_pool" "this" {
  name = "${var.application}-${terraform.workspace}"

  auto_verified_attributes = ["email"]

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}

resource "aws_cognito_user_pool_client" "this" {
  name = "${var.application}-${terraform.workspace}"

  user_pool_id                  = aws_cognito_user_pool.this.id
  generate_secret               = false
  prevent_user_existence_errors = "ENABLED"

  explicit_auth_flows = [
    "ALLOW_CUSTOM_AUTH",
    "ALLOW_REFRESH_TOKEN_AUTH",
    "ALLOW_USER_PASSWORD_AUTH",
    "ALLOW_USER_SRP_AUTH",
  ]
}
