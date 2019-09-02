# go-fork-cmd

**fork** utility will run a given number of instances of a script or command while ensuring the number of executions does not exceed an informed limit.

This can be specialy useful, when one requires to run the same command or script multiple times but still want to control the volume of executions, like for instance to simulate concurent access to an API using `cURL`.

The same result can be achieved in bash with a `while` loop counting current executions, however `fork` benefits from the small memory footprint of goroutines (around 230KB for each thread) and better safety/control provided by channels for efficiently managing the execution queue.

## Download

Download the binaries from the [latest release](https://github.com/fsilveir/go-fork-cmd/releases/latest) to your local machine and place it in your PATH and follow the instructions below for usage.

## Usage

```
$ fork -help
This utility will concurrently run a given number of instances of a script or command,
while ensuring the number of executions does not exceed an informed limit.

Usage:
fork -c [ <script_path> | <command> ] -t <total_executions> -l <limit>

Examples:
$ fork -c "curl google.com -n 50 -l 10
$ fork -c "/tmp/test.sh" -n 50 -l 10

This will execute the informed script/command 50 times, limiting concurrency to 10 at the time.

Options:
   -c        Path for the source script or command to be executed.
   -t        Total number of concurrent executions (default is 2).
   -l        Limit of instances to be concurrently executed (default is 10).

```

## Building from Source

To build the binaries from source, execute the following (if you already have GOPATH configured on your local machine):

```bash
go get github.com/fsilveir/go-fork-cmd
```
Or clone the repository directly with the following command:

```bash
git clone git@github.com:fsilveir/go-fork-cmd.git
```

After succesfully downloading the files from the repository, execute the script `build.sh`, as shown below`:

```bash
~/go/src/github.com/fsilveir/go-fork-cmd $ ./build.sh
Revision is 1172888
Building GOOS=windows GOARCH=amd64...
Building GOOS=windows GOARCH=386...
Building GOOS=linux GOARCH=amd64...
Building GOOS=linux GOARCH=386...
Building GOOS=freebsd GOARCH=amd64...
Building GOOS=freebsd GOARCH=386...
Building GOOS=darwin GOARCH=amd64...
Building GOOS=darwin GOARCH=386...
fork_darwin_386: OK
fork_darwin_amd64: OK
fork_freebsd_386: OK
fork_freebsd_amd64: OK
fork_linux_386: OK
fork_linux_amd64: OK
fork_windows_386.exe: OK
fork_windows_amd64.exe: OK
```

## Get support

Create an [issue](https://github.com/fsilveir/go-fork-cmd/issues) if you want to report a problem or ask for a new functionality any feedback is highly appreciated!
