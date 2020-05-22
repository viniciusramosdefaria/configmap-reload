package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var (
	volumeDirs  volumeDirsFlag
	processPath processFlag
	wg          sync.WaitGroup
	pid         string
)

func main() {
	flag.Var(&volumeDirs, "volume-dir", "the config map volume directory to watch for updates; may be used multiple times")
	flag.Var(&processPath, "process-path", "the process path to send SIGHUP to when the specified config map volume directory has been updated")
	flag.Parse()

	if err := argValidator(); err != nil {
		log.Fatal(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	wg.Add(1)

	go eventHandler(watcher)

	directoryWatcher(watcher)

	wg.Wait()
}

type volumeDirsFlag []string

type processFlag []string

func (v *volumeDirsFlag) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v *volumeDirsFlag) String() string {
	return fmt.Sprint(*v)
}

func (p *processFlag) Set(value string) error {
	*p = append(*p, value)
	return nil
}

func (p *processFlag) String() string {
	return fmt.Sprint(*p)
}

func argValidator() error {
	if len(volumeDirs) < 1 {
		flag.Usage()
		return errors.New("Missing volume-dir")
	}

	if len(processPath) < 1 {
		flag.Usage()
		return errors.New("Missing process-path")
	}
	return nil
}

func eventHandler(watcher *fsnotify.Watcher){
	for {
		select {
		case event := <-watcher.Events:
			switch {
			case event.Op&fsnotify.Write == fsnotify.Write:
				for _, p := range processPath {
					cmd := exec.Command("/usr/bin/pidof", p)
					var out bytes.Buffer
					var errExec bytes.Buffer
					cmd.Stdout = &out
					cmd.Stderr = &errExec
					cmd.Run()
					osReturn := out.String()
					if len(osReturn) > 0 {
						pid, _ := strconv.Atoi(strings.TrimSuffix(osReturn, "\n"))
						process, _ := os.FindProcess(pid)
						process.Signal(syscall.SIGHUP)
						log.Printf("successfully triggered reload for %v, pid: %v", p, out.String())
					}
				}
			}

		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
	defer wg.Done()
}

func directoryWatcher(watcher *fsnotify.Watcher){
	for _, d := range volumeDirs {
		log.Printf("Watching directory: %q", d)
		if err := watcher.Add(d); err != nil {
			log.Fatal(err)
		}
	}
}