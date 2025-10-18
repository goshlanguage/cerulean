# Postmortem

## Overview

We originally started Cerulean as an "artisanally crafted" project where we did everything by hand. This worked fine for the "minimum-minimum viable product" but we quickly realized it was going to be unwieldy to manage this in the long term as we began to support more endpoints.

## Context

1. Azure has an HTTP API for interacting with your cloud resources
1. [Microsoft's AutoRest](https://github.com/Azure/autorest) takes the API definitions (written in a Microsoft-specific subset of Swagger) and generates code for the clients of the Azure API
1. `azure-sdk-for-go` (the Go client for the Azure API) is one of the generated clients (which means it uses HTTP/AutoRest under the hood)

## Blockers

- Sheer size/scale of manually creating/maintaining all the API mocks for each endpoint
- There's not a clear standard set for the API pathing in the Azure API
- API endpoints are versioned via arbitrary dates and endpoints reference other endpoints in no discernible pattern which makes it difficult to determine which endpoints reference/use other endpoints
    - (There's no "latest" version which created a scaffolding nightmare)

## Designs Attempted

1. Manually mocking out unlinked, fake backend objects that were interacted with via the official API paths
1. Took fake backend objects and tried to find links between them to create a full, mock database of "real" data

## Designs Considered

1. Automatically capturing API calls and responses and using those to mock out valid JSON payloads and such
    - Wouldn't automatically determine links between objects (e.g. a security group JSON wouldn't be automatically returned when you hit the "all resources" API endpoint)

## Final Thoughts

It is still likely possible to use the AutoRest code generator to create mocks and scaffolding, but there was enough overhead involved that it felt unwieldy and not worth the effort at the time.

We could have potentially gained more up-front traction had we gone with our initial plans of pulling in a customer or two early on to keep the quality of the project (e.g. documentation, etc.) as well as getting feedback on where we could offer the most benefit with the least effort.
