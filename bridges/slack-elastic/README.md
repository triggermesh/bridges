# Slack to Elasticsearch

This sample bridge indexes message written to a channel the Slack bot is a member of.

## Components

This bridge uses

- Slack source as the producer of events
- A broker that receives events from slack
- A trigger that connects the broker with the target
- An Elasticsearch target

## Parametrization

Customization of URL parameter and secret for Elasticsearch authentication are required.  Optionally the Slack source can also be customized to verify received messages signature.

- Refer to [Slack docs](../../docs/sources/slack.md) for configuring Slack
- Refer to [Elasticsearch docs](../../docs/targets/elasticsearch.md) for configuring Elasticsearch.
