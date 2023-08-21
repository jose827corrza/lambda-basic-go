package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleRequest(ctx context.Context, data MyStruct) (string, error) {
	return fmt.Sprintf("Hello %s!", data.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
