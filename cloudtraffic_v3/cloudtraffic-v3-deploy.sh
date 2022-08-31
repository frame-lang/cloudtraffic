#!/usr/bin/sh

PROJECT_DIR="/root/cloudtraffic"
BACKEND_DIR="/root/cloudtraffic/cloudtraffic_v3/utils-service"
FRONTEND_DIR="/root/cloudtraffic/cloudtraffic_v2/frontend-react"
UTILS_SERVICE_BUILD="cloudtraffic-v3-utils-service"

cd $PROJECT_DIR
git pull origin main

cd $BACKEND_DIR

echo "Deleting previous backend build...." 
rm -rf $UTILS_SERVICE_BUILD
echo "Creating BE Build..."
go build -o $UTILS_SERVICE_BUILD

SERVICEFILE="/lib/systemd/system/trafficlightbackendv3.service"

if [ -f "$SERVICEFILE" ]; then
echo "Restarting backend service..."
    systemctl daemon-reload
    service trafficlightbackendv3 restart
else
    echo "Creating backend service..."
    touch $SERVICEFILE
    echo "
    [Unit]
    Description=Traffic Light Backend V3
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    WorkingDirectory=$BACKEND_DIR
    ExecStart=$BACKEND_DIR/$UTILS_SERVICE_BUILD
    [Install]
    WantedBy=multi-user.target" | tee -a $SERVICEFILE > /dev/null
    systemctl daemon-reload
    service trafficlightbackendv3 restart
    systemctl enable trafficlightbackendv3.service
fi

cd $FRONTEND_DIR
echo "Deleting previous frontend build...."
rm -rf build
echo "Installing Dependencies..."
npm install
echo "Creating Build..."
npm run build

echo "Running FE as a service..."
pm2 restart v3-frontend
pm2 save

echo "Cloud Traffic V2 Deployment is Done."
exit 1
