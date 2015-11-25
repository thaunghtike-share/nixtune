#!/usr/bin/env bash


echo "fs.file-max = 100000" >> /etc/sysctl.conf
echo "* - nofile unlimited" >> /etc/security/limits.d/00_anatma_knight_limits.conf

sysctl -p
