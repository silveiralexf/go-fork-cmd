// Package go-fork-cmd will concurrently run a given number of instances of a script or command
// while ensuring the number of executions does not exceed an informed limit.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const (
	errMsg = (`ERROR: You did not specify a valid command or failed to pass the proper options. Exiting!

Use "-help" or "-h" for usage instructions.
`)
	helpMsg = (`This utility will concurrently run a given number of instances of a script or command,
while ensuring the number of executions does not exceed an informed limit.

Usage:
fork -c [ <script_path> | <command> ] -t <total_executions> -l <limit>

Examples:
$ fork -c "curl google.com -n 50 -l 10
$ fork -c "/tmp/test.sh" -n 50 -l 10

This will execute the informed script/command 50 times, limiting concurrency to 10 at the time.

Options:`)
)

var (
	wg      sync.WaitGroup
	command = flag.String("c", "", "Path for the source script  or command to be executed.")
	total   = flag.Int("t", 2, "Total number of concurrent executions (default is 2).")
	max     = flag.Int("l", 10, "Limit of instances to be concurrently executed (default is 10).")
)

func main() {

	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Println(helpMsg)
		order := []string{"c", "t", "l"}
		for _, name := range order {
			flag := flagSet.Lookup(name)
			fmt.Printf("   -%s   ", flag.Name)
			fmt.Printf("     %s\n", flag.Usage)
		}
	}
	flag.Parse()

	// Creates a channel with a size of the total number of executions to be made
	ch := make(chan int, *total)

	// max number of goroutines that will concurrently run
	wg.Add(*max)

	for i := 0; i < *max; i++ {
		go func() {
			for {
				num, ok := <-ch
				if !ok { // if there is nothing to do and the channel has been closed then end the goroutine
					wg.Done()
					return
				}

				// Execute a script or command depending on the input of the user
				if *command != "" {
					cmdExec(command, total, max, num)

				} else {
					fmt.Println(errMsg)
					os.Exit(2)
				}
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	for i := 0; i < *total; i++ {
		ch <- i // add i to the queue
	}

	close(ch) // This tells the goroutines there's nothing else to do
	wg.Wait() // Wait for the threads to finish
}

func cmdExec(command *string, total, max *int, num int) {

	args := strings.Split(*command, " ")
	cmd := exec.Command(args[0], args[1:]...)

	var stdout, stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("ERROR: Failed with the following error:\n\n%s\n\n", err)
		os.Exit(1)
	}
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out %d:\n%s\nerr:\n%s\n", num, outStr, errStr)
}
