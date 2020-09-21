# piha

![CodeQL](https://github.com/nullv01d/piha/workflows/CodeQL/badge.svg)

A twitter bot that automatically tweets a rant.

## What's with the name?
Named after the screaming bird [piha](https://en.wikipedia.org/wiki/Screaming_piha), it represents the bot's action - ranting/screaming. 😜

## Why is it built?
1. I wanted to practice my recently acquired Golang skills.
2. My internet service provider was extremely terrible at addressing requests. Tired of following up with them, I wrote this bot to do that for me repeatedly and automatically until my issues are resolved.
3. We all need a twitter bot at some point in our life! 😉

## How do I set it up?
1. You will need a twitter developer account. Follow the instructions [here](https://developer.twitter.com/en/portal) to set up a new twitter bot app and gather the following secrets:
    * API_KEY
    * API_KEY_SECRET
    * ACCESS_TOKEN
    * ACCESS_TOKEN_SECRET
2. run `cp .env.example .env`
3. edit `.env` file and set the secrets gathered from step #1
4. edit `.env` file and setup the below configurations
    * RANT_USER - the twitter handle of the service provider
    * RANT_DATE - the date in YYYY-MM-DD format when your issue with the service provider occurred
    * RANT_TEMPLATE - the rant tweet template. DO NOT change the text inside `{{}}` including the braces.

## How do I run it?
1. run `go mod download`
2. run `go run piha.go`

## TODO
1. Document repeated invocation via a Cron job or its alternative
2. Explore building binaries
3. Explore GitHub workflow integration
