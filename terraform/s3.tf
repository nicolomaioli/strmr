resource "aws_s3_bucket" "videos" {
  bucket = "${var.application}-videos-${terraform.workspace}"
  acl    = "private"

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = [
      "HEAD",
      "GET",
      "PUT",
      "POST",
      "DELETE",
    ]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
