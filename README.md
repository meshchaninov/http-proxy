# http-proxy

[In dockerhub](https://hub.docker.com/repository/docker/meshchaninov/http-proxy/)

Simple http server using [goproxy](https://github.com/elazarl/goproxy) with or without auth.

## Start container with proxy
```shell
docker run -d --name http-proxy -p <PROXY_PORT>:80 -e PROXY_USER=<PROXY_USER> -e PROXY_PASS=<PROXY_PASSWORD> meshchaninov/http-proxy
```

- For auth-less mode just do not pass **PROXY_USER** and **PROXY_PASS**.  
- Change **PROXY_PORT** for custom proxy port

## Test running service
### with auth
```bash
curl --proxy http://<PROXY_USER>:<PROXY_PASS>@<DOCKER_MACHINE_IP>:<PROXY_PORT> https://ya.ru 
```
### without auth
```bash
curl --proxy http://<DOCKER_MACHINE_IP>:<PROXY_PORT> https://ya.ru 
```
result must show docker host ip (for bridged network)
