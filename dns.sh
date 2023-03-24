#!/bin/bash

# Set the IP address for the kiosk.local domain name
ip_address="127.0.0.1"

# Check if the kiosk.local entry already exists in the hosts file
if grep -q "kiosk.local" /etc/hosts; then
  echo "The kiosk.local entry already exists in the hosts file."
else
  # Add the kiosk.local entry to the hosts file
  echo "$ip_address kiosk.local" | sudo tee -a /etc/hosts > /dev/null
  echo "The kiosk.local entry has been added to the hosts file."
fi