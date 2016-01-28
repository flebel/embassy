Embassy acts as a gateway to APIs.

This software is a work in progress and is not suited for use in production environments. Backwards incompatibilities and regressions are to be expected.

# Motivation

As IoT devices become mainstream, deployments are likely to experience an unhealthy duplication of credentials and logic across their IoT fleet to communicate with third party APIs. Embassy can be used instead of, or as a gateway to a more complex IoT management platform.

# Getting started

1. Write custom ambassador if needed
1. Configure `config.json` file
1. `go run cmd/main.go`

# Usage

A minimum of 3 settings must be set for every ambassador:

        Ambassador    The name of the ambassador, as found in the ambassador's `Name` constant
        Path          Path at which the ambassador endpoint should be exposed
        HTTPVerb      HTTP verb at which the ambassador endpoint should be exposed, only `GET`, `HEAD` and `POST` are supported at this time

A fourth setting named `Configuration` can be required by an ambassador. Consult the ambassador's documentation and code for the available settings.

## Built-in ambassadors

Embassy comes with a generic ambassador that allows for creating simple API gateways from a JSON configuration file, and an ambassador for the `jsonip.com` service. The `config.json` file contains an example of how to use either ambassador, with the generic ambassador configured to query `jsonip.com`, essentially resulting in an identical behavior and HTTP response for those two ambassadors.

## Using the generic ambassador

The `generic` ambassador can be used for simple one off communications over HTTP and accepts an URL and HTTP verb over which a request should be made. Support for passing custom headers and form data is on the roadmap.

Configuration::

        URL         string # e.g. "http://jsonip.com"
        HTTPVerb    string # "GET" or "POST"

## Writing a custom ambassador

Embassy comes pre-configured with `generic` and `jsonip` ambassadors. Both ambassadors produce the same result, however, the `jsonip` ambassador is hardcoded to query the `http://jsonip.com` API and retrieve `embassy`'s external IP address, and the `generic` ambassador is configured to perform a `GET` HTTP request to `http://jsonip.com/` through the `config.json` file.

The `jsonip` ambassador can be used as a template for writing a simple custom ambassador to support functionality lacking in the `generic` ambassador.

