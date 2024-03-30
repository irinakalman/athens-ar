#!/bin/bash

ENV_FILE=".env"

# check if ENV_FILE not exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: $ENV_FILE does not exist."
    exit 1
fi

# check if ENV_FILE doesn't have content
if [ ! -s "$ENV_FILE" ]; then
    echo "Error: $ENV_FILE is empty."
    exit 1
fi

# check values
required_vars=("DB_POSTGIS_DB" "DB_POSTGIS_USER" "DB_POSTGIS_PASSWORD")
for var in "${required_vars[@]}"; do
    if ! grep -q "^$var=" "$ENV_FILE"; then
        echo "Error: $var is not set in $ENV_FILE"
        exit 1
    fi

    value=$(grep "^$var=" "$ENV_FILE" | cut -d'=' -f2-)
    if [ -z "${value}" ]; then
        echo "Error: $var is empty in $ENV_FILE"
        exit 1
    fi
done

echo "env_validation.sh : Success"
exit 0