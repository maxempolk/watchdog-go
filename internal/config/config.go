package config

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"stat_by_sites/domain/endpoint"
)

type Mode string

const (
	ModeDefault Mode = "default"
	ModeFile    Mode = "file"
	ModeSites   Mode = "sites"

	defaultConfigFileName = "data.json"
)

type RuntimeConfig struct {
	Mode      Mode
	FilePath  string
	Endpoints []endpoint.EndpointConfig
}

func ParseRuntimeConfig(args []string) (*RuntimeConfig, error) {
	if len(args) == 0 {
		return parseDefaultMode()
	}

	if args[0] == string(ModeSites) {
		return parseSitesMode(args[1:])
	}

	return parseFileMode(args)
}

func parseDefaultMode() (*RuntimeConfig, error) {
	path := filepath.Join(".", defaultConfigFileName)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"default mode: file %q was not found in current directory",
				defaultConfigFileName,
			)
		}
		return nil, fmt.Errorf("default mode: failed to access %q: %w", path, err)
	}

	return &RuntimeConfig{
		Mode:     ModeDefault,
		FilePath: path,
	}, nil
}

func parseFileMode(args []string) (*RuntimeConfig, error) {
	fs := flag.NewFlagSet("watchdog", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	var filePath string
	fs.StringVar(&filePath, "file", "", "path to config file")

	if err := fs.Parse(args); err != nil {
		return nil, fmt.Errorf("invalid arguments: %w", err)
	}

	rest := fs.Args()
	if filePath == "" {
		if len(rest) > 0 {
			return nil, fmt.Errorf(
				"unknown mode: use no args (default), --file <path>, or sites <host:interval> ...",
			)
		}
		return nil, fmt.Errorf("file mode requires --file <path>")
	}

	if len(rest) > 0 {
		if containsSitesKeyword(rest) {
			return nil, fmt.Errorf("argument conflict: --file cannot be used together with sites mode")
		}
		return nil, fmt.Errorf("file mode does not allow positional arguments: %s", strings.Join(rest, " "))
	}

	if err := validateFileExists(filePath); err != nil {
		return nil, err
	}

	return &RuntimeConfig{
		Mode:     ModeFile,
		FilePath: filePath,
	}, nil
}

func parseSitesMode(items []string) (*RuntimeConfig, error) {
	if len(items) == 0 {
		return nil, fmt.Errorf("sites mode requires at least one <host:interval> entry")
	}

	endpoints := make([]endpoint.EndpointConfig, 0, len(items))
	for _, item := range items {
		if looksLikeFileFlag(item) {
			return nil, fmt.Errorf("argument conflict: --file cannot be used together with sites mode")
		}

		host, intervalToken, err := splitHostInterval(item)
		if err != nil {
			return nil, fmt.Errorf("invalid sites entry %q: %w", item, err)
		}

		interval, err := parseInterval(intervalToken)
		if err != nil {
			return nil, fmt.Errorf("invalid interval in %q: %w", item, err)
		}

		cfg := endpoint.NewDefaultEndpointConfig(host, interval)
		cfg, err = endpoint.NormalizeEndpointConfig(cfg)
		if err != nil {
			return nil, fmt.Errorf("invalid sites entry %q: %w", item, err)
		}

		endpoints = append(endpoints, cfg)
	}

	return &RuntimeConfig{
		Mode:      ModeSites,
		Endpoints: endpoints,
	}, nil
}

func parseInterval(raw string) (time.Duration, error) {
	trimmed := strings.TrimSpace(raw)
	if d, err := time.ParseDuration(trimmed); err == nil {
		if d <= 0 {
			return 0, fmt.Errorf("interval must be > 0")
		}
		return d, nil
	}

	// Numeric values in sites mode are treated as seconds: example.com:10 -> 10s.
	seconds, err := strconv.Atoi(trimmed)
	if err != nil || seconds <= 0 {
		return 0, fmt.Errorf("use Go duration format (e.g. 5s, 1m) or positive integer seconds")
	}

	return time.Duration(seconds) * time.Second, nil
}

func splitHostInterval(value string) (string, string, error) {
	idx := strings.LastIndex(value, ":")
	if idx <= 0 || idx == len(value)-1 {
		return "", "", fmt.Errorf("expected host:interval")
	}

	host := strings.TrimSpace(value[:idx])
	interval := strings.TrimSpace(value[idx+1:])
	if host == "" || interval == "" {
		return "", "", fmt.Errorf("host and interval are required")
	}

	return host, interval, nil
}

func validateFileExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file mode: file %q does not exist", path)
		}
		return fmt.Errorf("file mode: failed to access %q: %w", path, err)
	}

	return nil
}

func containsSitesKeyword(args []string) bool {
	for _, arg := range args {
		if arg == string(ModeSites) {
			return true
		}
	}
	return false
}

func looksLikeFileFlag(arg string) bool {
	return arg == "--file" ||
		arg == "-file" ||
		strings.HasPrefix(arg, "--file=") ||
		strings.HasPrefix(arg, "-file=")
}
