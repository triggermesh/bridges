# Slack to Slack

This sample bridge reads messages from a specfied Slack channel and posts them to another specified channel.

## Components

This bridge uses

- Slack source as the producer of events
- A Broker that receives message events from the Slack Source
- A Trigger that connects the Broker with the Slack Target 
- A Slack Target


## Setup

- When creating and configuring the Slack Source & Target 'Bot's'. You will need to grant proper OAuth roles, as well as include them into thier respective intended channel(s). 



- Refer to [Slack Source docs](../../docs/sources/slack.md) for configuring a Slack Source
- Refer to [Slack Target docs](../../docs/targets/slack.md) for configuring a Slack Target.
