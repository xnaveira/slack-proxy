# Slack proxy

###Proxy messages to a Slack Incoming WEBHOOK_URL

Make sure that you create a webhook in slack and get the url.

Run the container

```
export WEBHOOK_URL=<your_webhook_url>
docker run -ti -e WEBHOOK_URL=$WEBHOOK_URL -p 8080:8080 xnaveira/slack-proxy
```

Send a message to your channel through slack-proxy

```
curl -v -X POST -H 'Content-type: application/json' \
--data '{"text":"This is a message"}' http://localhost:8080/message
