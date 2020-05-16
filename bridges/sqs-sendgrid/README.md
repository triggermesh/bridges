# Send SQS messages via Email with SendGrid

Create a secret containing your SendGrid API Key

```
kubectl apply -f secret.yaml
```

Edit the Bridge manifest to configure your SQS ARN and your AWS API credentials secret

Then create the bridge

```
kubectl apply -f sqs-sendgrid.yaml
```
