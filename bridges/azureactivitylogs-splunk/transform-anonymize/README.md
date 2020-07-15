# Azure Activity Logs data anonymization

This code demonstrates how to transform the data contained in CloudEvents originating from the [TriggerMesh Azure
Activity Logs event source][aalsource-docs] in order to anonymize people's identity and IP addresses.

The example is written in Go and released as a container image at
`gcr.io/triggermesh/bridge-examples/azureactivitylogs-splunk-transform-anonymize`.

[aalsource-docs]: https://docs.triggermesh.io/sources/azureactivitylogs/
