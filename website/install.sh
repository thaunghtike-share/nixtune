#!/usr/bin/env bash

VERSION=0.5.1

function download_autotune {
    curl -o /tmp/autotune.tar.gz https://assets.anatma.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz
    cd /tmp && tar zxvf autotune.tar.gz
    sudo mv autotune /usr/local/bin/autotune
}

function install_autotune {
    profile=$1 ; shift

    echo "Autotune WILL NOT update the values."
    echo "First verify that the values being changed are okay."
    echo "Run it again with: \"autotune signature $profile\""
    autotune signature -write=false $profile
}

download_autotune
install_autotune "$@"
