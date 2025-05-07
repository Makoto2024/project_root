#!/bin/bash
bazel run //golang_server/cmd/server:server_image_tar
bazel run //cc_server/server:server_image_tar
