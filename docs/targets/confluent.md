# Confluent event target

This event target integrates with Confluent, forwarding Cloud Event messages to topics.

## Contents

- [Confluent event target](#confluent-event-target)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
  - [Creating an Elasticsearch Target](#creating-an-elasticsearch-target)

## Prerequisites

A Confluent clusterand a set of credentials:

- User and password to the Elasticsearch cluster.
- An APIKey instead of User and password.
- CACertificate if using self-signed certificate and `SkipVerify` is not configured.
- Version 7.x is preferred.

## Creating an Elasticsearch Target

Anonymized Elasticsearch API object:

```yaml
apiVersion: targets.triggermesh.io/v1alpha1
kind: ConfluentTarget
metadata:
  name: <TARGET-NAME>
spec:
  topic: <TOPIC-MESSAGES-ARE-BEING-SENT-TO>
  bootstrapservers:
    secretKeyRef:
      name: <SECRET-CONTAINING-BOOTSTRAP-SERVERS>
      key: <SECRET-KEY-CONTAINING-BOOTSTRAP-SERVERS>
  username:
    secretKeyRef:
      name: <SECRET-CONTAINING-USERNAME>
      key: <SECRET-KEY-CONTAINING-USERNAME>
  password:
    secretKeyRef:
      name: <SECRET-CONTAINING-PASSWORD>
      key: <SECRET-KEY-CONTAINING-PASSWORD>
```

If topic is not informed a default one byt the name `tmkafka` will be used. Bootstrap servers, user name and password are mandatory.

When not existing the topic will be created.