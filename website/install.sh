#!/usr/bin/env bash

VERSION=0.5.4

function welcome_strum {
    echo "Welcome to Acksin STRUM ($VERSION) Installation"
    echo "Copyright (c) 2016. Acksin, LLC."
    echo "https://acksin.co/strum"
    echo ""
    echo "STRUM will be installed in the following location:"
    echo ""
    echo "    /usr/local/bin/strum"
    echo ""
}

function download_strum {
    curl -s -o /tmp/strum.tar.gz https://assets.acksin.com/strum/${VERSION}/strum-${VERSION}-$(uname)-$(uname -i).tar.gz
    cd /tmp && tar zxf strum.tar.gz
    sudo mv strum /usr/local/bin/strum
}

welcome_strum
download_strum
install_strum "$@"
