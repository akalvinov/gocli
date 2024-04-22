# gocli

## What is gocli 

gocli is a console tool that simplifies acquiring data from AWS Secrets manager. gocli can validate existance of env variables it is optional feature. 

## Configuration

There are two ways of configuring gocli - Env variables or gocli.yaml file

### Parameters to configure

secretName (gocli_SECRETNAME) - name of the secret in AWS secret manager to acquire data from

envFileName (gocli_ENVFILENAME) - filename to save acquired secrets

validate - list of Env variables to validate

secretsMap - map of AWS secrets and filenames of files to save

slackUrl (gocli_SLACKURL) - Slack url for sending messages

### Commands

secrets - Acquires data from AWS Secret manager and saves it to env file and secret files
validate - Validate list of specified variables to be set in environment
