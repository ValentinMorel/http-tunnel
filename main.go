package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"

	"github.com/ValentinMorel/http-tunnel/handlers"
	"github.com/gliderlabs/ssh"
)

func main() {

	go func() {
		http.HandleFunc("/list", handlers.List)
		http.HandleFunc("/download", handlers.Download)
		log.Fatal(http.ListenAndServe(":3000", nil))
	}()

	ssh.Handle(func(s ssh.Session) {
		id := rand.Intn(math.MaxInt)
		handlers.Tunnels[id] = make(chan handlers.Tunnel)

		s.Write([]byte(fmt.Sprintf("tunnel ID -> %d\n", id))) // Send the tunnel ID to the SSH client
		fmt.Println("tunnel ID -> ", id)
		tunnel := <-handlers.Tunnels[id]
		fmt.Println("tunnel is ready")

		_, err := io.Copy(tunnel.W, s)
		if err != nil {
			log.Fatal(err)
		}
		close(tunnel.Done)
		s.Write([]byte("tunnel successfully established and closed\n"))
	})

	log.Fatal(ssh.ListenAndServe(":2222", nil))
}
