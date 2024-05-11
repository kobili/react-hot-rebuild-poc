# React Hot Rebuild POC
This project demonstrates the use of the `--watch` flag when using Vite's build system
to trigger automatic rebuilds of a React app when its source file changes.

The static files are then served by a Go HTTP server.

## Requirements
- Docker
- Go v1.21.1+, but that's just what I developed with

## Setup
- Build  and watch the UI by running
    ```shell
    make build-ui
    ```
- Start the HTTP server
    ```shell
    make start-server
    ```
