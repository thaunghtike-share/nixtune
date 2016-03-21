#!/usr/bin/env bash

VERSION=0.1.0

function welcome_strum {
    echo "Welcome to Acksin STRUM ($VERSION) Installation"
    echo "Copyright (c) 2016. Acksin, LLC."
    echo "https://acksin.co/strum"
    echo ""
}

function download_strum {
    echo "Downloading for $(uname)-$(uname -i)"
    curl -s -o /tmp/strum.tar.gz https://assets.acksin.com/strum/${VERSION}/strum-${VERSION}-$(uname)-$(uname -i).tar.gz
    tar zxf strum.tar.gz
}

welcome_strum
download_strum
