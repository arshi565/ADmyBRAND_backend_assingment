package main

import (
    "os"
)

// Config is a struct that contains the configuration options for the program.
type Config struct {
    // Port is the port that the server listens on.
    Port string

    // LogLevel is the logging level for the server.
    LogLevel string

    // MaxRequests is the maximum number of requests that the server can handle simultaneously.
    MaxRequests int
}

// GetConfig reads the values of the configuration options from environment variables.
func GetConfig() (*Config, error) {
    cfg := &Config{
        Port:        "8080",
        LogLevel:    "info",
        MaxRequests: 10,
    }

    // Read the port from the PORT environment variable, if it is set.
    if port := os.Getenv("PORT"); port != "" {
        cfg.Port = port
    }

    // Read the logging level from the LOG_LEVEL environment variable, if it is set.
    if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
        cfg.LogLevel = logLevel
    }

    // Read the max requests from the MAX_REQUESTS environment variable, if it is set.
    if maxRequests := os.Getenv("MAX_REQUESTS"); maxRequests != "" {
        max, err := strconv.Atoi(maxRequests)
        if err != nil {
            return nil, fmt.Errorf("invalid value for MAX_REQUESTS: %s", maxRequests)
        }
        cfg.MaxRequests = max
    }

    return cfg, nil
}
