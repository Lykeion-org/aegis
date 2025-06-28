# Aegis
Aegis is the service that handles all security related processes and creates and validates JWT-tokens. Aegis is approachable via GRPC and REST, although REST is mainly used for debugging and administrative purposes, since Aegis should not be exposed to the end user directly.

## TODO:
- Move config variables to config file
- Implement ping request from client, so orchestrator can get information about the service

## Config
Currently the variables that can be configured are the following:
- REST_PORT
- GRPC_PORT
_ JWT_SECRET

