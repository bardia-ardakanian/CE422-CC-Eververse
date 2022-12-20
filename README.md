# CE422-CC-Eververse
Cloud Computing course project

## Usage
```bash
docker-compose up --build -d --force-recreate
```
Browse to `http://localhost` and you should see the landing page.

## Kubernetes
### Deployment
```bash
chmod +x deploy.sh
./deploy.sh
‍‍‍```

### Testing method 1
```bash
chmod +x demo.sh
./demo.sh
```
### Testing method 2
```bash
kubectl run webkit --image=bardiaardakanian/webkit -i --tty -- sh
# curl -v -Ss http://eververse-service.default.svc.cluster.local:8080
```

## Requests

Browse to `localhost/get?name=<coin_name>` to get exchange rate of the given digital currency. for example for getting the exchange rate of bitcoin `localhost/get?name=BTC` can be used.
Digital currency exchange rates will be cached in the redis cache for 5 minutes by default.

## Configure

Following fields are configurable.
- COIN-API-KEY
- Web MAPPING Port (80 by default)
- Web EXPOSE Port (1323 by default)
- Redis MAPPING Port (6379 by default)
- Redis EXPOSE Port (6379 by default)
- Redis cache expiration by minutes (5 by default)
