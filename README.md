# Drone Pushover Plugin

A simple drone plugin to send push notifications via Pushover.

Usage:

```
kind: pipeline
name: default

steps:
- name: pushover  
  image: awerv/drone_pushover
  settings:
    user:
      from_secret: pushover_user
    token:
      from_secret: pushover_token
    title: test
    message: lorem ipsum...
```

It's recommended to supply the user key and the token via secrets.

Check [this](https://pushover.net/api/subscriptions) for more information regarding creating a Pushover subscription.