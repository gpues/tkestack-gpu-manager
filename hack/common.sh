#!/usr/bin/env bash

readonly PACKAGE="tkestack.io/gpu-manager"
readonly BUILD_IMAGE_REPO=plugin-build
readonly LOCAL_OUTPUT_IMAGE_STAGING="${ROOT}/go/images"
readonly IMAGE_FILE=${IMAGE_FILE:-"thomassong/gpu-manager"}
readonly PROTO_IMAGE="proto-generater"

function plugin::cleanup() {
  rm -rf ${ROOT}/go
}

function plugin::cleanup_image() {
  docker rm -vf ${PROTO_IMAGE}
}



function plugin::version::ldflag() {
  local key=${1}
  local val=${2}
  echo "-X ${PACKAGE}/pkg/version.${key}=${val}"
}


function plugin::source_targets() {
  local targets=(
    $(find . -mindepth 1 -maxdepth 1 -not \(        \
        \( -path ./go \) -prune  \
      \))
  )
  echo "${targets[@]}"
}

function plugin::fmt_targets() {
  local targets=(
    $(find . -not \(  \
        \( -path ./go \
        -o -path ./vendor \
        \) -prune \
        \) \
        -name "*.go" \
        -print \
    )
  )
  echo "${targets[@]}"
}
