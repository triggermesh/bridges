# Flow Stacks

Flow Stacks are a seamless way of bundling kubernetes objects for deployment and readiness.

It is a single manifest controlled by the [TriggerFlow](https://github.com/triggermesh/triggerflow) operator that contains an arbitrary number of components, each of which needs to be a valid kubernetes object.

## Features

TriggerFlow allows for each Stack to:

- Deploy all items in the bundle
- Check status for those components that comply with [status conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties)
- Reconcyle components to make sure they match their definition at the Stack

Flow Stacks are focused at TriggerMesh use cases which heavily involve Knative Eventing, but can be used with any kubernetes object.