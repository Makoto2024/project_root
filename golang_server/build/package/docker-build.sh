#!/bin/bash
# Assume running in project_root/
docker build -f golang_server/build/package/Dockerfile -t "makoto2024/project_root_foo" .
