`bphc` - Braspag Health Checking Tool
==============================

**Work in Progress!**

## Features

### Liveness check
```bash
bphc alive example.com
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
go get github.com/BraspagDevelopers/bphc
```

## Usage
**`bphc <command> [flags]`** 
There is no root functionality. You must inform a command in order to perform any relevant action.

### Command `alive`
**`bphc alive <Base_URL> [-v|--verbose] [--path path]`**

Sends a `GET` HTTP request to a site in order to check its liveness.
If the site returns a status code in the range 200-299, it will be considered alive.
If the site returns any other status code, the check will fail.

When the check succedes, it will be produce an exit code of `0`. Any failure will produce a difference exit code. Additionaly, there will always be a message in `STDOUT` when the check fails.
#### Arguments
* **`Base_URL`:** The base URL to use in the request. If _scheme_ is absent, `https://` will be used.
#### Flags
* **`-v` or `--verbose`:** Enable verbose mode