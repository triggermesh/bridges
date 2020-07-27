# Google Storage to AWS EventBridge

This bridge sends events from Google storage to AWS Eventbridge

## Configure a Google Storage Bucket

In order to get notifications, see [Google Storage](https://github.com/triggermesh/bringyourown/tree/master/sources/python/googlestorage) source.

## Apply the Bridge

```
kubectl apply -f google-eventbridge.yaml
``` 

## Configure AWS EventBridge to Act on events

See [https://docs.triggermesh.io/targets/awseventbridge/](https://docs.triggermesh.io/targets/awseventbridge/)
