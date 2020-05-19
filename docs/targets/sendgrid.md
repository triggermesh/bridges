# SendGrid event target

This event target integrates with SendGrid using received Cloud Event messages to send emails.

## Contents

- [SendGrid event target](#sendgrid-event-target)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
  - [Creating SendGrid Target](#creating-sendgrid-target)
  - [Expected CloudEvents](#expected-cloudevents)

## Prerequisites

A SendGrid account API token

## Creating SendGrid Target

Anonymized SendGrid API object:

```yaml
apiVersion: targets.triggermesh.io/v1alpha1
kind: SendgridTarget
metadata:
  name: <TARGET-NAME>
spec:
  defaultFromName: <SENDER-NAME>
  defaultFromEmail: <SENDER-EMAIL>
  defaultToName: <ADDRESSER-NAME>
  defaultToEmail: <ADDRESSER-EMAIL>
  apiKey:
    secretKeyRef:
      name: <SECRET-CONTAINING-APIKEY>
      key: <SECRET-KEY-CONTAINING-APIKEY>
```

API key is a mandatory field. All default fields will be used if any of them is not present at the event message received by the target.

## Expected CloudEvents

SendGrid Target will create an email with CloudEvent JSON payloads that conform to this format:

```JSON
{
    "message":"sent via Triggermesh",
    "fromemail": "lana@triggermesh.io",
    "fromname": "Lana W.",
    "toemail": "lilly@triggermesh.io",
    "toname": "Lilly W.",
    "media_urls": ["https://66.media.tumblr.com/tumblr_lrbu1l9BJk1qgzxcao1_250.gifv"],
  }
```

All from and to fields will be default to Target configured fields if not present.
