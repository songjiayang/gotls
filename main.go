package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	cer, err := tls.LoadX509KeyPair("./server.crt", "./server.key")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	l, err := tls.Listen("tcp", "localhost:9090", config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go sendMsg()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			reader := bufio.NewReader(c)

			go func() {
				for {
					data, _, err := reader.ReadLine()
					if err != nil {
						log.Printf("reader.ReadLine(): %v", err)
						return
					}

					log.Printf("receive data from client: %s", string(data))
				}
			}()

			for {
				c.Write([]byte("pong\n"))
				time.Sleep(time.Second)
			}

		}(conn)
	}
}

func sendMsg() {
	rootCert, err := ioutil.ReadFile("./server.crt")
	if err != nil {
		log.Fatal(err)
	}

	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(rootCert); !ok {
		log.Fatal("failed to parse root certificate")
	}

	config := &tls.Config{RootCAs: roots}
	conn, err := tls.Dial("tcp", "localhost:9090", config)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	go func() {
		for {
			data, _, err := reader.ReadLine()
			if err != nil {
				log.Printf("reader.ReadLine1(): %v", err)
				return
			}

			log.Printf("receive data from server: %s", string(data))
		}
	}()

	for {
		conn.Write([]byte("ping\n"))
		time.Sleep(time.Second)
	}
}
