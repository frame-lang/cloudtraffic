# Cloud Traffic V1

This repository contains the source code to demonstrate the Traffic Light wrokflow. In this initial version, the front-end and back-end communicates with help of web sockets and the traffic light data will get stored in disk storage.
The data will get stored in a `data/` directory on V1 root level. Data of each user will be stored on seaparate file. The name of the file in `data/` directory is the connectionID of a user. 

---

## Tech Stack Used

- React
- GO
- WebSockets
- Frame Lang

---

## Development/Local Setup


### Pre-requisites

- [Node](vultr.com/docs/install-nvm-and-node-js-on-ubuntu-20-04/)
- [GO](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04)

### Clone the repository

```
git clone https://github.com/frame-lang/cloudtraffic.git
```

### Front-end

- Go to the front-end folder (`cloudtraffic/cloudtraffic_v1/frontend-react`) and install dependencies
```
npm install
```

- Start the server
```
npm start
```

### Back-end

- Go to the back-end folder (`cloudtraffic/cloudtraffic_v1/backend-go`) and install dependencies.
```
go install
```

- Start the server
```
go run main.go
```

---

## Deployment

1. With Bash Script (Both Front-end and Back-end at once)

```
curl -sSf https://raw.githubusercontent.com/frame-lang/cloudtraffic/main/cloudtraffic_v1/cloudtraffic-v1-deploy.sh | sh
```

2. Manual deployment

    a.  Front-end

    - Move to V1 front-end folder
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
    pm2 restart v1
    ```
    - Save pm2 progress
    ```
    pm2 save
    ```   

    b. Back-end

    - Move to V1 back-end folder
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
    sudo service trafficlightbackendv1 restart
    ```