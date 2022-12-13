# Mac Address Look Up CLI Tool
simple client for MAC address lookup

# How to run
1. build the docker image via `docker build -t <name> .`
2. run the container via `docker run <name>`
3. run the container with the `-h` flag to list out the supported cli flags
```
docker run brightsign -h
Usage of /brightsign:
  -addr string
    	the mac address to look up
  -key string
    	the secret api key
  -v	print out all data from the response
```
4. you'll need a macaddress.io access token and provide it to the cli via the -key flag
5. provide a mac address to lookup via the -addr flag
```
docker run brightsign --key <access-token> --addr 44:38:39:ff:ef:57  | jq .
```