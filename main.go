package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: whoiscli <domain>")
		return
	}

	domain := os.Args[1]
	server := "whois.verisign-grs.com:43" // .com/.net ドメイン向け

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error connecting to whois server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "%s\r\n", domain)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
