# SERVERS

## Overview

Each folder contains a different architecture pattern for servers.

## OTP

### Overview

This architecture pattern is used for sending one time password (OTP) to users.

![OTP](./authentication/assets/otp.png)

> Go - Gin - GORM - RabbitMQ - Redis - SMTP - Mailhog - Docker - Docker Compose - Terraform - DDD

## BFF (Backend For Frontend)

### Overview

This architecture pattern is used for aggregating and formatting data from different services for the frontend. One of the main benefits of this pattern is that it can reduce the number of calls to the backend and improve the performance of the frontend.

### Client call for BFF
![Dashboard](./BFF/assets/dashboard.png)

### BFF call for services
![BFF Communication](./BFF/assets/bff_comunication.png)

> Go - Gin - Docker - Docker Compose - gRPC - restApi - DDD - Concurrency

## Authentication

### Overview

This architecture pattern is used for authentication and authorization of users that uses JWT. That way, the frontend can store the token in the browser keeping the user logged in until the token expires.

![Authentication](./authentication/assets/auth.png)

> Go - Gin - GORM - Docker - Docker Compose - Postgres - JWT - DDD
