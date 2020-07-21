# Zendesk ticket to Zendesk tag

![Thumb up](./assets/thumb-up.gif)

This bridge is designed to complete the following event *flow*:

1. Receive _New Ticket_ events from the Zendesk event source.
1. Determine the Ticket's sentiment via [AWS Comprehend][aws-comprehend] using a custom event transformation.
1. Label the new Zendesk Ticket using the Zendesk event target, based on the outcome of the AWS Comprehend sentiment
   analysis.

## Architecture overview

An instance of the Zendesk event source sends events to a central event broker. A transformation service consumes those
events based on their type, interacts with AWS Comprehend, and pushes new events back to the broker. An instance of the
Zendesk event target consumes events resulting from the transformation.

```
+--------+       +--------+       +--------+
| source |-------> broker |-------> target |
+--------+  (1)  +---^----+  (3)  +--------+
                     |
		 (2) |
                     |
                 +---v-----+       +-------+
                 | transf. <- - - ->  AWS  |
                 +---------+       +-------+
```

## Prerequisites

### Transformation service configuration

The AWS Comprehend transformation component requires the following parameters:

* [**AWS Access Key ID and Secret Access Key**][aws-creds]: credentials to authenticate API requests to AWS Comprehend.
* [**Region**][aws-region]: AWS region in which the transformation service interacts with the AWS Comprehend API.
* [**Language**][comprehend-lang]: the language AWS Comprehend should expect in the body of the Zendesk Ticket.

### Secrets

The event bridge reads several secrets, such as the credentials mentioned above, from the following Kubernetes Secret
objects:

* `zendesk-api`: credentials for the Zendesk source and target
* `awscomprehend`: credentials for the Transformation service (AWS Comprehend sentiment analysis)

Those Secrets need to exist in the TriggerMesh user's namespace for the bridge components to start. Please refer to the
[secrets.yaml](./secrets.yaml) file for an example of Secret objects and their expected attributes.

## More Info

Please refer to the [TriggerMesh documentation][tm-doc] for more information about configuring the aforementioned event
sources and targets.

[aws-comprehend]: https://aws.amazon.com/comprehend/
[aws-creds]: https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys
[aws-region]: https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/
[comprehend-lang]: https://docs.aws.amazon.com/comprehend/latest/dg/supported-languages.html
[tm-doc]: https://docs.triggermesh.io/
