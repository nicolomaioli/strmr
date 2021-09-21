resource "aws_s3_bucket" "video-in" {
  bucket = "${var.application}-${terraform.workspace}-video-in"
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

  tags = local.tags
}

resource "aws_s3_bucket" "video-out" {
  bucket = "${var.application}-${terraform.workspace}-video-out"
  acl    = "private"

  cors_rule {
    allowed_headers = ["*"]
    allowed_origins = ["*"]
    max_age_seconds = 3000
    allowed_methods = [
      "GET",
    ]
  }

  tags = local.tags
}
