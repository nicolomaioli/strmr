variable "application" {
  description = "the name of the application"
  type        = string
  default     = "strmr"
}

variable "region" {
  description = "the region we are applying to"
  type        = string
  default     = "eu-west-1"
}

variable "cloudfront" {
  description = "cloudfront distribution configuration"
  type = object({
    aliases             = list(string)
    whitelist_locations = list(string)
    acm_certificate_arn = string
  })
}

variable "route53" {
  description = "route53 configuration"
  type = object({
    zone      = string
    subdomain = string
  })
}
