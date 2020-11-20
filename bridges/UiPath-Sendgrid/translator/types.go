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

// Request is the structure of the event we expect to receive.
type Request struct {
	Type       string `json:"Type"`
	QueueItems []struct {
		ID              int64 `json:"id"`
		SpecificContent struct {
			Invoice      string `json:"invoice"`
			Email        string `json:"email"`
			CustomerName string `json:"customerName"`
		} `json:"SpecificContent"`
	} `json:"QueueItems"`
}

// Response is the structure of the event we send in response to requests.
type Response struct {
	Message string `json:"message"`
	ToEmail string `json:"toEmail"`
	Subject string `json:"subject"`
}
