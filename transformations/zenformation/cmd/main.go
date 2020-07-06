package main

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	trans "github.com/triggermesh/zenformation/pkg"
)

func main() {

	// The default client is HTTP.
	z := &trans.Zenformation{}
	c, err := cloudevents.NewDefaultClient()
	z.CeClient = c
	if err != nil {
		fmt.Println("failed to create client, %v", err)
		z.ErrorHandler(err)
	}
	ctx := context.Context(context.Background())

	z.Start(ctx.Done())

}
