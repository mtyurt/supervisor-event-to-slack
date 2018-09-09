# supervisor-event-to-slack

This showcase repository listens to supervisor events and sends these events to Slack properly. Generic [supervisor-event-handler](https://github.com/mtyurt/supervisor-event-handler) is used to process supervisor events.

# installation

- Download from Releases 
- Install from source:


```bash
git clone git@github.com:mtyurt/supervisor-event-to-slack
cd supervisor-event-to-slack
vgo build
```

# usage
This program needs to be provided to supervisor configuration as event listener. A bare minimum configuration would be:

```bash
$ cat > eventlistener.ini <<EOF
[eventlistener:status_listener]
command=/path/to/supervisor-event-to-slack
events=PROCESS_STATE
autostart=true
environment=SLACK_TOKEN="your-slack-token",SLACK_CHANNNEL="channel-to-post-messages"
EOF
```

Copy this file to supervisor.d config directory, by default `/etc/supervisor.d/`, and restart supervisord. 

Check out http://supervisord.org/events.html for event types and their payloads.

# troubleshooting
- Make sure Slack token and channel are correct.
- Check if `status_listener` is working via `supervisorctl status all`, restart supervisord if it does not.
- Check `/var/log/supervisord.log` and `/var/log/supervisor/status_listener-*.log` for further information.

# licence
The BSD 3-Clause License - see [LICENSE](https://github.com/mtyurt/supervisor-event-to-slack/blob/master/LICENSE) for more details.

