# Triggermesh Bridges ![](./everybridge.png)

Bridges are the core object of TriggerMesh's EveryBridge offering available in Beta for free in the [TriggerMesh Cloud](https://cloud.triggermesh.io).

They represent a seamless way of building event flows made of event sources and event targets. In this repository you will find some sample Bridges that can be used in the TriggerMesh Cloud.

A `Bridge` is a single manifest that contains an arbitrary number of components, each of which needs to be a valid kubernetes object. It was inspired by the [Application CRD](https://github.com/kubernetes-sigs/application). TriggerFlow is the internal controller which manages `Bridges`.

> A note on Open Source. Currently the Bridge controller is not open source, it is still under heavy development. We plan to and to some extent already are contributing the ideas behind Bridges and Targets to the upstream Knative community. Stay tune for a fast evolving space...

## Features

TriggerFlow allows for each Bridge to:

- Deploy all objects in the manifest
- Check status for those components that comply with [status conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties)
- Reconcile components to make sure they match their definition at the Bridge
- Contain any event Sources
- Contain any event Targets
- Contain any Knative serving/eventing objects

## Secrets

> For Bridges to be able to connect to SaaS offerings and Public Clouds services, you need to store your own credentials as a Kubernetes secret in the TriggerMesh Cloud. Without those credentials, the Bridges will not work.

## Catalog

The following sample Bridges are available and can be deployed on [https://cloud.triggermesh.io](https://cloud.triggermesh.io) .

| From | To | Manifests |
|------|----|---|
|GitHub|Display|[yaml](./bridges/github-display/)|
|GitHub|Twilio|[yaml](./bridges/github-twilio/)|
|GitLab|Display|[yaml](./bridges/gitlab-display/)|
|GitLab|AWS SQS,Kinesis,SNS|[yaml](./bridges/gitlab-aws/)|
|AWS Kinesis|Confluent|[yaml](./bridges/kinesis-confluent/)|
|AWS SQS| SendGrid|[yaml](./bridges/sqs-sendgrid/)|
|AWS SQS| Twilio|[yaml](./bridges/sqs-twilio/)|
|CronJob| ElasticSearch|[yaml](./bridges/cronjob-elastic/)|
|CronJob| Confluent|[yaml](./bridges/cronjob-confluent/)|
|CronJob| Tekton|[yaml](./bridges/cronjob-tekton/)|
|Salesforce| Elasticsearch|[yaml](./bridges/salesforce-elastic/)|

Any combination from a documented event [Source](./docs/sources/README.md) and a documented event [Target](./docs/targets/README.md) can be made.

## Contact Us

If you need dedicated help with those Bridges or want to discuss on-premises use, feel free to contact TriggerMesh **info@triggermesh.com**

## Support

We would love your feedback on those bridges so do not hesitate to let us know what is wrong and how we could improve them, just file an [issue](https://github.com/triggermesh/bridges/issues/new)

## Code of Conduct

These Bridges are by no means part of [CNCF](https://www.cncf.io/) but we abide by its [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
