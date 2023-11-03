#!/bin/bash



OS_TOKEN="<bot_token>"

OS_DATE=$(date)

OS_HOSTNAME=$(hostname)

#$USER

OS_IP=$(env | grep SSH_CLIENT | cut -d '=' -f 2 | cut -d ' ' -f 1)

OS_ID=$(curl -s "https://api.telegram.org/bot$OS_TOKEN/getUpdates" | jq -r '.result[0].message.chat.id')



curl -s -X POST -H 'Content-Type: application/json' -d "{\"chat_id\": $OS_ID, \"text\": \"User login detected:\n$OS_DATE\n$OS_HOSTNAME\n$OS_IP\n$USER\", \"disable_notification\": false}" "https://api.telegram.org/bot$OS_TOKEN/sendMessage"  > /dev/null 2>&1
