package main

import (
	goflag "flag"
	"time"

	"github.com/dongyiyang/simplegoserver/server"
	"github.com/spf13/pflag"
)

// TODO(thockin): This is temporary until we agree on log dirs and put those into each cmd.
func init() {
	goflag.Set("logtostderr", "true")
}

func main() {
	InitFlag()
	go generateCPULoad()
	server.ListenAndServe()
}

func InitFlag() {
	// pflag.CommandLine.SetNormalizeFunc(WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	pflag.Parse()
}

func generateCPULoad() {
	for {
		time.Sleep(10 * time.Millisecond)
	}
}
