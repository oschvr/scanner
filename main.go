package main

import (
	"fmt"
	"net"
	"os"
	"sort"

	"github.com/mbndr/figlet4go"
)

func worker(host string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	// Figlet ASCII banner
	ascii := figlet4go.NewAsciiRender()
	renderOpts := figlet4go.NewRenderOptions()
	renderOpts.FontColor = []figlet4go.Color{
		figlet4go.ColorCyan,
	}
	renderStr, _ := ascii.RenderOpts("scanner", renderOpts)
	fmt.Print(renderStr)
	fmt.Println("----------------")

	// Vars
	var host string
	if len(os.Args) > 1 {
		host = os.Args[1]
	} else {
		fmt.Println("Error 🔴: No host given !")
		os.Exit(1)
	}

	portCount := 655365

	// Helpers
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	// Scanner logic
	// ------------
	// Send ports and results channels to worker func
	// Worker func will call net.Dial to verify open port
	fmt.Println("Warning 🌕: Scanning... !")
	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, results)
	}

	go func() {
		for i := 0; i < portCount; i++ {
			ports <- i
		}
	}()

	// Append open ports to list
	// Iterate list and show results
	for i := 1; i < portCount; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Connection to %s succesful on port %d", host, port)
		fmt.Println("")
	}

	if len(openports) >= 1 {
		fmt.Println("Success ✅: Scan complete")
	}
}
