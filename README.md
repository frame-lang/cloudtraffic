# cloudtraffic


## Perquisites
1. Clone the project   
```git clone https://github.com/frame-lang/cloudtraffic.git```
2. [Golang](https://go.dev/doc/install)
3. [NodeJs](https://nodejs.org/en/download/)
4. [Nginx](https://www.nginx.com/resources/wiki/start/topics/tutorials/install/)

## Deployment

1. Create Compute Engine on GCP   
1.1 Select perferred machine configuration.   
1.2 Create a firewall rule for backend service PORT(Currently, it is 8000). [Docs](https://cloud.google.com/vpc/docs/firewalls)
2. change directory to backend-go   
```cd backend-go```
3. compiled backend  
```go install && go build```
4. Create a unix service  
4.1 create service config  
```sudo touch /lib/systemd/system/trafficlightbackend.service```
4.2 Paste below code in service config  
    ```
    # ExecStart has path of binary which generated by go build
    [Unit]
    Description=Traffic Light Backend go
    [Service]
    Type=simple
    Restart=always
    RestartSec=5s
    ExecStart=/home/gauravsaxena_ongraph/backend-go/persistenttrafficlight
    [Install]
    WantedBy=multi-user.target
    ```
    4.3 Start the service  
    ```sudo service trafficlightbackend start```
5. Build frontend  
5.1 Install dependencies && make build   
```sudo npm install && sudo npm run build```
5.2 Point nginx to frontend build  
     ```
     # change root /var/www/html to root /home/gauravsaxena_ongraph/frontend-react/build
     sudo nano /etc/nginx/sites-available/default
     ```
    5.3 Restart Nginx service  
    ```sudo service nginx restart```



