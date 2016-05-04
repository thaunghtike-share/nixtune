#!/usr/bin/env bash

VERSION=0.5.4

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
    curl -s -o autotune.tar.gz https://assets.acksin.com/autotune/${VERSION}/autotune-${VERSION}-$(uname -i).tar.gz
    tar zxf autotune.tar.gz
    rm autotune.tar.gz
}

function install_autotune {
    profile=$1 ; shift

    echo "These are the kernel values that will be changing on this machine"
    echo "for the ${profile} profile."
    echo ""

    autotune sig  $profile

    # echo ""
    # echo "If you okay with setting these values run the following command:"
    # echo ""
    # echo "  sudo /usr/local/bin/autotune signature $profile"
    # echo ""
}

welcome_autotune
download_autotune
install_autotune "$@"
