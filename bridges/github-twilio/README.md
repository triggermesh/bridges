# Github to Twilio

This sample bridge sends Github push events to a mobile phone as SMS using Twilio.

## Components

This bridge uses

- GithubSource as the producer of events
- A broker that receives events from github
- A broker that receives Twilio compliant events.
- A transformation ad-hoc service
- A subscription that will read Github broker, send events to transformation service and put the result in the Twilio broker
- A trigger that connects the Twilio broker with the target
- A Twilio target

## Parameterization

Customization of GitHub secrets and owner/repo, and Twilio secrets and default numbers are required.

- Refer to [GithubSource docs](https://knative.dev/docs/eventing/samples/github-source/) for the configuration and source event filtering.
- Refer to [Twilio docs](../../docs/targets/twilio.md) for customization of Twilio token and defaulting.
