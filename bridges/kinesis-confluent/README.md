# Forward Kinesis messages to a Confluent Kafka cluster

Create a secret containing your AWS API credentials


Create a secret containing your Confluent credentials

```
kubectl apply -f secret.yaml
```

Configure the Bridge to your liking (e.g change the ARN of the Kinesis stream) and create it

```
kubectl apply -f kinesis-confluent.yaml
```

Now put a message in your Kinesis stream and see it appear in your Confluent cluster, for example

```
aws kinesis put-record --stream-name everybridge --partition-key 123 --data 'hi everyone'
```
