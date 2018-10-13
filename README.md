# Flood-Alert-App
More Detailed info + instructionms coming soon

## App-Flood-App
GRPC Server - Waits for calls from `App-Flood-Ctl` 
List of calls that the client can make:
- check-alerts - Requests info from api and checks for severe weather (currently just certain weather codes + rain in the past 3 hours above 3 hours)
- test-alerts - Tests alerts methods - Currently just Text Message and Email 
- mute-notifications - mutes notifications of each alert type 
	- 1 = Call (Not implemented yet)
	- 2 = Text Message
	- 3 = Email

## App-Flood-Ctl
Small command-line app, that makes requests to the Flood-Server

To make requests locally i'd recommend installing it locally and following the instructions below (once they are finished)
The dockerfile is used for the cron job aspect of the project. Its hard coded to make requests to localhost, for use inside the cluster.

### Installation:
TODO: Start with how to setup the cluster / setup with gcp
 - On your Local Machine:
	`You must have go installed and a running version of the app running`
	1.  By runing `go install` the binary will be built and put into your path
	2. run `app-flood-ctl` to get a list of commands
	3. Before you can make requests locally you must the set the enviroment variable `KUBERNETES_IP` to the external ip created by the service



## App-Flood-Build

Handles creating and populating the kubernetes cluster

CronJob: In order to check requests, kubernetes manages a cronjob that runs the following command `app-flood-ctl check-alerts` every 30min. This is how the app is able to check alerts.
