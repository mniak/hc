`hc` - HealthCheck Tool
==============================

**Work in Progress!**

## Features

### Liveness check
`hc`

### Health check

## Installation

### gobinaries.com method

Install `hc` to `/usr/local/bin`
```bash
curl -sf https://gobinaries.com/mniak/hc | sh
```

You can also specify a custom directory where to download the binary file
```bash
# Install on the current directory
curl -sf https://gobinaries.com/mniak/hc | PREFIX=. sh
```
```bash
# Install on /tmp
curl -sf https://gobinaries.com/mniak/hc | PREFIX=/tmp sh
```


### `go get` method
```bash
go get github.com/mniak/hc
```

## Usage
**`hc command [flags]`** 
There is no root functionality. You must inform a command in order to perform any relevant action.


### Command `alive`
Sends a `GET` HTTP request to a site in order to check its liveness.
If the site returns a status code in the range 200-299, it will be considered alive.
If the site returns any other status code, the check will fail.

When the check succedes, it will be produce an exit code of `0`. Any failure will produce a difference exit code. Additionaly, there will always be a message in `STDOUT` when the check fails.

> **TODO:** document the flags