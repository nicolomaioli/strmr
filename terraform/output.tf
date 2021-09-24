output "user_pool_id" {
  value = aws_cognito_user_pool.this.id
}

output "user_pool_client_id" {
  value = aws_cognito_user_pool_client.this.id
}

output "identity_pool_id" {
  value = aws_cognito_identity_pool.this.id
}

output "api_gateway_id" {
  value = aws_apigatewayv2_api.this.id
}

output "api_gateway_arn" {
  value = aws_apigatewayv2_api.this.arn
}

output "s3_video_in_arn" {
  value = aws_s3_bucket.video-in.arn
}

output "s3_video_in_name" {
  value = aws_s3_bucket.video-in.id
}

output "s3_video_out_arn" {
  value = aws_s3_bucket.video-out.arn
}

output "s3_video_out_id" {
  value = aws_s3_bucket.video-out.id
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

output "mediaconvert_role_arn" {
  value = aws_iam_role.mediaconvert.arn
}

output "mediaconvert_role_name" {
  value = aws_iam_role.mediaconvert.name
}
