# #!/bin/bash

sudo apt update && sudo apt install git wget nginx -y

Go_latest="1.18"

Go_current=$(go version | egrep -o "([0-9]{1,}\.)+[0-9]{1,}")

if [ $Go_current == $Go_latest ]
then
source $HOME/.profile
else
echo "## NEW VERSION FOUND GO-$Go_latest ###"
sudo rm -rf /usr/local/go go$Go_latest.linux-amd64.tar.gz
sudo wget https://go.dev/dl/go$Go_latest.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go$Go_latest.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.profile
source $HOME/.profile
echo "## GO-$Go_latest Installed ###"
fi


DIR="$HOME/cloudtraffic"
cd $HOME
if [ -d "$DIR" ]; then
    cd $DIR
    git pull https://github.com/frame-lang/cloudtraffic.git

else

    git clone https://github.com/frame-lang/cloudtraffic.git

fi

cd $DIR/backend-go

echo "#### Delete previous backend build.... ####" 
rm -rf persistenttrafficlight
echo "#### Installing Dependencies... ####"
go install
echo "#### Creating Build... ####"
go build

SERVICEFILE="/lib/systemd/system/trafficlightbackend.service"

if [ -f "$SERVICEFILE" ]; then
echo "#### Restarting backend service... ####"
    sudo service trafficlightbackend restart
else
    echo "#### Creating backend service... ####"
    sudo touch $SERVICEFILE
    sudo echo "
    [Unit]
    Description=Traffic Light Backend go
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    ExecStart=$DIR/backend-go/persistenttrafficlight
    [Install]
    WantedBy=multi-user.target" | sudo tee -a $SERVICEFILE > /dev/null
    sudo service trafficlightbackend restart
    sudo systemctl enable trafficlightbackend.service
fi

cd $DIR/frontend-react
echo "#### Delete previous frontend build.... ####"
sudo rm -rf build
echo "#### Installing Dependencies... ####"
sudo npm install
echo "#### Creating Build... ####"
sudo npm run build

sudo chmod -R 777 $DIR/frontend-react/build

sudo sed -i "s|/var/www/html|$DIR/frontend-react/build|g" /etc/nginx/sites-available/default

echo "#### Restarting Nginx service... ####"

sudo service nginx restart

echo "Deployment is Done"
