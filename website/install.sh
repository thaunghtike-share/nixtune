#!/usr/bin/env bash

VERSION=0.5.2

function welcome_autotune {
    echo "Welcome to Anatma Autotune ($VERSION) Install"
    echo "Copyright (c) 2015-2016. Abhi Yerra."
    echo "https://anatma.co/autotune"
    echo ""
    echo "We will be installing Autotune on your system into the following path:"
    echo "  - /usr/local/bin"
    echo ""
}

function download_autotune {
    curl -o /tmp/autotune.tar.gz https://assets.anatma.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz
    cd /tmp && tar zxvf autotune.tar.gz
    sudo mv autotune /usr/local/bin/autotune
}

function install_autotune {
    profile=$1 ; shift

    echo "Since we are installing from the web Anatma Autotune"
    echo "WILL NOT update the system values. It will only display"
    echo "what values will change."
    echo ""
    echo "Verify that the values being changed are okay."
    echo "Run it again with: \"autotune signature $profile\""
    echo ""
    autotune signature -write=false $profile
}

welcome_autotune
download_autotune
install_autotune "$@"
