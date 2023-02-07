package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"iScore-api/api"
)

func main() {
	lambda.Start(api.Run())
}
