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

import eventhub "github.com/Azure/azure-event-hubs-go/v3"

const anonString = "*****"

// EventRecords is an aggregate of Azure Activity Log events.
// The schema for Activity Logs originating from Event Hubs is documented at
// https://docs.microsoft.com/en-us/azure/azure-monitor/platform/activity-log-schema#schema-from-storage-account-and-event-hubs.
type EventRecords struct {
	Data *struct {
		Records []map[string]interface{} `json:"records,omitempty"`
	} `json:"Data,omitempty"`

	*eventhub.Event
}

// Anonymize performs an in-place anonymization of the events' data contained
// in the EventRecords.
func (r *EventRecords) Anonymize() {
	if r.Data == nil {
		return
	}

	for _, r := range r.Data.Records {
		anonymize(r)
	}
}

func anonymize(e map[string]interface{}) {
	for _, f := range anonFields {
		if _, isSet := e[f]; isSet {
			e[f] = anonString
		}
	}

	identityIface, isSet := e["identity"]
	if !isSet {
		return
	}
	identity, ok := identityIface.(map[string]interface{})
	if !ok {
		return
	}

	identityClaimsIface, isSet := identity["claims"]
	if !isSet {
		return
	}
	identityClaims, ok := identityClaimsIface.(map[string]interface{})
	if !ok {
		return
	}

	for _, f := range anonIdentityClaimFields {
		if _, isSet := identityClaims[f]; isSet {
			identityClaims[f] = anonString
		}
	}
}

var anonFields = []string{
	"callerIpAddress",
}

var anonIdentityClaimFields = []string{
	"name",
	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/givenname",
	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name",
	"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/surname",
	"ipaddr",
}
