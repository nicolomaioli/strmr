data "aws_route53_zone" "this" {
  name = var.route53.zone
}

resource "aws_route53_record" "cf" {
  for_each = toset(var.route53_record_types)
  type     = each.key
  zone_id  = data.aws_route53_zone.this.zone_id
  name     = "${var.cloudfront.subdomain}.${var.route53.zone}"

  alias {
    name                   = aws_cloudfront_distribution.vod.domain_name
    zone_id                = aws_cloudfront_distribution.vod.hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "api" {
  for_each = toset(var.route53_record_types)
  type     = each.key
  zone_id  = data.aws_route53_zone.this.zone_id
  name     = aws_apigatewayv2_domain_name.this.domain_name

  alias {
    name                   = aws_apigatewayv2_domain_name.this.domain_name_configuration[0].target_domain_name
    zone_id                = aws_apigatewayv2_domain_name.this.domain_name_configuration[0].hosted_zone_id
    evaluate_target_health = false
  }
}
