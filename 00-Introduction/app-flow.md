# Introduction

Build a simple service to understand general dev workflow.

## Local

Local service to do some Docker

## Simple Service

Simple request/resonse scenario. Demonstrates:

- building clients and services on Dapr
- language portability
- HTTP and gRPC communication protocols

### Sender Client (Python)

Simple Python HTTP client that sends messages to the Receiver Service

### Receiver Service (Node)

Simple Node HTTP server that prints messgages for each request it receives.

## Deployment

Deployed to K8s cluster that is running Dapr. Uses deployment yaml.
