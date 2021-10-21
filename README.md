# terraform-provider-luis
[![GitHub release](https://img.shields.io/github/release/BESTSELLER/terraform-provider-luis.svg)](https://github.com/BESTSELLER/terraform-provider-luis/releases/)
[![CircleCI](https://circleci.com/gh/BESTSELLER/terraform-provider-luis.svg?style=svg)](https://circleci.com/gh/BESTSELLER/terraform-provider-luis)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=BESTSELLER_terraform-provider-luis&metric=alert_status)](https://sonarcloud.io/dashboard?id=BESTSELLER_terraform-provider-luis)
![GitHub All Releases](https://img.shields.io/github/downloads/bestseller/terraform-provider-luis/total)
![GitHub Releases](https://img.shields.io/github/downloads/BESTSELLER/terraform-provider-luis/latest/total)

Terraform provider for managing luis.ai

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

### Testing

`go test`


## Generate Documentation

Docs are located in docs/.  These are autogenerated using a tool called `tfplugindocs`.

Install the tfplugindocs tool using:

```
$ go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
```

Generate the documentation at the root of this repository

```
$ tfplugindocs
```

