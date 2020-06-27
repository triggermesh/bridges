# Notes on the Oracle Source Bridges

These bridges are designed to connect the Oracle DB Event Source to multiple
targets.  Focusing on the concept of an e-commerce platform, the scenario is
that there are two types of transactions to watch for from the single source
  * New orders coming in
  * Running out of products

In one scenario, a serverless function will be invoked. In the other scenario,
a message will be posted to either Slack or Zendesk.

There is a single [translators](translators/) directory that contain slightly
different translators depending on whether the zendesk or slack target is used.
Both are built as docker images using the `build.sh` script.  This is where the
reading of the database change event is made and the creation of the target
events are produced.

## Prerequisites

To utilize any of the sources, a few attributes will need to be tweaked.  This
is in addition to the required secrets for each known target.

### SinkBinding

There is a namespace that is hardcoded, and must match the target namespace the
bridge will be deployed into.  Without setting this, the SinkBinding will attempt
to bind to the service running in the same namespace as the EveryBridge controller.

### Slack

A default channel of `test` is used for the scenario, but should reflect a channel
that already exists

### Zendesk

This requires an active Zendesk subdomain and point-of-contact that is registered
within Zendesk.