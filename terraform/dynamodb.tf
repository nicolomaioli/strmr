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
    type = "N"
  }

  attribute {
    name = "JobStatus"
    type = "S"
  }

  global_secondary_index {
    name            = "UsernameCreatedAtIndex"
    hash_key        = "Username"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "JobStatusCreatedAtIndex"
    hash_key        = "JobStatus"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  tags = local.tags
}
