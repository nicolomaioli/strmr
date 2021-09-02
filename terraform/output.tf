output "user_pool_id" {
  value = aws_cognito_user_pool.this.id
}

output "user_pool_client_id" {
  value = aws_cognito_user_pool_client.this.id
}

output "s3_videos_arn" {
  value = aws_s3_bucket.videos.arn
}
