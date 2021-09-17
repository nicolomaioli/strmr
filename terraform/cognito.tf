resource "aws_cognito_user_pool" "this" {
  name = "${var.application}-${terraform.workspace}"

  auto_verified_attributes = ["email"]

  admin_create_user_config {
    allow_admin_create_user_only = true
  }

  tags = local.tags
}

resource "aws_cognito_user_group" "basic" {
  name         = "basic"
  user_pool_id = aws_cognito_user_pool.this.id
  description  = "A basic user on Strmr"
  precedence   = 1
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

resource "aws_cognito_identity_pool" "this" {
  identity_pool_name = "${var.application}-${terraform.workspace}"

  allow_unauthenticated_identities = false

  cognito_identity_providers {
    client_id               = aws_cognito_user_pool_client.this.id
    provider_name           = aws_cognito_user_pool.this.endpoint
    server_side_token_check = false
  }

  tags = local.tags
}

data "aws_iam_policy_document" "cognito_authenticated_role" {
  statement {
    actions = [
      "sts:AssumeRoleWithWebIdentity"
    ]

    principals {
      type        = "Federated"
      identifiers = ["cognito-identity.amazonaws.com"]
    }

    condition {
      test     = "StringEquals"
      variable = "cognito-identity.amazonaws.com:aud"
      values   = [aws_cognito_identity_pool.this.id]
    }

    condition {
      test     = "ForAnyValue:StringLike"
      variable = "cognito-identity.amazonaws.com:amr"
      values   = ["authenticated"]
    }
  }
}

resource "aws_iam_role" "authenticated" {
  name               = "${var.application}-authenticated-${terraform.workspace}"
  assume_role_policy = data.aws_iam_policy_document.cognito_authenticated_role.json
}

data "aws_iam_policy_document" "cognito_authenticated_role_policy" {
  statement {
    actions = [
      "s3:DeleteObject",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject",
      "s3:PutObjectAcl",
    ]

    resources = [
      "${aws_s3_bucket.videos.arn}",
      "${aws_s3_bucket.videos.arn}/*",
    ]
  }
}

resource "aws_iam_role_policy" "authenticated" {
  name   = "${var.application}-authenticated-${terraform.workspace}"
  role   = aws_iam_role.authenticated.id
  policy = data.aws_iam_policy_document.cognito_authenticated_role_policy.json
}

resource "aws_cognito_identity_pool_roles_attachment" "this" {
  identity_pool_id = aws_cognito_identity_pool.this.id

  roles = {
    "authenticated" = aws_iam_role.authenticated.arn
  }
}
