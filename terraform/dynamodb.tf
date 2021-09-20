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
    name = "JobStatus"
    type = "S"
  }

  global_secondary_index {
    name            = "UsernameIndex"
    hash_key        = "Username"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "JobStatusIndex"
    hash_key        = "JobStatus"
    projection_type = "ALL"
  }

  tags = local.tags
}
