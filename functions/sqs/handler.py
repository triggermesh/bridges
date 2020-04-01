import boto3
import json
import os

# Create an SQS client
sqs = boto3.client('sqs')
sqs_queue_url = os.environ['QUEUE_URL']

def send(event, context):
    response = sqs.send_message(QueueUrl=sqs_queue_url,
                                MessageBody=json.dumps(event))
    return response
