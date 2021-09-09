package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handle)
}

func handle() {
	fmt.Println("hello world!")
}
