#!/usr/bin/env bash

PRODUCT=strum
VERSION=0.3.0

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

    curl -s --data "v=1&tid=UA-75403807-1&cid=2d7de14a2283b05c956524b2878ab7fa23c9b179&t=event&ec=Download&ea=Download%20STRUM&el=Validation%20Download%20STRUM%20${VERSION}%20Curl" https://www.google-analytics.com/collect > /dev/null
}

welcome_strum
download_strum
