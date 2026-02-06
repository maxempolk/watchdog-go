# Watchdog TUI

Terminal UI for monitoring HTTP endpoints.

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

```json
{
  "resources": [
    { "url": "https://example.com", "interval": "5s" },
    { "url": "https://api.example.com", "interval": "30s" }
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
