resource "aws_cloudfront_origin_access_identity" "video-out" {
  comment = "Origin access identity for the video-out bucket"
}

data "aws_iam_policy_document" "video-out" {
  statement {
    actions   = ["s3:GetObject"]
    resources = ["${aws_s3_bucket.video-out.arn}/*"]

    principals {
      type        = "AWS"
      identifiers = [aws_cloudfront_origin_access_identity.video-out.iam_arn]
    }
  }
}

resource "aws_s3_bucket_policy" "video-out" {
  bucket = aws_s3_bucket.video-out.id
  policy = data.aws_iam_policy_document.video-out.json
}

resource "aws_cloudfront_distribution" "video-out" {
  origin {
    domain_name = aws_s3_bucket.video-out.bucket_regional_domain_name
    origin_id   = aws_s3_bucket.video-out.id

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.video-out.cloudfront_access_identity_path
    }
  }

  enabled         = true
  is_ipv6_enabled = true
  aliases         = var.cloudfront.aliases

  viewer_certificate {
    acm_certificate_arn = var.acm_certificate_arn.edge
    ssl_support_method  = "sni-only"
  }

  restrictions {
    geo_restriction {
      restriction_type = "whitelist"
      locations        = var.cloudfront.whitelist_locations
    }
  }

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD", "OPTIONS"]
    target_origin_id = aws_s3_bucket.video-out.id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }

      headers = [
        "Access-Control-Request-Headers",
        "Access-Control-Request-Method",
        "Origin",
      ]
    }

    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
    viewer_protocol_policy = "redirect-to-https"
  }

  tags = local.tags
}
