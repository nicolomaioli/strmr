resource "aws_dynamodb_table" "this" {
  name = "${var.application}-videos-${terraform.workspace}"

  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "UserId"
  range_key    = "CreatedAt"

  attribute {
    name = "UserId"
    type = "S"
  }

  attribute {
    name = "VideoId"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  attribute {
    name = "Type"
    type = "S"
  }

  global_secondary_index {
    name            = "VideoIndex"
    hash_key        = "VideoId"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "TypeCreatedAtIndex"
    hash_key        = "Type"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
