# Caddy API gateway demo
Use caddy as microservice API gateway in docker swarm.

## Requirements
* Docker Engine 1.13+
* docker-compose 1.11.2+

## Notice
The image in docker-compose is build by Dockerfile.app and Dockerfile.gateway in this demo.
I've host it in docker hub, so you can use it in next step directly.
You can also build the image by yourself, then don't forget to change the image name in docker-compose.yml

## Run Server
Pull images from docker hub(~30M):
```bash
docker-compose pull
```
Run
```bash
docker stack deploy -c docker-compose.yml caddy
```

## Try
Because of caddy's proxy connection reuse feature, you can't test this demo by 
visit http://{your-host-ip}:2015  in your browser.
I write a client to send concurrent requests.  
If you have docker, 
```bash
docker run --rm -it muninn/caddy-microservice:client
```
You can also use unix shell version:
```bash
cd client
sh test.sh
```
app will tell you which container it is running in.
