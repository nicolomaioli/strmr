resource "aws_media_convert_queue" "this" {
  name = "${var.application}-${terraform.workspace}"

  pricing_plan = "ON_DEMAND"
  status       = "ACTIVE"

  tags = local.tags
}

data "aws_iam_policy_document" "mediaconvert_role_policy" {
  statement {
    actions = [
      "s3:Get*",
      "s3:List*",
    ]

    resources = [
      "${aws_s3_bucket.videos.arn}/*",
    ]
  }

  statement {
    actions = [
      "s3:Get*",
      "s3:List*",
      "s3:Put*",
    ]

    resources = [
      "${aws_s3_bucket.vod.arn}/*",
    ]
  }
}

resource "aws_iam_role_policy" "mediaconvert" {
  name   = "${var.application}-mediaconvert-${terraform.workspace}"
  role   = aws_iam_role.mediaconvert.id
  policy = data.aws_iam_policy_document.mediaconvert_role_policy.json
}

data "aws_iam_policy_document" "mediaconvert_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["mediaconvert.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "mediaconvert" {
  name               = "${var.application}-mediaconvert-${terraform.workspace}"
  assume_role_policy = data.aws_iam_policy_document.mediaconvert_role.json
}
