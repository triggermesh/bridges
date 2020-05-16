# Cronjob to Elasticsearch

This sample bridge indexes a JSON message produced regulary from a cronjob.

## Components

This bridge uses

- The pingsource as the producer of events
- A broker that receives events from pingsource
- A trigger that connects the broker with the target
- An Elasticsearch target

## Parametrization

- Refer to [Pingsource docs](https://github.com/knative/docs/tree/master/docs/eventing/samples/ping-source) for the JSON being produced.
- Refer to [Elasticsearch docs](../../docs/targets/elasticsearch.md) for customization of credentials and index.
