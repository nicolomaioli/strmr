output "user_pool_id" {
  value = aws_cognito_user_pool.this.id
}

output "user_pool_client_id" {
  value = aws_cognito_user_pool_client.this.id
}

output "identity_pool_id" {
  value = aws_cognito_identity_pool.this.id
}

output "s3_videos_arn" {
  value = aws_s3_bucket.videos.arn
}

output "dynamodb_table_name" {
  value = aws_dynamodb_table.this.name
}

output "dynamodb_table_arn" {
  value = aws_dynamodb_table.this.arn
}

output "mediaconvert_queue_arn" {
  value = aws_media_convert_queue.this.arn
}

output "mediaconvert_queue_id" {
  value = aws_media_convert_queue.this.id
}
