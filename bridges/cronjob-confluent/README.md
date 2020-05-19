# Cronjob to Confluent

This sample bridge sends a JSON message produced regulary from a cronjob to confluent.

## Components

This bridge uses

- The pingsource as the producer of events
- A broker that receives events from pingsource
- A trigger that connects the broker with the target
- A Confluent target

## Parametrization

Customization of URL and secret creation for Confluent is required.

- Refer to [Pingsource docs](https://github.com/knative/docs/tree/master/docs/eventing/samples/ping-source) for the JSON being produced.
- Refer to [Confluent docs](../../docs/targets/confluent.md) for customization of credentials and topic.
