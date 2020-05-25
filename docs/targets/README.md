# Targets for Bridges

Targets are custom objects available in TriggerMesh hosted EveryBridge. Currently you can use:

* [Elasticsearch](./elasticsearch.md) via the kind `ElasticSearchTarget`
* [Confluent](./confluent.md) via the kind `ConfluentTarget`
* [SendGrid](./sendgrid.md) via the kind `SendgridTarget`
* Tekton via the kind `TektonTarget`
* [Twilio](./twilio.md) via the kind `TwilioTarget`
* AWS SQS via the kind `AWSTarget`
* AWS Lambda via the kind `AWSTarget`
* AWS Kinesis via the kind `AWSTarget`
* AWS SNS via the kind `AWSTarget`

Each Target has its own `Spec` and may need security credentials containing API keys.

## Implementation

Technically Targets are Kubernetes CRDs managed by a controller. Targets are adressable and can receive events.
