resource "aws_s3_bucket" "videos" {
  bucket = "${var.application}-videos-${terraform.workspace}"
  acl    = "private"

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
