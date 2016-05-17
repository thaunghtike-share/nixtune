#!/usr/bin/env bash

PRODUCT=autotune
VERSION=0.6.2

function welcome_autotune {
    echo "Welcome to Acksin Autotune ($VERSION) Installation"
    echo "Copyright (c) 2016. Acksin, LLC."
    echo "https://acksin.co/autotune"
    echo ""
    echo "Autotune will be installed in the following location:"
    echo ""
    echo "    ${PWD}"
    echo ""
}

function download_autotune {
    curl -s -o ${PRODUCT}.tar.gz https://assets.acksin.com/${PRODUCT}/${VERSION}/${PRODUCT}-$(uname)-$(uname -i)-${VERSION}.tar.gz
    tar zxf ${PRODUCT}.tar.gz
    rm ${PRODUCT}.tar.gz
}

function install_autotune {
    profile=$1 ; shift

    curl -s --data "v=1&tid=UA-75403807-1&cid=2d7de14a2283b05c956524b2878ab7fa23c9b179&t=event&ec=Download&ea=Autotune%20Curl" https://www.google-analytics.com/collect > /dev/null

    echo "These are the kernel values that will be changing on this machine"
    echo "for the ${profile} profile."
    echo ""

    ./autotune sig  $profile
}

welcome_autotune
download_autotune
install_autotune "$@"
