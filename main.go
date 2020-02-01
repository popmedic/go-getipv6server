package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<html>", "<body>", "<h1>IP Echo</h1>")
		defer fmt.Fprint(w, "</body>", "</html>")

		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if nil != err {
			fmt.Fprintf(w, "error (%q): %q is not in ip:port format", err, r.RemoteAddr)
		}

		uIP := net.ParseIP(ip)
		if nil == uIP {
			fmt.Fprintf(w, "error: %q is not in ip:port format", ip)
			return
		}

		f := r.Header.Get("X-Forwarded-For")

		fmt.Fprint(w, "<p>", uIP, "</p>")
		fmt.Fprint(w, "<p>Forwarded:<br/>", strings.Join(strings.Split(f, ", "), "<br/>\n"))
	})

	log.Print("server starting")
	log.Fatal(http.ListenAndServe(":8099", nil))
}
