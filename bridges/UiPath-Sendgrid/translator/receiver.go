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
	"fmt"
	"log"
	"net/http"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
)

const (
	uiPathEventType               = "io.triggermesh.uipath"
	uiPathQueueItemAddedEventType = "queueItem.added"
	sendgridEmailEventType        = "io.triggermesh.sendgrid.email.send"

	transformTypeExtension = "transformtype"
)

const eventSource = "io.triggermesh.transformations.ui-sendgrid"

const connTimeout = 20 * time.Second

// Receiver runs a CloudEvents receiver.
type Receiver struct {
	ceClient   cloudevents.Client
	compClient comprehendiface.ComprehendAPI

	lang string
}

func (recv *Receiver) handleQueueItemAdded(req Request) (Response, error) {
	var response Response
	for _, d := range req.QueueItems {
		if d.SpecificContent.Email != "" {
			response.ToEmail = d.SpecificContent.Email
		}

		if d.SpecificContent.Invoice != "" {
			response.Message = d.SpecificContent.Invoice
		}

		if d.SpecificContent.CustomerName != "" {
			response.Subject = "Hello: " + d.SpecificContent.CustomerName + ". Your Triggermesh Invoice is Avalible."
		}
	}

	if response.ToEmail == "" {
		return response, fmt.Errorf("no Email provided")
	}

	return response, nil
}

// handleEvent accepts a CloudEvent and responds to the sender.
func (recv *Receiver) handleEvent(ctx context.Context, e cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if typ := e.Type(); typ != uiPathEventType {
		return nil, cloudevents.NewHTTPResult(http.StatusAccepted, "Ignoring event type %q", typ)
	}

	var req Request
	if err := e.DataAs(&req); err != nil {
		log.Print(err)
		return nil, cloudevents.NewHTTPResult(http.StatusBadRequest,
			"Received data can not be deserialized to Sendgrid event: %s", err)
	}

	log.Printf("Processing event from source %q", e.Source())

	var resp Response
	switch typ := req.Type; typ {
	case uiPathQueueItemAddedEventType:
		x, err := recv.handleQueueItemAdded(req)
		if err != nil {
			fmt.Errorf("an erro proccessing the added queue item: ")
			fmt.Println(err)
			return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError,
				"Failed to analyze: %s", err)
		}
		resp = x
	default:
		fmt.Errorf("cannot process specified UiEvent Type (not CE)")
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError,
			"Failed to analyze: %s", nil)
	}

	// check that a message exists in the response
	if resp.Message == "" {
		fmt.Errorf("Did not get a message in transformation: ")
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError,
			"Failed to analyze: %s", nil)

	}

	event := cloudevents.NewEvent(cloudevents.VersionV1)
	event.SetType(sendgridEmailEventType)
	event.SetSource(eventSource)
	event.SetTime(time.Now())
	event.SetID("hi-mom")
	event.SetExtension(transformTypeExtension, e.Type())
	err := event.SetData(cloudevents.ApplicationJSON, resp)
	if err != nil {
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError, "Failed to set event data: %s", err)
	}

	fmt.Println("sent event:")
	fmt.Println(event.String())

	return &event, cloudevents.NewHTTPResult(http.StatusCreated, "Event transformed")
}
