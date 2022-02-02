terraform {
  required_version = "~> 1.1.2"

  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 3.65"
    }
  }

  backend "s3" {
    bucket = "gingerbreadtemplate-terraform"
    key = "prod/terraform.tfstate"
    region = "eu-west-2"
  }
}

provider "aws" {
  region = "eu-west-2"
  max_retries = 5
}

provider "aws" {
  alias = "acm_provider"
  region = "us-east-1"
  max_retries = 5
}