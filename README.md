This template is an example of serverless integration between GitHub and Slack built on AWS Lambda. It enables you to process and forward events from GitHub to Slack instantly via webhooks meaning you can modify output based on your own needs, either per some criteria or with information in output. To create your own integration just follow instructions below. Expected output will look something like this:

![image](https://github.com/mantil-io/template-github-to-slack/blob/master/images/gh2s_image.png) 

# Prerequisites

This template is created with Mantil. To download [Mantil CLI](https://github.com/mantil-io/mantil#installation) on Mac or Linux use Homebrew

```
brew tap mantil-io/mantil
brew install mantil
```

or check [direct download links](https://github.com/mantil-io/mantil#installation).

To deploy this application you will need:
- An [AWS account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/)
- A GitHub account with a repository where you have admin rights
- A Slack account with the right to create apps

# Installation

To locally create a new project from this template run:

```
mantil new app --from github-to-slack
cd app
```

# Configuration 

Before deploying your application you will need to create a Slack webhook and add it as an environment variable for your function which will be used to post notifications to your Slack channel.

Detailed instructions on how to create a webhook can be found [here](https://slack.com/help/articles/115005265063-Incoming-webhooks-for-Slack).

Once your webhook is created you need to add URL to the `config/environment.yml` file as env variable for your function.

```
project:
  env:
    SLACK_WEBHOOK: # add your slack webhook here
```

# Deploying an application

Note: If this is the first time you are using Mantil you will first need to install Mantil Node on your AWS account. For detailed instructions please follow these simple, one-step [setup instructions](https://github.com/mantil-io/mantil/blob/master/docs/getting_started.md#setup)

```
mantil aws install
```

After configuring the environment variable you can proceed with application deployment.

```
mantil deploy
```

This command will create a new stage for your project with default name `development` and deploy it to your node.

Now you can output the stage endpoint with `mantil env -u`. The API endpoint for your function will have the name of that function in the path, in our case that is `$(mantil env -u)/star`.
With this URL we can now create a Github webhook which will invoke our Lambda function on each star to our Github repository.

# Setting up Github webhook

To create new Github webhook, go to your repository and choose `Settings - Webhooks - Add webhook`.

Under `Payload URL` add your function endpoint, choose `application/json` as content type and for the trigger events choose `Let me select individual events` and toggle `Watches`.

If you did everything correctly Github will send you a test trigger which will invoke your function and send notification to your Slack channel.

Congratulations, you just created and deployed a fully functional serverless AWS Lambda application!

# Modification

If you want different behavior out of your project you can add more triggers by creating new webhooks and new functions. Examples of payloads for all Github events can be found in their [docs](https://docs.github.com/en/developers/webhooks-and-events/webhooks/webhook-events-and-payloads).

Adding new function will be demonstrated with `fork` function which sends slack notification every time your repository gets forked. Implementation of this function is already available in the repository.

New function was created with

```
mantil generate api fork
```

This generated necessary files for our new lambda function which was then further edited to suit our needs. In the case of `fork` function `api/fork.go` contains necessary logic for this trigger.

Together with the function another Github trigger was created containing `$(mantil env -u)/fork` payload URI and `Fork` as an event.

After each change you have to deploy your changes with `mantil deploy`, or instruct Mantil to automatically deploy all saved changes with `mantil watch`.

For more detailed instructions please refer to [Mantil documentation](https://github.com/mantil-io/mantil#documentation).

# Cleanup

To remove the created stage with all resources from your AWS account destroy it with

```
mantil stage destroy development
```

# Final thoughts

With this template you learned how to create a simple serverless application with AWS Lambda via Mantil where Lambda is invoked by webook. Check out our [documentation](https://github.com/mantil-io/mantil#documentation) to find more interesting templates. 

If you have any questions or comments on this concrete template or would just like to share your view on Mantil contact us at [support@mantil.com](mailto:support@mantil.com) or create an issue.
