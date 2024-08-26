[![Actions Status](https://github.com/unobravo/pulumi-mysql/actions/workflows/release.yml/badge.svg)](https://github.com/unobravo/pulumi-mysql/actions)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fmysql.svg)](https://www.npmjs.com/package/@unobravo/pulumi-mysql)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-mysql/sdk/v3/go)](https://pkg.go.dev/github.com/unobravo/pulumi-mysql/sdk/go)
[![Static Badge](https://img.shields.io/badge/Terraform_Provider-v3.0.63-purple)](https://github.com/petoju/terraform-provider-mysql)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# MySQL Resource Provider

The MySQL resource provider for Pulumi lets you manage MySQL resources in your cloud programs.  To use
this package, please [install the Pulumi CLI first](https://pulumi.io/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @unobravo/pulumi-mysql

or `yarn`:

    $ yarn add @unobravo/pulumi-mysql

### Python

To use from Python, install using `pip`:

    $ pip install git+https://github.com/unobravo/pulumi-mysql.git@v0.0.0#subdirectory=sdk/python

> [!NOTE]
> NOT PUBLISHED ON PyPI

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/unobravo/pulumi-mysql/sdk

### .NET

> [!WARNING]
> NOT PUBLISHED

## Configuration

The following configuration points are available:

- `mysql:endpoint` (required) - The address of the MySQL server to use. Most often a "hostname:port" pair, but may also
  be an absolute path to a Unix socket when the host OS is Unix-compatible. Can be set via `MYSQL_ENDPOINT` environment variable.
- `mysql:username` (required) - Username to use to authenticate with the server. Can be set via `MYSQL_USERNAME` environment variable.
- `mysql:password` - (optional) Password for the given user, if that user has a password. Can be set via `MYSQL_PASSWORD` environment variable.
- `mysql:tls` - (optional) The TLS configuration. One of false, true, or skip-verify. Defaults to `false`. Can be set via
  `MYSQL_TLS_CONFIG` environment variable.
- `mysql:proxy` - (Optional) Proxy socks url, can also be sourced from `ALL_PROXY` or `all_proxy` environment variables.

## Reasons

This provider bridges [petoju/terraform-provider-mysql](https://github.com/petoju/terraform-provider-mysql) to Pulumi

Currently the official [pulumi-mysql](https://www.pulumi.com/docs/reference/pkg/mysql) uses the [hashicorp/terraform-provider-mysql](https://github.com/hashicorp/terraform-provider-mysql/tree/master) which has been archived in 2021.
Based on [this issue](https://github.com/pulumi/pulumi-mysql/issues/145), Pulumi is still investigating the migration from an archived terraform provider to a community fork.

## Reference

For further information on the resources supported by the Terraform provider, please visit the [documentation](https://registry.terraform.io/providers/petoju/mysql/3.0.63/docs).
