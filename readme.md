# CogFlow

CogFlow is a command-line tool designed to simplify interactions with Amazon Cognito user pools and groups. It provides an intuitive way to manage users and groups within your Amazon Cognito setup using simple commands.

## Installation

To install CogFlow, use the following command:

```shell
go install github.com/abakermi/cogflow/cmd/cogflow
```

## Usage

**CogFlow** provides several commands to manage users and groups in Amazon Cognito. Here are some examples:
### List all available groups:
```shell
cogflow group list
```
### Create a new group:
```
cogflow group create -g groupname
```
### Delete a group:
```
cogflow group delete -g groupname
```
### Enable a user:
```shell
cogflow user enable -u userid
```
### Disable a user:
```shell
cogflow user disable -u userid
```
### Delete a user:
```shell
cogflow user delete -u userid
```