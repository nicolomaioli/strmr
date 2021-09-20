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

variable "acm_certificate_arn" {
  description = "arn of acm certificates"
  type = object({
    edge     = string
    regional = string
  })
}

variable "cloudfront" {
  description = "cloudfront distribution configuration"
  type = object({
    aliases             = list(string)
    whitelist_locations = list(string)
    subdomain           = string
  })
}

variable "api_gateway" {
  description = "api gateway v2 configuration"
  type = object({
    subdomain = string
  })
}

variable "route53" {
  description = "route53 configuration"
  type = object({
    zone = string
  })
}

variable "route53_record_types" {
  description = "a list of route53 record types to be created"
  type        = list(string)
}
