#!/bin/bash
source localhost.env

sleep 20 && /bin/goose -dir "${MIGRATION_DIR}" postgres "${MIGRATION_DSN}" up -v