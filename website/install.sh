#!/usr/bin/env bash

PRODUCT=strum
VERSION=0.2.2

function welcome_strum {
    echo "Welcome to Acksin STRUM ($VERSION) Installation"
    echo "Copyright (c) 2016. Acksin, LLC."
    echo "https://www.acksin.com/strum"
    echo ""
}

function download_strum {
    echo "Downloading STRUM for $(uname)-$(uname -i)"
    curl -s -o strum.tar.gz https://assets.acksin.com/${PRODUCT}/${VERSION}/${PRODUCT}-$(uname)-$(uname -i)-${VERSION}.tar.gz
    tar zxf strum.tar.gz
    echo "STRUM is installed as ${PWD}/strum"
}

welcome_strum
download_strum
