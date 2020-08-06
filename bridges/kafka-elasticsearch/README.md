# Kafka to Elasticsearch

This sample Bridge indexes messages sent through Kafka into an Elasticsearch cluster.

## Components

This bridge uses

- Kafka source as the producer of events
- A broker that receives events from Kafka
- A trigger that connects the broker with the target
- An Elasticsearch target

## Parametrization

Customization of URL parameter and secret for Elasticsearch authentication are required. Kafka source is a community component that  will also need to be configured.

- Refer to [Kafka repo](https://github.com/knative/eventing-contrib/tree/master/kafka/source) for configuring the Kafka instance
- Refer to [Elasticsearch docs](../../docs/targets/elasticsearch.md) for configuring Elasticsearch.
