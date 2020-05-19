# Forward GitLab push events to AWS

These are a collection of sample bridges that will send a post notification to
  * AWS Lambda
  * AWS Simple Notification Service (SNS)
  * AWS Simple Queueing Service (SQS)
  * AWS Kinesis

All require secrets for both GitLab and AWS. Update the `100-secrets.yaml` file
with the correct credentials and apply:

    kubectl apply -f 100-secrets.yaml


Configure the Bridge to your liking.  Changes will be required such as to the ARN
to reflect the location of the service being called as well as the GitLab
respository to instrument and the type of events to look for.

_NOTE: Kinesis requires a partition to be specified in addition to the ARN_

Lastly, create the bridge:

    kubectl apply -f gitlab-awslambda.yaml
