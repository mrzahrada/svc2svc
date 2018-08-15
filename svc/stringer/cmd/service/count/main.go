package main

import (
	"github.com/mrzahrada/aws-lambda-go/lambda"
	"github.com/mrzahrada/svc2svc/svc/stringer/internal/service"
)

func main() {
	lambda.Start(service.CountHandler())
}
