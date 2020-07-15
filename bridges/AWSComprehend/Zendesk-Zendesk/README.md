# Zendesk to Zendesk with sentiment analysis via AWS Comprehend

## Bridge Intention

This bridge is designed to complete the following event *flow*:

* Recieve new Ticket Event's from Zendesk

* Determine the Ticket's sentiment via [AWS Comprehend](https://aws.amazon.com/comprehend/) 

* Create a label for the new Ticket's baised on the outcome of [AWS Comprehend](https://aws.amazon.com/comprehend/)

### Prerequisites

From [AWS](https://aws.amazon.com/):

* [AWS Access key ID](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html)

* [AWS Secret Access key](https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html)

* [Region](https://aws.amazon.com/premiumsupport/knowledge-center/vpc-find-availability-zone-options/#:~:text=To%20find%20which%20Availability%20Zones,options%20in%20the%20Region%20selector.)

* Language - This will tell Comprehend what language to expect (`en` denotes english)

From [Zendesk](https://www.zendesk.com/):

* [API Token](https://support.zendesk.com/hc/en-us/articles/226022787-Generating-a-new-API-token)

* [subdomain](https://support.zendesk.com/hc/en-us/articles/221682747-Where-can-I-find-my-Zendesk-subdomain-)

### Deployment

After populating the fields denoted `#EXAMPLE` in the `bridge.yaml` file located in the `Zendesk-Zendesk` folder you can deploy by executing the following command from within this directory.

    kubectl -n <namespace> apply -f bridge.yaml

### Project pages for brige compoments

Both of the [Zendesk Source](https://github.com/triggermesh/knative-sources) and [Zenformation](https://github.com/JeffNaef-Triggermesh/transformations/tree/master/zenformation) are open source.

The Zendesk Target is not open source at the moment.
