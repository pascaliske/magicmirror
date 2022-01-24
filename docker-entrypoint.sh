#!/bin/sh
# -*- coding: utf-8 -*-

# ensure default ids
PUID=${PUID:-911}
PGID=${PGID:-911}

# modify container user
groupmod -o -g "$PGID" unknown
usermod -o -u "$PUID" unknown

# exec container command
su-exec unknown:unknown "$@"
