#!/usr/bin/env bash

VERSION=0.5.4

function welcome_autotune {
    echo "Welcome to Acksin Autotune ($VERSION) Installation"
    echo "Copyright (c) 2015-2016. Abhi Yerra."
    echo "https://acksin.co/autotune"
    echo ""
    echo "We will be installing Autotune on your system into the following location:"
    echo ""
    echo "  - /usr/local/bin/autotune"
    echo ""
}

function download_autotune {
    curl -s -o /tmp/autotune.tar.gz https://assets.acksin.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz
    cd /tmp && tar zxf autotune.tar.gz
    sudo mv autotune /usr/local/bin/autotune
}

function install_autotune {
    profile=$1 ; shift

    echo "These are the kernel values that will be changing on this machine"
    echo "for the ${profile} profile."
    echo ""
    autotune signature -write=false $profile

    echo ""
    echo "If you okay with setting these values run the following command:"
    echo ""
    echo "  sudo /usr/local/bin/autotune signature $profile"
    echo ""
}

welcome_autotune
download_autotune
install_autotune "$@"
