# go-cli

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9a71350a5de64095a7f175170fc81137)](https://app.codacy.com/gh/edwinvautier/go-cli?utm_source=github.com&utm_medium=referral&utm_content=edwinvautier/go-cli&utm_campaign=Badge_Grade_Settings)
[![Go](https://github.com/edwinvautier/go-cli/actions/workflows/go.yml/badge.svg)](https://github.com/edwinvautier/go-cli/actions/workflows/go.yml)
[![CodeQL](https://github.com/edwinvautier/go-cli/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/edwinvautier/go-cli/actions/workflows/codeql-analysis.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![version](https://img.shields.io/badge/version-0.0.8-red)
[![Go Reference](https://pkg.go.dev/badge/github.com/edwinvautier/go-cli.svg)](https://pkg.go.dev/github.com/edwinvautier/go-cli)
[![codecov](https://codecov.io/gh/edwinvautier/go-cli/branch/main/graph/badge.svg?token=1USTLF2NA0)](https://codecov.io/gh/edwinvautier/go-cli)

A CLI to initialize and work on go projects, mainly designed for API's.

## Install CLI

>💡 You need to have go installed correctly on your machine. More informations are available in the wiki.

Install the CLI by running :

```sh
go get github.com/edwinvautier/go-cli
```

Run `go-cli -h` in order to know more about commands.

## Initialize a project

You can initialize a project in your working directory by running the `create` command.

```sh
go-cli create
# or with app name
go-cli create my-app-name
```

The CLI will eventually ask you your **git username**, the **DB management system** you'd like to use and if you want to **dockerize** the application or not.

## Install a bundle

You can install bundles by using the install command of the CLI.
This command will look for a bundle located inside the bundles folder of the repository and install it.

```sh
go-cli install authenticator
```

### Bundle API

Each bundle should have the following elements :

- `templates` folder
- `install.sh` script
- `README.md` file to explain how bundle works and how to use it

The templates part must follow the same filetree as the project that is created.

## Make new model

With **go-cli**, you can use the `make model` command. It will create a new model and repository file with fields of your choice !

```sh
go-cli make model modelYouWant
```

## Make CRUD

**go-cli** can generate your controllers, to do so, you just have to use the `make crud` command :

```sh
go-cli make crud entityName
```

**go-cli** will eventually asks you to run the `go-cli update` command, that reads the models files, and parse their fields to the config.
