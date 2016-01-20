#!/usr/bin/env bash

VERSION=0.5.1

function download_autotune {
    curl -o /tmp/autotune.tar.gz https://assets.anatma.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz
    cd /tmp && tar jxvf autotune.tar.gz
    sudo mv autotune /usr/local/bin/autotune
}

function install_autotune {
    profile=$1 ; shift

    autotune signature $profile
}

download_autotune
install_autotune "$@"
