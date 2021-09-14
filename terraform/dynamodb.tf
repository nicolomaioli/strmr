resource "aws_dynamodb_table" "this" {
  name = "${var.application}-videos-${terraform.workspace}"

  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ID"

  attribute {
    name = "Username"
    type = "S"
  }

  attribute {
    name = "ID"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  attribute {
    name = "Status"
    type = "S"
  }

  global_secondary_index {
    name            = "UsernameCreatedAtIndex"
    hash_key        = "Username"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "StatusCreatedAtIndex"
    hash_key        = "Status"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
