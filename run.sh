#!/bin/bash
TARGET_OS="linux" FILE_EXT="" docker compose run --rm go-builder
# TARGET_OS="linux" FILE_EXT="" docker compose up --build go-builder && docker compose rm -fsv 
./build/forza-telemetry-receiver-linux
