resource "aws_dynamodb_table" "this" {
  name = "${var.application}-videos-${terraform.workspace}"

  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "Username"
  range_key    = "CreatedAt"

  attribute {
    name = "Username"
    type = "S"
  }

  attribute {
    name = "VideoID"
    type = "S"
  }

  attribute {
    name = "CreatedAt"
    type = "S"
  }

  attribute {
    name = "FileType"
    type = "S"
  }

  global_secondary_index {
    name            = "VideoIndex"
    hash_key        = "VideoID"
    projection_type = "ALL"
  }

  global_secondary_index {
    name            = "TypeCreatedAtIndex"
    hash_key        = "FileType"
    range_key       = "CreatedAt"
    projection_type = "ALL"
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
