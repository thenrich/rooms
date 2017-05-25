# Rooms -- Ephemeral conference rooms powered by Twilio and App Engine

### Setup your Twilio account

 **!!! You will be charged by-the-minute for Twilio's phone service !!!**

Head over to [Twilio](https://twilio.com) and setup an account. Purchase a phone number and grab your auth token from the [console](https://www.twilio.com/console).

### Setup a Google App Engine project

It's easiest to use the `gcloud` tool to do this which is available [here](https://cloud.google.com/sdk/gcloud/).

Run `gcloud projects create PROJECT_NAME` to create a new project.

### Setup Rooms

First, make sure you have [Go](https://golang.org/) installed.

Then, `go get github.com/thenrich/rooms`. Modify `src/github.com/thenrich/rooms/app.yaml`, replacing `YOUR_TWILIO_API_KEY` with your Twilio auth token mentioned above and `YOUR_APP_ENGINE_URL` with `https://[PROJECT_NAME].appspot.com`.

Now, deploy the app to App Engine:

`gcloud --project [PROJECT_NAME] beta app deploy`

If `gcloud` complains about `GOPATH` dependencies, make sure `rooms` is in your `GOPATH`.

### Finish Twilio setup

Go back to your Twilio console and add an incoming webhook to the phone number you purchased. The incoming webhook should be a `GET` request to `https://[PROJECT_NAME].appspot.com/calls/incoming`

### Test

Call your Twilio number and enter any conference ID when prompted. You should be forwarded to the conference room with wait music playing.
