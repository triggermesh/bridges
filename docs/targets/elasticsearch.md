# Elasticsearch event target

This event target integrates with Elasticsearch using received Cloud Event messages to index documents.

## Contents

- [Elasticsearch event target](#elasticsearch-event-target)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
  - [Creating an Elasticsearch Target](#creating-an-elasticsearch-target)
  - [Expected CloudEvents](#expected-cloudevents)

## Prerequisites

A Elasticsearch cluster and a set of credentials:

- Version 7.x is preferred.
- User and password to the Elasticsearch cluster.
- An APIKey instead of User and password.
- CACertificate if using self-signed certificate and `SkipVerify` is not configured.
- Version 7.x is preferred.

## Creating an Elasticsearch Target

Anonymized Elasticsearch API object:

```yaml
apiVersion: targets.triggermesh.io/v1alpha1
kind: ElasticsearchTarget
metadata:
  name: <TARGET-NAME>
spec:
  connection:
    addresses:
      - <ELASTICSEARCH-URL>
    skipVerify: <true|false>
    caCert: <ELASTICSEARCH-CA-CERTIFICATE>
    apiKey:
      secretKeyRef:
        name: <SECRET-CONTAINING-APIKEY>
        key: <SECRET-KEY-CONTAINING-APIKEY>
    username: <ELASTICSEARCH-USERNAME>
    password:
      secretKeyRef:
        name: <SECRET-CONTAINING-PASSWORD>
        key: <SECRET-KEY-CONTAINING-PASSWORD>
  indexName: <ELASTICSEARCH-INDEX>
```

Connection must include at least one address, including protocol scheme and port.

- example: `https://elasticsearch-server:9200`

The connection must be filled with one of:

- `username` and `password`
- `apiKey`

If the Elasticsearch cluster is being served using a self-signed certificate the CA can be added, or TLS verify can be skipped:

- `caCert` for adding the PEM string for the certificate.
- `skipVerify` set to true for skip checking certificates.

Received events will be indexed using `indexName` as the elasticsearch index.

## Expected CloudEvents

Elasticsearch Target will forward any JSON payload at the CloudEvent to be indexed.

```console
curl -v http://elasticsearchtarget-es-indexinge5d0adf0209a48c23fa958aa1b8ecf0b.default.svc.cluster.local \
 -X POST \
 -H "Content-Type: application/json" \
 -H "Ce-Specversion: 1.0" \
 -H "Ce-Type: something.to.index.type" \
 -H "Ce-Source: some.origin/intance" \
 -H "Ce-Id: 536808d3-88be-4077-9d7a-a3f162705f79" \
 -d '{"message":"thanks for indexing this message","from": "Triggermesh targets", "some_number": 12}'
```
