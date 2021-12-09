#!/usr/bin/bash

# Quit on error.
set -e
# Treat undefined variables as errors.
set -u


function main {
    local gandalf_uid="${1:-}"
    local gandalf_gid="${2:-}"

    # Change the uid
    if [[ -n "${gandalf_uid:-}" ]]; then
        usermod -u "${gandalf_uid}" root
    fi
    # Change the gid
    if [[ -n "${gandalf_gid:-}" ]]; then
        groupmod -g "${gandalf_gid}" root
    fi

    # Setup permissions on the run directory where the sockets will be
    # created, so we are sure the app will have the rights to create them.

    # Make sure the folder exists.
    mkdir /var/run/sockets
    # Set owner.
    chown root:root /var/run/sockets
    # Set permissions.
    chmod u=rwX,g=rwX,o=--- /var/run/sockets
}


main "$@"