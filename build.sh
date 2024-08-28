#!/bin/bash
# A script to compile a Go application for multiple OS using Docker Compose

# Declare an associative array with OS as key and file extension as value
declare -A os_list=(
    ["linux"]=""
    ["windows"]=".exe"
    ["darwin"]=""  # macOS typically doesn't use an extension for executables
)

# Iterate over each operating system in the associative array
for os in "${!os_list[@]}"; do
    extension=${os_list[$os]}
    echo "Starting compilation for $os..."
    TARGET_OS=$os FILE_EXT=$extension docker compose run --rm go-builder
done

echo "Compilation complete. Check the /build volume."
