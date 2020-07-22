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
	"net/http"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
)

const (
	ticketCreatedEventType = "com.zendesk.ticket.created"
	tagCreateEventType     = "com.zendesk.tag.create"

	transformTypeExtension = "transformtype"
)

const eventSource = "io.triggermesh.transformations.zendesk-sentiment-tag"

const connTimeout = 20 * time.Second

// Receiver runs a CloudEvents receiver.
type Receiver struct {
	ceClient   cloudevents.Client
	compClient comprehendiface.ComprehendAPI

	lang string
}

// handleEvent accepts a CloudEvent and responds to the sender.
func (recv *Receiver) handleEvent(ctx context.Context, e cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if typ := e.Type(); typ != ticketCreatedEventType {
		return nil, cloudevents.NewHTTPResult(http.StatusAccepted, "Ignoring event type %q", typ)
	}

	var req Request
	if err := e.DataAs(&req); err != nil {
		log.Print(err)
		return nil, cloudevents.NewHTTPResult(http.StatusBadRequest,
			"Received data can not be deserialized to Zendesk event: %s", err)
	}

	log.Printf("Processing event from source %q", e.Source())

	sentiment, err := recv.analyzeSentiment(ctx, req.Ticket.Description)
	if err != nil {
		log.Printf("Failed to analyze ticket sentiment: %s", err)
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError,
			"Failed to analyze ticket sentiment: %s", err)
	}

	log.Print("Comprehend responded with sentiment: ", *sentiment)

	resp := &Response{
		ID:  req.Ticket.ID,
		Tag: *sentiment.Sentiment,
	}

	event := cloudevents.NewEvent(cloudevents.VersionV1)
	event.SetType(tagCreateEventType)
	event.SetSource(eventSource)
	event.SetTime(time.Now())
	event.SetExtension(transformTypeExtension, e.Type())
	err = event.SetData(cloudevents.ApplicationJSON, resp)
	if err != nil {
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError, "Failed to set event data: %s", err)
	}

	return &event, cloudevents.NewHTTPResult(http.StatusCreated, "Event transformed")
}

// analyzeSentiment requests a sentiment analysis of the given ticket
// description from AWS Comprehend.
func (recv *Receiver) analyzeSentiment(ctx context.Context, ticketDesc string) (*comprehend.DetectSentimentOutput, error) {
	var dSI comprehend.DetectSentimentInput
	dSI.SetLanguageCode(recv.lang)
	dSI.SetText(ticketDesc)

	connCtx, cancel := context.WithTimeout(ctx, connTimeout)
	defer cancel()
	return recv.compClient.DetectSentimentWithContext(connCtx, &dSI)
}
