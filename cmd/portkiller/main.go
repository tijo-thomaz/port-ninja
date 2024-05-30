package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/tijo-thomaz/port-ninja/internal/ports"
	"github.com/tijo-thomaz/port-ninja/internal/process"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: port-ninja <port1> <port2> ... <portN>")
		os.Exit(1)
	}

	help := flag.Bool("help", false, "Display this help message")
	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		flag.Usage()
	}

	portsToKill := []int{}
	for _, arg := range flag.Args() {
		port, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Invalid port number: %s\n", arg)
			os.Exit(1)
		}
		portsToKill = append(portsToKill, port)
	}

	for _, port := range portsToKill {
		pids, err := ports.GetPIDsByPort(port)
		if err != nil {
			fmt.Printf("Error getting PIDs for port %d: %v\n", port, err)
			continue
		}

		if len(pids) > 0 {
			for _, pid := range pids {
				err = process.KillProcess(pid)
				if err != nil {
					fmt.Printf("Error killing process with PID %d: %v\n", pid, err)
				} else {
					fmt.Printf("Successfully killed process with PID %d on port %d\n", pid, port)
				}
			}
		} else {
			fmt.Printf("No process found on port %d\n", port)
		}
	}
}
