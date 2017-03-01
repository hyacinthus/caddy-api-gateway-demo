# Caddy API gateway demo
Use caddy as microservice API gateway in docker swarm.

## RUN
```bash
docker stack deploy -c docker-compose.yml caddy
```

## TRY
visit http://{your-host-ip}:2015  
```
Hello, world! I am 72ec231a5eba :)

Hello, world! I am c1a1bc3e79a1 :)

Hello, world! I am 3e4f9f63c1a0 :)
```
app will tell you which container it is running in.
