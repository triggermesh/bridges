# Twilio event target

This event target integrates with Twilio using received Cloud Event messages to compose and send SMS.

## Contents

- [Twilio event target](#twilio-event-target)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
  - [Creating an Twilio Target](#creating-an-twilio-target)
  - [Expected CloudEvents](#expected-cloudevents)

## Prerequisites

A Twilio [configured account](https://support.twilio.com/hc/en-us/articles/223136027-Auth-Tokens-and-How-to-Change-Them) is required to run this target:

- Register a Twilio account
- Purchase a phone number with
- Retrieve from Twilio Dashbard Account SID
- Retrieve from Twilio Dashbard Auth Token

## Creating an Twilio Target

Anonymized Twilio API object:

```yaml
apiVersion: targets.triggermesh.io/v1alpha1
kind: TwilioTarget
metadata:
  name: <TARGET-NAME>
spec:
  defaultPhoneFrom: "<PHONE-FROM>"
  defaultPhoneTo: "<PHONE-TO>"
  sid:
    secretKeyRef:
      name: "<YOUR-SID-SECRET>"
      key: "<YOUR-SID-SECRET-KEY>"
  token:
    secretKeyRef:
      name: "<YOUR-TOKEN-SECRET>"
      key: "<YOUR-TOKEN-SECRET-KEY>"
```

- `sid` and `token` are needed to connect to Twilio.
- `defaultPhoneFrom` will usually be configured matching the phone number you have purchased at Twilio. It is not mandatory and can be overriden at each received message.
- `defaultPhoneTo` is not mandatory and will be used if the received Cloud Event message does not inform one.

Refer to [Twilio docs for number formating](https://www.twilio.com/docs/lookup/tutorials/validation-and-formatting?code-sample=code-lookup-with-international-formatted-number).

## Expected CloudEvents

Twilio Target expect a JSON payload at the CloudEvent that includes:

- `message`: text to be sent.
- `media_urls`: array of URLs pointing to JPG, GIF or PNG resources.
- `from`: phone sourcing the communication. Optional if provided by the TWilioTarget.
- `to`: phone destination. Optional if provided by the TwilioTarget.

Example:

```JSON
{
    "from": "+1111111111",
    "to": "+2222222222",
    "message":"Hello from Triggermesh using Twilio!",
    "media_urls": ["https://66.media.tumblr.com/tumblr_lrbu1l9BJk1qgzxcao1_250.gifv"],
}
```
