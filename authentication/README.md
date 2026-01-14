
## OTP

## Description

This is a simple OTP service that uses RabbitMQ to send OTPs to users.

![otp](./assets/otp.png)

change the .env.example file to .env and fill the values

run 

```bash

docker compose up -d
```

create a acount in http//localhost:8080/auth/register

send a forgot password request in http//localhost:8080/auth/forgot

check the opt code in email in http://localhost:8025

verify the opt code in http//localhost:8081/otp/validation



This is a simple authentication server that uses JWT to authenticate users.

## Auth

![Auth Server](./assets/auth.png)

openssl rand -hex 32

