#!/usr/bin/env bash

VERSION=0.0.0

function download_autotune {
    sudo https://assets.anatma.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz
}

function install_autotune {
    profile=$1 ; shift

    autotune signature $profile
}

install_autotune "$@"
