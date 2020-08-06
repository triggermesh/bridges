# Kafka to event-display service

This sample logs messages sent throught Kafka into an `event-display` service.

## Components

This bridge uses

- Kafka source as the producer of events
- A broker that receives events from Kafka
- A trigger that connects the broker with the service
- An event display service

## Parametrization

Kafka source is a community component, refer to [Kafka repo](https://github.com/knative/eventing-contrib/tree/master/kafka/source) for configuration.
