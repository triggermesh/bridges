# Forward Kinesis messages to a Confluent Kafka cluster

This bridge listens to a kinesis stream and sends received messages to a Confluent topic.

## Components

This bridge uses

- AWS Kinesis source that ingest events into Triggermesh.
- A broker that receives events from Kinesis.
- A Confluent Target instance.
- A trigger that connects the broker with the target.

## Parametrization

Customization of Confluent's instance is required to provide credentials, servers, topics and connection options. Kinesis source is also customizable to provide ARN and credentials

- Refer to [AWS docs](../../docs/sources/aws.md) for configuring AWS Kinesis.
- Refer to [Confluent docs](../../docs/targets/confluent.md) for configuring Confluent.
