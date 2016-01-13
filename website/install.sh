#!/usr/bin/env bash

function install_autotune {
    profile=$1 ; shift

    autotune signature $profile
}

install_autotune "$@"
