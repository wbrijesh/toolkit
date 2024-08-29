# Toolkit

Toolkit is a Go library providing utilities for API development.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)

## Installation

To install Toolkit, use the following command:

```bash
go get brijesh.dev/toolkit
```

## Usage

Here's a quick example of how to use Toolkit:

```go
package main

import (
 "net/http"
 "time"

 "brijesh.dev/toolkit/middleware"
 "brijesh.dev/toolkit/router"
)

type response struct {
 Message   string      `json:"message"`
 Error     error       `json:"error"`
 Data      interface{} `json:"data"`
 RequestID string      `json:"request_id"`
}

func main() {
 r := router.NewRouter()

 r.Use(middleware.Logger)
 r.Use(middleware.RequestIDMiddleware)
 r.Use(middleware.RateLimit(2, time.Second))

 r.GET("/health", healthCheckHandler)

 http.ListenAndServe(":8080", r)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
 response := response{
  Message:   "Okay",
  Error:     nil,
  Data:      nil,
  RequestID: middleware.GetRequestID(r),
 }
 router.SendResponse(w, http.StatusOK, response)
}
```

## Features

- **Router**: A lightweight wrapper around net/http, providing an intuitive API for defining routes and middleware.
- **Middleware**:
  - Logger: Logs incoming requests with details like method, path, and response time.
  - Rate Limiter: Implements rate limiting to protect your API from abuse.
  - Request ID: Assigns a unique ID to each request for easier tracking and debugging.
- **BUID (Brijesh's Unique Identifier)**: A custom unique identifier generator, similar to UUID but with its own algorithm.
- **Benchmarking**: A package to easily benchmark your functions and measure performance.
