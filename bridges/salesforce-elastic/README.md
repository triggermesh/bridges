# Salesforce to Elasticsearch

This bridge template connects Salesforce with Elasticsearch, sending any type of event from the former into
a single index. Two files need to be customized to get the bridge working:

- Secrets: contain sensitive authentication data for both ends.
- Bridge: contains Triggermesh blocks whose elements need to be customized to your environment.

## Prerequisites

- Salesforce account configured as [stated by our guide](https://docs.triggermesh.io/sources/salesforce/).
- Elasticsearch cluster connection and credentials data.

### Secrets

The secrets file contains placeholders for

- Salesforce signing cert key, paired with the certificate configured at Salesforce's Connected App.
- Elasticsearch password. Other authentication mechanisms are also supported.

### Bridge

The bridge is composed of 4 elements

- Broker: is the subscribable bus where events will be sent.

- Salesforce source:

  Replace channel with the topic for the events that the source will be listening to. ClientID needs to also be informed along with the reference to the secret that was created at the previous step.

  The sink element at this block references the broker that will receive the events.

- Elasticsearch target:

  Replace URL (including port), username and index to store events. Refer to [Elasticsearch docs](../../docs/targets/elasticsearch.md) for configuring this target.

- Trigger:

  Contains a reference to the broker to subscribe to, and a reference to the Elasticsearch target where all unfiltered events will be posted.


