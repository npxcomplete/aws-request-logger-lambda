package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(func(ctx context.Context, event interface{}) (resp interface{}, err error) {
		// all events arrive as JSON objects
		bs, err := json.Marshal(event)
		if err != nil {
			fmt.Println("Error serializing json event.")
			return
		}

		// The entire purpose of this lambda func
		fmt.Println(string(bs))

		// play nice with api gateway / http clients
		err = nil
		resp = events.APIGatewayProxyResponse{
			StatusCode:        200,
			Body:              "{\"goodbye\": \"world\"}",
		}
		return
	})
}
