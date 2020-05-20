# Confluent event target

This event target integrates with Confluent forwarding Cloud Event messages to topics.

## Contents

- [Confluent event target](#confluent-event-target)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
  - [Creating an Confluent Target](#creating-an-confluent-target)

## Prerequisites

A Confluent cluster and a [set of credentials](https://docs.confluent.io/current/cloud/using/api-keys.html):

- User and password to the Confluent cluster.
- Bootstraps servers address.
- Topic (can be dynamically created)

## Creating an Confluent Target

Anonymized Confluent API object:

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
