# Slack to Elasticsearch

This sample bridge indexes message written to a channel the Slack bot is a member of.

## Components

This bridge uses

- Slack source as the producer of events
- A broker that receives events from slack
- A trigger that connects the broker with the target
- An Elasticsearch target

## Parametrization

Customization of URL parameter for Elasticsearch and secret creation for both Elasticsearch and Slack is required.

- Refer to [Slack docs](../../docs/sources/slack.md) for configuring Slack
- Refer to [Elasticsearch docs](../../docs/targets/elasticsearch.md) for configuring Elasticsearch.
