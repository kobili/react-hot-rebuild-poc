# React Hot Rebuild POC
This project demonstrates the use of [`watchfiles`](https://pypi.org/project/watchfiles/)
to trigger automatic rebuilds of a React app when its source file changes.

The static files are then served by a Go HTTP server.

## Requirements
- Python v3.8 - 3.12 per watchfiles requirement
- Go v1.21.1+, but that's just what I developed with

## Setup
- install `watchfiles` using pip.
    ```shell
    pip install watchfiles
    ```
- Build and start watching by running
    ```shell
    ./watchbuild.sh
    ```
- Start the HTTP server
    ```shell
    make start-server
    ```
