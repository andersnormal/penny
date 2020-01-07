[![Build Status](https://travis-ci.org/andersnormal/penny.svg?branch=master)](https://travis-ci.org/axelspringer/penny)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Penny

Penny is a tool to run commands in K/V pre-configured environments. It is especially useful in [Docker](https://docker.io) containers to populate the environment with secrets.

The supported K/V providers are

* [System Manager Parameter Store](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-paramstore.html)
* libkv (comming soon)

## Getting Started

The most simple case is, that you want to execute a command with all the available parameters from the store. Folders are converted to `_` in the environment variable.

> it is recommended to use `PENNY_PATH` to configure the path to fetch
> all parameters are prefixed with `SSM` when not otherwise specified

```
penny --path /example run printenv
```

This would suggest that there is a key in `/example/parameter`.

## Usage

You can use `--help` to print out all the available options.

### --force

> this is especially helpful in the case of testing

This forces to run the process, without be able to fetch an env from the parameter store.

### PENNY_PATH (--path)

Configures the path in the System Manager Parameter Store from which to fetch the configs.

### PREFIX (--prefix)

Configures the prefix for the parameters that are in the path (e.g. SSM_MY_SECRET)

## License
[Apache-2.0](/LICENSE)
