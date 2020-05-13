# Triggermesh Bridges

Bridges are a seamless way of bundling kubernetes objects for deployment and readiness.

It is a single manifest controlled by the [TriggerFlow](https://github.com/triggermesh/triggerflow) operator that contains an arbitrary number of components, each of which needs to be a valid kubernetes object.

## Features

TriggerFlow allows for each Brige to:

- Deploy all items in the bundle
- Check status for those components that comply with [status conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties)
- Reconcyle components to make sure they match their definition at the Bridge

Bridges are focused at TriggerMesh use cases which heavily involve Knative Eventing, but can be used with any kubernetes object.

### Support

We would love your feedback on those bridges so do not hesitate to let us know what is wrong and how we could improve them, just file an [issue](https://github.com/triggermesh/bridges/issues/new)

### Code of Conduct

This plugin is by no means part of [CNCF](https://www.cncf.io/) but we abide by its [code of conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md)
