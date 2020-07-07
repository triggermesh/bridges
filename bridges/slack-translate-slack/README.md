# Slack to Slack with a translation

This bridge listens in a slack channel where a bot is invited and speaks the translated conversation in another channel.

You need:

* two Slack bots, or the same Slack bot configured for Event API and Web API.
* a Kubernetes secret holding the Slack Target bot token, optionally a second secret that contains the signing secret to verify signature for Slack Source received messages.

The source for the translation is at [https://github.com/sebgoa/transform/tree/master/translate](https://github.com/sebgoa/transform/tree/master/translate)
