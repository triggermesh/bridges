# Send SQS messages via Email with SendGrid

Create a secret containing your SendGrid API Key

## Components

This bridge uses

- SQS as a source for messages
- A broker that receives events from SQS
- A broker that receives SendGrid compliant events.
- A transformation ad-hoc service
- A subscription that will read SQS broker, send events to transformation service, and put the result in the SendGrid broker
- A trigger that connects the SendGrid broker with the target
- A Sendgrid target

## Parameterization

Customization of SQS and secret creation for SendGrid token is required.

- Refer to [Source docs](../../docs/sources/aws.md) for the message being produced.
- Refer to [Sendgrid docs](../../docs/targets/sendgrid.md) for customization of API key and default values.

You can also use kubectl and provided assets at this directory

```sh
# Edit with your SendGrid secret
kubectl apply -f sendgrid-secret.yaml

# Edit the Bridge manifest to configure your SQS ARN and your AWS API credentials secret
kubectl apply -f sqs-sendgrid.yaml
```
