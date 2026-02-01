# Watchdog TUI

**Watchdog TUI** is an educational TUI project written in Go for real-time monitoring of HTTP endpoints availability.

The project displays:
- current endpoint status (HTTP status code)
- response latency
- recent health-check history (trend)
- aggregated system health summary
- recent event logs

The interface is implemented as a terminal UI (TUI) using **Bubble Tea** and **Lip Gloss**.

---

## Features

- Periodic health checks for endpoints
- Immediate data refresh on startup
- Visual indicators for status and trends
- Tabular view of monitored endpoints
- In-app log viewer
- Keyboard-driven controls

---

## Tech Stack

- Go
- Bubble Tea (TUI framework)
- Lip Gloss (styling)
- Clean Architecture (UI / domain / repository)

---

## Project Status

⚠️ **Educational Project**

This project was created for learning purposes:
- learning Go
- building TUI applications
- understanding Bubble Tea’s event-driven model
- practicing separation of UI, domain, and infrastructure layers

It is **not intended for production use**.

---

## Controls

- `Q` — quit
- `R` — manual refresh
- `L` — show / hide logs

---

## License

MIT