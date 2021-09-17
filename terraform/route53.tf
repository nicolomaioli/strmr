data "aws_route53_zone" "this" {
  name = var.route53.zone
}

resource "aws_route53_record" "a" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "${var.route53.subdomain}.${var.route53.zone}"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.vod.domain_name
    zone_id                = aws_cloudfront_distribution.vod.hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "aaaa" {
  zone_id = data.aws_route53_zone.this.zone_id
  name    = "${var.route53.subdomain}.${var.route53.zone}"
  type    = "AAAA"

  alias {
    name                   = aws_cloudfront_distribution.vod.domain_name
    zone_id                = aws_cloudfront_distribution.vod.hosted_zone_id
    evaluate_target_health = false
  }
}
