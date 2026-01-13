# OTP

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

