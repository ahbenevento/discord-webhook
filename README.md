# Discord Webhooks

Very simple example to send Discord messages through webhooks in golang.

Is possibly define a list of webhooks and use one of this by a prefix channel ID or alias defined. See `discordwh.conf-example` for details.

Packages used:

- [gtuk/discordwebhook](https://github.com/gtuk/discordwebhook)

## Use

```console
discordwh [-u username] [-a avatar-url] channel "message"
```

- **username** Username used for message.
- **avatar-url** Avatar image URL.
- **channel** Full or partial channel ID, alias (configured) or channel URL.
- **message** Text of message.
