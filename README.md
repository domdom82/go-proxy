# go-proxy
Simple HTTP Proxy


# Deploying to CloudFoundry

Make sure you pass the target via

```
cf push go-proxy -c "go-proxy -t target https://my-target.com"
```
