# Slack event source

Slack Source enables integration between slack messages read by a bot users and the event platform.

## Contents

- [Slack event source](#slack-event-source)
  - [Contents](#contents)
  - [Prerequisites](#prerequisites)
    - [Creating an Slack Source instance](#creating-an-slack-source-instance)
  - [Produced CloudEvents](#produced-cloudevents)

## Prerequisites

A Slack Bot User that uses the RealTime API

1. Customize Slack adding a new **clasic** bot at https://api.slack.com/apps?new_classic_app=1
2. From Basic Information section, display Add features and functionality and select bots, then click on `Add Legacy Bot User`
3. Select `Install App` section and click on `Install App to Workspace`
4. Copy the bot user OAuth token (should begin with `xoxb-`)

### Creating an Slack Source instance

An instance of the Slack Source is created by creating a manifest at your cluster where it is informed of:

- The namespace where the instance of the source adapter will run.
- The kubernetes secret and key that host the bot token copied when configuring the Slack bot.
- An optional threadiness parameter in case we need more than one thread for sink dispatching.
- The sink addressable where events will be sent.

```yaml
apiVersion: sources.triggermesh.io/v1alpha1
kind: SlackSource
metadata:
  name: triggermesh-knbot
spec:
  slackToken:
    secretKeyRef:
      name: slack
      key: token
  threadiness: 1
  sink:
    ref:
      apiVersion: eventing.knative.dev/v1beta1
      kind: Broker
      name: slack
```


## Produced CloudEvents

The Slack Source creates a cloud event for each message written at a channel where the bot is added and also to direct messages to the bot.

- type: `com.slack/message`
- source: `com.slack.<WORKSPACE>`
- subject: `<CHANNEL-WHERE-THE-MESSAGE-WAS-HEARD>`
- data: JSON structure that contains:

```json
   {
     "user_id": "<USER-WRITING-THE-MESSAGE>",
     "text": "<MESSAGE-CONTENTS>"
   }
```
