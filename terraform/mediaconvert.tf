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

resource "aws_iam_role" "mediaconvert" {
  name = "${var.application}-mediaconvert-${terraform.workspace}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "mediaconvert.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "mediaconvert" {
  name = "${var.application}-mediaconvert-${terraform.workspace}"
  role = aws_iam_role.mediaconvert.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:Get*",
        "s3:List*",
        "s3:Put*"
      ],
      "Resource": [
        "${aws_s3_bucket.videos.arn}/*"
      ]
    }
  ]
}
EOF
}
