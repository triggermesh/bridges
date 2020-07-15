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
	"encoding/json"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const (
	activityLogEventType   = "com.microsoft.azure.monitor.activity-log"
	transformTypeExtension = "transformtype"
)

// Handler runs a CloudEvents receiver.
type Handler struct {
	cli cloudevents.Client
}

// NewHandler returns a new Handler for the given CloudEvents client.
func NewHandler(c cloudevents.Client) *Handler {
	return &Handler{
		cli: c,
	}
}

// Run starts the handler and blocks until it returns.
func (h *Handler) Run(ctx context.Context) error {
	return h.cli.StartReceiver(ctx, h.receive)
}

// receive implements the handler's receive logic.
func (h *Handler) receive(e cloudevents.Event) (*cloudevents.Event, cloudevents.Result) {
	if typ := e.Type(); typ != activityLogEventType {
		return nil, cloudevents.NewHTTPResult(http.StatusAccepted, "Ignoring event type %q", typ)
	}

	var eventData EventRecords
	if err := json.Unmarshal(e.Data(), &eventData); err != nil {
		return nil, cloudevents.NewHTTPResult(http.StatusBadRequest,
			"Received data can not be deserialized to Activity Log records: %s", err)
	}

	eventData.Anonymize()

	e.SetExtension(transformTypeExtension, e.Type()) // this transformation does not alter the event type
	err := e.SetData(cloudevents.ApplicationJSON, &eventData)
	if err != nil {
		return nil, cloudevents.NewHTTPResult(http.StatusInternalServerError, "Failed to set event data: %s", err)
	}

	return &e, cloudevents.NewHTTPResult(http.StatusCreated, "Event transformed")
}
