#!/usr/bin/sudo sh

PROJECT_DIR="/root/cloudtraffic"
BACKEND_DIR="/root/cloudtraffic/cloudtraffic_v1/backend-go"
FRONTEND_DIR="/root/cloudtraffic/cloudtraffic_v1/frontend-react"

cd $PROJECT_DIR
git pull origin main

cd $BACKEND_DIR

echo "Deleting previous backend build...." 
rm -rf cloudtraffic_v1
echo "Creating BE Build..."
go build

SERVICEFILE="/lib/systemd/system/trafficlightbackendv1.service"

if [ -f "$SERVICEFILE" ]; then
echo "Restarting backend service..."
    sudo systemctl daemon-reload
    sudo service trafficlightbackendv1 restart
else
    echo "Creating backend service..."
    sudo touch $SERVICEFILE
    sudo echo "
    [Unit]
    Description=Traffic Light Backend V1
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    WorkingDirectory=$BACKEND_DIR
    ExecStart=$BACKEND_DIR/cloudtraffic_v1
    [Install]
    WantedBy=multi-user.target" | sudo tee -a $SERVICEFILE > /dev/null
    sudo systemctl daemon-reload
    sudo service trafficlightbackend restart
    sudo systemctl enable trafficlightbackend.service
fi

cd $FRONTEND_DIR
echo "Deleting previous frontend build...."
sudo rm -rf build
echo "Installing Dependencies..."
sudo npm install
echo "Creating Build..."
sudo npm run build

echo "Running FE as a service..."
pm2 restart v1
pm2 save

echo "Cloud Traffic V2 Deployment is Done."
exit 1
