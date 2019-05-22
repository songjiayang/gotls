# Go TLS Example

### How to run

```
go run main.go
```

The output is: 

```bash
16:15 $ go run main.go 
2019/05/22 16:15:30 receive data from server: pong
2019/05/22 16:15:30 receive data from client: ping
2019/05/22 16:15:31 receive data from client: ping
2019/05/22 16:15:31 receive data from server: pong
2019/05/22 16:15:32 receive data from client: ping
2019/05/22 16:15:32 receive data from server: pong
2019/05/22 16:15:33 receive data from server: pong
2019/05/22 16:15:33 receive data from client: ping
2019/05/22 16:15:34 receive data from server: pong
2019/05/22 16:15:34 receive data from client: ping
```

### RSA Generate

- command example:
 
```bash
$ openssl genrsa -out server.key 2048
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

More command please check link: https://bbengfort.github.io/programmer/2017/03/03/secure-grpc.html .