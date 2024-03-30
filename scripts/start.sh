#!/bin/bash

bash ./scripts/permissions.sh
if [ $? -ne 0 ]; then
    echo "start.sh : Error : permissions.sh failed"
    exit 1
fi

bash ./scripts/env_validation.sh
if [ $? -ne 0 ]; then
    echo "start.sh : Error : env_validation.sh failed"
    exit 1
fi

docker-compose up -d --build