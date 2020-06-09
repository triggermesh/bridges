# Slack to Slack with a translation

This bridge listens in a slack channel where a bot is invited and speaks the translated conversation in another channel.

You need:

* two Slack bots
* two Kubernetes secrets holding the Slack bot token
* a bridge

The source for the translation is at [https://github.com/sebgoa/transform/tree/master/translate](https://github.com/sebgoa/transform/tree/master/translate)
