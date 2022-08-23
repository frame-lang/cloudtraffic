# Cloud Traffic V3

This repository contains the source code to demonstrate the Traffic Light wrokflow. In this version, we have four different components.

1. Front-end server: Hosted on GCP VM (Compute Engine),  Written in React.
2. Back-end server (Utils service):  Hosted on GCP VM (Compute Engine),  Written in GO. It is used to communicate between Front-end and TL service.
3. TL service: Serverless on Google Cloud Function,
 Written in Go. It contains all the logic for State machine and manager.
4. Redis DB: Connected via GCP Memorystore. Used to store Traffic light data.

Flow:
> UI(React app) <-- WebSockets --> Util Service (GO app) <-- RabbitMQ --> TL service (Cloud function) <-- Serverless VPC Connector --> Redis (Memorystore)

## Tech Stack Used

- React
- GO
- Frame Lang
- WebSockets
- GCP VM (Compute Engine)
- Google Cloud Function
- Memorystore
- Serverless VPC Connector
- Docker
- RabbitMQ

## Development/Local Setup

### Pre-requisites

- [Node](vultr.com/docs/install-nvm-and-node-js-on-ubuntu-20-04/)
- [GO](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04)

### Clone the repository

```
git clone https://github.com/frame-lang/cloudtraffic.git
```

### Front-end

- Go to the front-end folder (`cloudtraffic/cloudtraffic_v2/frontend-react`) and install dependencies
```
npm install
```

- Start the server
```
npm start
```

### Back-end

- Go to the back-end folder (`cloudtraffic/cloudtraffic_v2/backend-go`) and install dependencies.
```
go install
```

- Start the server
```
go run main.go
```

## Deployment

1. With Bash Script (Both Front-end and Back-end at once)

    - Switch to Root user
    ```
    sudo su
    ``` 
    - Run the bash script
    ```
    curl -sSf https://raw.githubusercontent.com/frame-lang/cloudtraffic/main/cloudtraffic_v3/cloudtraffic-v3-deploy.sh | sh
    ```

2. Manual deployment

    Switch to Root user
    ```
    sudo su
    ```

    a.  Front-end

    - Move to V2 front-end folder
    - Install dependencies
    ```
    npm install
    ```
    - Create Build
    ```
    npm run build
    ```
    - Reload pm2 instance
    ```
    pm2 restart v3-frontend
    ```
    - Save pm2 progress
    ```
    pm2 save
    ```   

    b. Back-end

    - Move to V2 back-end folder
    - Install dependencies
    ```
    go install
    ```
    - Build Go app
    ```
    go build
    ```
    - Reload the **systemd** service
    ```
    sudo systemctl daemon-reload
    sudo service trafficlightbackendv3 restart
    ```