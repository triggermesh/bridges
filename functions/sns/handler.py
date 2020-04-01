import boto3
import json

# Create an SNS client
sns = boto3.client('sns')

def send(event, context):
    # Publish a simple message to the specified SNS topic
    response = sns.publish(
        TopicArn='arn:aws:sns:us-east-1:587264368683:triggermesh',    
        Message=json.dumps(event),    
    )

    # Print out the response
    return response
