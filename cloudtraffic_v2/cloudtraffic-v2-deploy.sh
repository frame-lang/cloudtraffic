#!/usr/bin/sh

PROJECT_DIR="/root/cloudtraffic"
BACKEND_DIR="/root/cloudtraffic/cloudtraffic_v2/backend-go"
FRONTEND_DIR="/root/cloudtraffic/cloudtraffic_v2/frontend-react"

cd $PROJECT_DIR
git pull origin main

cd $BACKEND_DIR

echo "Deleting previous backend build...." 
rm -rf cloudtraffic_v2
echo "Creating BE Build..."
go build

SERVICEFILE="/lib/systemd/system/trafficlightbackendv2.service"

if [ -f "$SERVICEFILE" ]; then
echo "Restarting backend service..."
    systemctl daemon-reload
    service trafficlightbackendv2 restart
else
    echo "Creating backend service..."
    touch $SERVICEFILE
    echo "
    [Unit]
    Description=Traffic Light Backend V2
    [Service]
    Type=simple
    Restart=always
    Environment=GOOGLE_APPLICATION_CREDENTIALS=$BACKEND_DIR/pubsub-system-key.json
    RestartSec=5s
    WorkingDirectory=$BACKEND_DIR
    ExecStart=$BACKEND_DIR/cloudtraffic_v2
    [Install]
    WantedBy=multi-user.target" | tee -a $SERVICEFILE > /dev/null
    systemctl daemon-reload
    service trafficlightbackend restart
    systemctl enable trafficlightbackend.service
fi

cd $FRONTEND_DIR
echo "Deleting previous frontend build...."
rm -rf build
echo "Installing Dependencies..."
npm install
echo "Creating Build..."
npm run build

echo "Running FE as a service..."
pm2 restart v2 
pm2 save

echo "Cloud Traffic V2 Deployment is Done."
exit 1
