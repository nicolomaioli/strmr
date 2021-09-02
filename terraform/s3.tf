resource "aws_s3_bucket" "videos" {
  bucket = "${var.application}-videos-${terraform.workspace}"
  acl    = "private"

  cors_rule {
    allowed_headers = ["*"]
    allowed_origins = ["*"]
    max_age_seconds = 3000
    allowed_methods = [
      "HEAD",
      "GET",
      "PUT",
      "POST",
      "DELETE",
    ]
    expose_headers = [
      "x-amz-server-side-encryption",
      "x-amz-request-id",
      "x-amz-id-2",
      "ETag"
    ]
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
