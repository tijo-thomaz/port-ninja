package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: kill_port <port1> <port2> ... <portN>")
		os.Exit(1)
	}

	help := flag.Bool("help", false, "Display this help message")
	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		flag.Usage()
	}

	ports := []int{}
	for _, arg := range flag.Args() {
		port, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Invalid port number: %s\n", arg)
			os.Exit(1)
		}
		ports = append(ports, port)
	}

	for _, port := range ports {
		pids, err := getPIDsByPort(port)
		if err != nil {
			fmt.Printf("Error getting PIDs for port %d: %v\n", port, err)
			continue
		}

		if len(pids) > 0 {
			for _, pid := range pids {
				err = killProcess(pid)
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

func getPIDsByPort(port int) ([]int, error) {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf(":%d", port))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out.String(), "\n")
	var pids []int
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) > 1 {
			pid, err := strconv.Atoi(fields[1])
			if err == nil {
				pids = append(pids, pid)
			}
		}
	}

	return pids, nil
}

func killProcess(pid int) error {
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	return cmd.Run()
}
