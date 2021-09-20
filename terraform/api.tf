resource "aws_apigatewayv2_api" "this" {
  name                         = "${var.application}-api-${terraform.workspace}"
  protocol_type                = "HTTP"
  disable_execute_api_endpoint = true

  cors_configuration {
    allow_origins = ["*"]
    allow_methods = ["*"]
    allow_headers = [
      "x-amz-date",
      "x-api-key",
      "x-amz-security-token",
      "x-amz-user-agent",
      "Content-Type",
      "Authorization",
    ]
  }

  tags = local.tags
}

resource "aws_apigatewayv2_domain_name" "this" {
  domain_name = "${var.api_gateway.subdomain}.${var.route53.zone}"

  domain_name_configuration {
    certificate_arn = var.acm_certificate_arn.regional
    endpoint_type   = "REGIONAL"
    security_policy = "TLS_1_2"
  }

  tags = local.tags
}
