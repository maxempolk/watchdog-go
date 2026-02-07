# Watchdog TUI

Terminal UI for monitoring HTTP endpoints.

## Features

- Three launch modes: `default`, `--file`, `sites`
- Configurable request method per endpoint: `HEAD`, `GET`, `POST`
- Custom request headers and request body
- Response validation with `expected_status` and `expected_body_contains`
- Per-endpoint timeout and redirect policy (`follow_redirects`)
- Logs panel with scroll and freeze/unfreeze
- State panel with last status and result for all endpoints

## Quick Start

1. Build or run directly:
   - `go run .`
   - or `go build -o watchdog && ./watchdog`
2. Use one of the launch modes below.

## Launch Modes

### 1) Default mode

Runs with no args and loads `./data.json`.

```bash
go run .
```

If `data.json` is missing, startup fails with an error.

### 2) File mode

Loads config from an explicit file path.

```bash
go run . --file ./data1.json
```

### 3) Sites mode

Pass endpoints directly in CLI as `host:interval`.

```bash
go run . sites example.com:5s api.local:10 127.0.0.1:1m
```

`10` means `10s`.

## Config Format

Minimal config:

```json
{
  "resources": [
    { "url": "https://example.com", "interval": "5s" },
    { "url": "https://api.example.com", "interval": "30s" }
  ]
}
```

Extended config (optional fields):

```json
{
  "resources": [
    {
      "url": "https://api.example.com/health",
      "interval": "10s",
      "method": "GET",
      "headers": {
        "Accept": "application/json",
        "Authorization": "Bearer token"
      },
      "body": "",
      "expected_status": 200,
      "expected_body_contains": "\"ok\":true",
      "timeout": "5s",
      "follow_redirects": false
    }
  ]
}
```

## Controls

- `Q` quit
- `R` manual refresh
- `L` show/hide logs
- `F` freeze/unfreeze logs
- `S` show/hide state panel
- `Up/Down` or `J/K` scroll logs

## Notes

Educational project, not production-ready.
