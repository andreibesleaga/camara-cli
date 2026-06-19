# Camara CLI

The official CLI for the Camara REST API.

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Go

To test or install the CLI locally, you need [Go](https://go.dev/doc/install) version 1.22 or later installed.

```sh
go install 'github.com/andreibesleaga/camara-cli/cmd/camara@latest'
```

Once you have run `go install`, the binary is placed in your Go bin directory:

- **Default location**: `$HOME/go/bin` (or `$GOPATH/bin` if GOPATH is set)
- **Check your path**: Run `go env GOPATH` to see the base directory

If commands aren't found after installation, add the Go bin directory to your PATH:

```sh
# Add to your shell profile (.zshrc, .bashrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"
```

<!-- x-release-please-end -->

### Running Locally

After cloning the git repository for this project, you can use the
`scripts/run` script to run the tool locally:

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
camara [resource] <command> [flags...]
```

```sh
camara customerinsights:scoring retrieve \
  --bearer-token 'My Bearer Token'
```

For details about specific commands, use the `--help` flag.

### Environment variables

| Environment variable                                      | Required |
| --------------------------------------------------------- | -------- |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_BEARER_TOKEN`                                     | yes      |
| `CAMARA_DEVICE_LOCATION_NOTIFICATIONS_API_KEY`            | yes      |
| `CAMARA_NOTIFICATIONS_API_KEY`                            | yes      |
| `CAMARA_POPULATION_DENSITY_DATA_NOTIFICATIONS_API_KEY`    | yes      |
| `CAMARA_REGION_DEVICE_COUNT_NOTIFICATIONS_API_KEY`        | yes      |
| `CAMARA_CONNECTIVITY_INSIGHTS_NOTIFICATIONS_API_KEY`      | yes      |
| `CAMARA_SIM_SWAP_NOTIFICATIONS_API_KEY`                   | yes      |
| `CAMARA_DEVICE_ROAMING_STATUS_NOTIFICATIONS_API_KEY`      | yes      |
| `CAMARA_DEVICE_REACHABILITY_STATUS_NOTIFICATIONS_API_KEY` | yes      |
| `CAMARA_CONNECTED_NETWORK_TYPE_NOTIFICATIONS_API_KEY`     | yes      |

### Global flags

- `--bearer-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--customer-insights-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-swap-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--kyc-age-verification-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--kyc-fill-in-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--kyc-match-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--tenure-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--number-recycling-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--otp-validation-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--call-forwarding-signal-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-location-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--population-density-data-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--region-device-count-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--web-rtc-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--connectivity-insights-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--quality-on-demand-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-identifier-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--sim-swap-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-roaming-status-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-reachability-status-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--connected-network-type-token` (can also be set with `CAMARA_BEARER_TOKEN` env var)
- `--device-location-notifications-api-key` (can also be set with `CAMARA_DEVICE_LOCATION_NOTIFICATIONS_API_KEY` env var)
- `--notifications-api-key` (can also be set with `CAMARA_NOTIFICATIONS_API_KEY` env var)
- `--population-density-data-notifications-api-key` (can also be set with `CAMARA_POPULATION_DENSITY_DATA_NOTIFICATIONS_API_KEY` env var)
- `--region-device-count-notifications-api-key` (can also be set with `CAMARA_REGION_DEVICE_COUNT_NOTIFICATIONS_API_KEY` env var)
- `--connectivity-insights-notifications-api-key` (can also be set with `CAMARA_CONNECTIVITY_INSIGHTS_NOTIFICATIONS_API_KEY` env var)
- `--sim-swap-notifications-api-key` (can also be set with `CAMARA_SIM_SWAP_NOTIFICATIONS_API_KEY` env var)
- `--device-roaming-status-notifications-api-key` (can also be set with `CAMARA_DEVICE_ROAMING_STATUS_NOTIFICATIONS_API_KEY` env var)
- `--device-reachability-status-notifications-api-key` (can also be set with `CAMARA_DEVICE_REACHABILITY_STATUS_NOTIFICATIONS_API_KEY` env var)
- `--connected-network-type-notifications-api-key` (can also be set with `CAMARA_CONNECTED_NETWORK_TYPE_NOTIFICATIONS_API_KEY` env var)
- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

### Passing files as arguments

To pass files to your API, you can use the `@myfile.ext` syntax:

```bash
camara <command> --arg @abe.jpg
```

Files can also be passed inside JSON or YAML blobs:

```bash
camara <command> --arg '{image: "@abe.jpg"}'
# Equivalent:
camara <command> <<YAML
arg:
  image: "@abe.jpg"
YAML
```

If you need to pass a string literal that begins with an `@` sign, you can
escape the `@` sign to avoid accidentally passing a file.

```bash
camara <command> --username '\@abe'
```

#### Explicit encoding

For JSON endpoints, the CLI tool does filetype sniffing to determine whether the
file contents should be sent as a string literal (for plain text files) or as a
base64-encoded string literal (for binary files). If you need to explicitly send
the file as either plain text or base64-encoded data, you can use
`@file://myfile.txt` (for string encoding) or `@data://myfile.dat` (for
base64-encoding). Note that absolute paths will begin with `@file://` or
`@data://`, followed by a third `/` (for example, `@file:///tmp/file.txt`).

```bash
camara <command> --arg @data://file.txt
```

## Linking different Go SDK versions

You can link the CLI against a different version of the Camara Go SDK
for development purposes using the `./scripts/link` script.

To link to a specific version from a repository (version can be a branch,
git tag, or commit hash):

```bash
./scripts/link github.com/org/repo@version
```

To link to a local copy of the SDK:

```bash
./scripts/link ../path/to/camara-go
```

If you run the link script without any arguments, it will default to `../camara-go`.
