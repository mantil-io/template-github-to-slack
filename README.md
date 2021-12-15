# template-github-to-slack
Send slack notification when someone stars your Github repository

# What problem does it solve?

This is example of Mantil app which sends notification to the Slack channel each time someone stars your Github repository.

# How do you install and run it?

To use this template you first need to create new Mantil project.

```
mantil new app --from github-to-slack
cd app
```

Before creating first stage and deploying your application you will need to create Slack webhook and add it as environment variable for your function.
For instructions on how to do that please check [What's necessary configuration?](#whats-necessary-configuration).

After configuring environment variable you can proceed with creation of the first stage.

```
mantil stage new development
```

This will create new stage called `development` and deploy it to your node.

# What's necessary configuration?

Before creating your first stage you will need to create Slack webhook which will be used to post notifications to your Slack channel.

Detailed instructions on how to create a webhook can be found [here](https://slack.com/help/articles/115005265063-Incoming-webhooks-for-Slack).

Once your webhook is created you can add URL to the `config/environment.yml` file as env variable for your function.

```
project:
  env:
    SLACK_WEBHOOK: # add your slack webhook here
```

Now you can create and deploy your first stage and output stage endpoint with `mantil env -u`

API endpoint for your function will have name of that function in the path. In our case that is `$(mantil env -u)/handler`.

With this URL we can now create Github webhook which will invoke our lambda function on each star to our Github repository.

To create new Github webhook, go to your repository and choose `Settings - Webhooks - Add webhook`.

Under `Payload URL` add your function endpoint, choose `application/json` as content type and for the trigger events choose `Let me select individual events` and toggle `Watches`.

If you did everything correctly Github will send you test trigger which will invoke your function and send notification to your Slack channel.
If something goes wrong you can always use `mantil logs` to check the logs of your function and try to pinpoint the problem.

# How does it work?

# How do you modify it?

If you want different behaviour out of your function you can edit trigger events of your Github webhook and make necessary changes to your code in `api/handler.go`
Examples of payloads for all Github events can be found in their [docs](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#watch).

After each change you have to deploy your changes with `mantil deploy`.

