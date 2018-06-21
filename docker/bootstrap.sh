#!/bin/bash

EXTRA_ARGS=$EXTRA_ARGS
if [ $LISTENPORT ]; then
    EXTRA_ARGS=$EXTRA_ARGS' -ListenPort='$LISTENPORT
fi

if [ $DBSERVER ]; then
    EXTRA_ARGS=$EXTRA_ARGS' -DBServer='$DBSERVER
fi

if [ $DBNAME ]; then
    EXTRA_ARGS=$EXTRA_ARGS' -DBName='$DBNAME
fi

if [ $ACCOUNT ]; then
    EXTRA_ARGS=$EXTRA_ARGS' -Account='$ACCOUNT
fi

if [ $PASSWORD ]; then
    EXTRA_ARGS=$EXTRA_ARGS' -Password='$PASSWORD
fi

echo $EXTRA_ARGS

/var/app/magicCenter $EXTRA_ARGS "$@"
