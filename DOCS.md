Use the wechat plugin to notify users during the Drone pipeline.

Please only use these paramters for now:

* `access_token` - The access token for authorization
* `agentid` - The agent id to send the message
* `msgurl` - The agent id to send the message
* `btntxt` - The text for the button on the card
* `title` - Title of the card
* `description` - Text description of the card

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
wechat:
  access-token: somelongasstoken
  agentid: 12345
  msgurl: http://acoolwebsite.com
  btntxt: click
  title: Title for the card
  description: This is the card body
```
