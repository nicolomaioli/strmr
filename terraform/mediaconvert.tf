resource "aws_media_convert_queue" "this" {
  name = "${var.application}-${terraform.workspace}"

  pricing_plan = "ON_DEMAND"
  status       = "ACTIVE"

  tags = {
    Application = var.application
    Environment = terraform.workspace
    Terraform   = true
  }
}
