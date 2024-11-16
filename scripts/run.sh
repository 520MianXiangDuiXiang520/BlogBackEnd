#!/bin/bash

function prepare_check() {
    cfgPath="/data/juneblog/config"
    logPath="/data/juneblog/log"
    if [ -v "$CFG_PATH" ]; then
         cfgPath="$CFG_PATH"
    fi

    if [ -v "$LOG_PATH" ]; then
         logPath="$LOG_PATH"
    fi
    mkdir -p "$cfgPath"
    mkdir -p "$logPath"
}

prepare_check
/app/juneblog

