Use the wechat plugin to notify users during the Drone pipeline.

Please only use these paramters for now:

* `corpid` - The corpid for authorization
* `corp_secret` - The corp secret for authorization
* `agent_id` - The agent id to send the message
* `msg_url` - The agent id to send the message
* `btntxt` - The text for the button on the card
* `title` - Title of the card
* `description` - Text description of the card

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  wechat:
    image: clem109/drone-wechat
    corpid: corpid
    corp_secret: secret
    agent_id: 1234567
    title: ${DRONE_REPO_NAME}
    description: "Build Number: ${DRONE_BUILD_NUMBER} failed. ${DRONE_COMMIT_AUTHOR} please fix. Check the results here: ${DRONE_BUILD_LINK} "
    msg_url: ${DRONE_BUILD_LINK}
    btn_txt: btn
    when:
      status: [ failure ]
```

If you want to add secrets that you set in the drone UI make sure you use the
correct naming scheme, all parts of the build can be hidden and set within the
Drone UI, please consult the [main.go](main.go) EnvVar:

```yaml
pipeline:
  wechat:
    image: clem109/drone-wechat
    secrets: [plugin_corpid, plugin_corp_secret, plugin_agent_id]
    title: ${DRONE_REPO_NAME}
    description: "Build Number: ${DRONE_BUILD_NUMBER} failed. ${DRONE_COMMIT_AUTHOR} please fix. Check the results here: ${DRONE_BUILD_LINK} "
    msg_url: ${DRONE_BUILD_LINK}
    btn_txt: btn
    when:
      status: [ failure ]
```
