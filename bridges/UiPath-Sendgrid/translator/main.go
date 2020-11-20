/*
Copyright (c) 2020 TriggerMesh Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func main() {
	ceCli, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("Unable to create CloudEvents client: %s", err)
	}

	r := Receiver{
		ceClient: ceCli,
	}

	log.Print("Running transformation receiver")

	ctx := context.Background()
	if err := r.ceClient.StartReceiver(ctx, r.handleEvent); err != nil {
		log.Fatal(err)
	}
}
