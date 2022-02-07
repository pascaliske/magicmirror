#!/bin/sh
# -*- coding: utf-8 -*-

# modify container user ids
groupmod -o -g "${PUID:-911}" unknown >/dev/null 2>&1
usermod -o -u "${PGID:-911}" unknown >/dev/null 2>&1

# exec container command
su-exec unknown:unknown "$@"
