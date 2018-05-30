# nosce
:postal_horn: Get metadata information of your EC2 instances

## Why

Every AWS instance has access to a special metadata service at the private endpoint http://169.254.169.254, which you can use to retrieve metadata about your instance such as `ami-id` and `local-ipv4`. In order to get these metadata information you must make HTTP calls to the http://169.254.169.254 at different URLs.

`nosce` is a tool that helps you get metadata about your instance without having make any manual HTTP calls and remembering any URLs.

## Install

1. Install it on your go path:

```sh
go get github.com/grvcoelho/nosce
```

## Usage

```
Usage:
  nosce [command]

Available Commands:
  get         Get a piece of metadata information
  help        Help about any command

Flags:
  -e, --endpoint string   Endpoint where the information will be fetched (default "http://169.254.169.254")
  -h, --help              help for nosce

Use "nosce [command] --help" for more information about a command.
```

## Examples

1. Get instance's ami-id:

```sh
$ nosce get ami-id
ami-1f3ca179
```

2. Get instance's local ipv4:

```sh
$ nosce get local-ipv4
10.10.0.1
```
