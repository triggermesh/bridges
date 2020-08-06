# Slack to Confluent Kafka cluster

This sample bridge pushes Slack messages from a Bot to a Confluent Kafka cluster topic.

## Components

This bridge uses

- Slack source as the producer of events.
- A broker that receives events from Slack.
- A Confluent Target instance.
- A trigger that connects the broker with the target.

## Parameterization

Customization of Confluent's instance is required. Optionally the Slack source can also be customized to verify received messages signature.

- Refer to [Slack docs](../../docs/sources/slack.md) for configuring Slack.
- Refer to [Confluent docs](../../docs/targets/confluent.md) for configuring Confluent.
