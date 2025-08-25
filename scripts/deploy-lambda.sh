#!/bin/bash
set -e  # exit immediately if a command fails

# -----------------------------
# Configuration
# -----------------------------
FUNCTION_NAME="fantasy-media-page"
GO_VERSION="1.25"
OUTPUT_BINARY="bootstrap"
ZIP_FILE="function.zip"
AWS_REGION="us-east-1"

# -----------------------------
# Build Go binary for Lambda
# -----------------------------
echo "Building Go binary..."
GOOS=linux GOARCH=arm64 go build -o $OUTPUT_BINARY main.go

# -----------------------------
# Package into zip
# -----------------------------
echo "Packaging binary into $ZIP_FILE..."
zip -r $ZIP_FILE $OUTPUT_BINARY

# -----------------------------
# Deploy to AWS Lambda
# -----------------------------
echo "Deploying to Lambda function: $FUNCTION_NAME..."
aws lambda update-function-code \
  --function-name $FUNCTION_NAME \
  --zip-file fileb://$ZIP_FILE \
  --region $AWS_REGION

echo "Deployment complete... Cleaning Up"

rm bootstrap function.zip
