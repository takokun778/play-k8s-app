# alpha

```bash
curl --header "Content-Type: application/json" --data '{"message": "hello"}' http://localhost:8080/alpha.v1.AlphaService/Info
```

```bash
grpcurl -protoset <(buf build -o -) -plaintext -d '{"message": "hello"}' localhost:8080 alpha.v1.AlphaService/Info
```
