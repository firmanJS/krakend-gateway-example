# API Gateway

Repository sebagai api gateway untuk mendistribusikan service api
## Start the service

* Via Make :

```bash
# Copy enviroment variables from .env.example to .env
cp .env.example .env
cp krakend/config/settings/env.example.json krakend/config/settings/env.json

# Copy Makefile
cp Makefile.example Makefile

# Build application
make docker-build

# Stop aplication
CTRL+C 
# then 
make docker-down

# After build you can run command with this
make docker-start
```