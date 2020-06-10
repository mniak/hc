`bphc` - Braspag Health Checking Tool
==============================

**Work in Progress!**

## Features

### Liveness check
```bash
bphc alive example.com
``` 

### Health check
```bash
bphc healthy example.com
``` 

## Installation

### [gobinaries.com](gobinaries.com) method

Install to `/usr/local/bin`
```bash
curl -sf https://gobinaries.com/BraspagDevelopers/bphc | sh
```

You can also specify a custom directory where to download the binary file
```bash
# Install on the current directory
curl -sf https://gobinaries.com/BraspagDevelopers/bphc | PREFIX=. sh
```
```bash
# Install on /tmp
curl -sf https://gobinaries.com/BraspagDevelopers/bphc | PREFIX=/tmp sh
```

### `go get` method
```bash
go get -u github.com/BraspagDevelopers/bphc
```

## Usage
**`bphc <command> [flags]`** 
There is no root functionality. You must inform a command in order to perform any relevant action.

### Flags
* **`-v` or `--verbose`:** Enable verbose mode

### Command `alive`
**`bphc alive <Base_URL> [-v|--verbose] [--path path]`**

Sends a `GET` HTTP request to a site in order to check its liveness.
If the site returns a status code in the range 200-299, it will be considered alive.
If the site returns any other status code, the check will fail.

When the check succedes, it will be produce an exit code of `0`. Any failure will produce a difference exit code. Additionaly, there will always be a message in `STDOUT` when the check fails.
#### Arguments
* **`Base_URL`:** The base URL to use in the request. If _scheme_ is absent, `http` will be used for localhost or IPs. `https` is used otherwise.

#### Flags
* **`--path path`:** The path for the liveness endpoint (default `/dfom.htm`)

### Command `healthy`
**`bphc healthy <Base_URL> [-v|--verbose] [--path path]`**

Sends a `GET` HTTP request to a site in order to check its health.
If the site returns a status code in the range 200-299 and the body is in JSON format and the value of the property `IsHealthy` is `true`, the site is considered healhty.
If not, the check will fail.

When the check succedes, it will be produce an exit code of `0`. Any failure will produce a difference exit code. Additionaly, there will always be a message in `STDOUT` when the check fails.
#### Arguments
* **`Base_URL`:** The base URL to use in the request. If _scheme_ is absent, `http` will be used for localhost or IPs. `https` is used otherwise.

#### Flags
* **`--path path`:** The path for the healthcheck endpoint (default `/healthcheck`)

