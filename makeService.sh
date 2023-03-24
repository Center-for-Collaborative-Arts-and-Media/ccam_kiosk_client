#!/bin/bash

# Set the name of the Docker Compose service you want to run at startup
SERVICE_NAME="kiosk-web"

# Set the path to the Docker Compose file
COMPOSE_FILE_PATH="~/ccam_kiosk_client/docker-compose.yml"

# Set the path to the systemd service file
SYSTEMD_SERVICE_PATH="/etc/systemd/system/docker-${SERVICE_NAME}.service"

# Create the systemd service file
#	tls /etc/caddy/certs/kiosk.local.pem /etc/caddy/certs/kiosk.local-key.pem

cat << EOF > ${SYSTEMD_SERVICE_PATH}
[Unit]
Description=Docker Compose ${SERVICE_NAME} service
Requires=docker.service
After=docker.service

[Service]
Type=oneshot
RemainAfterExit=yes
WorkingDirectory=$(dirname ${COMPOSE_FILE_PATH})
ExecStart=/usr/bin/docker-compose -f ${COMPOSE_FILE_PATH} up -d ${SERVICE_NAME}
ExecStop=/usr/bin/docker-compose -f ${COMPOSE_FILE_PATH} stop ${SERVICE_NAME}
ExecReload=/usr/bin/docker-compose -f ${COMPOSE_FILE_PATH} up -d ${SERVICE_NAME}

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd and start the service
systemctl daemon-reload
systemctl enable docker-${SERVICE_NAME}.service
systemctl start docker-${SERVICE_NAME}.service

echo "Service ${SERVICE_NAME} enabled to start on boot."
