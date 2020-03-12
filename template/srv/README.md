# Template Service

This is the Template service

Generated with

```
micro new github.com/eopenio/examples/template/srv --namespace=eopenio.micro --alias=template --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: eopenio.micro.srv.template
- Type: srv
- Alias: template

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./template-srv
```

Build a docker image
```
make docker
```