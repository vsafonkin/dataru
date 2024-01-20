package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/vsafonkin/dataru/config"
)

var cpuprofile = flag.String("cpuprofile", "", "dataru -cpuprofile <filename>")
var configPath = flag.String("config", "", "dataru -config <path>")

func main() {
	defer pprof.StopCPUProfile()
	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	if err := config.LoadConfig(*configPath); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	host := config.Host()
	fmt.Println(host)

	port := config.Port()
	fmt.Println(port)

	logger := config.Logger()
	fmt.Println(logger)

	names := config.Names()
	for _, v := range names {
		fmt.Println(v)
	}
}

func parseFlags() error {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(fmt.Sprintf("./%s/%s", "pprof", *cpuprofile))
		if err != nil {
			return fmt.Errorf("create cpu profile error: %w", err)
		}
		pprof.StartCPUProfile(f)
	}
	if *configPath == "" {
		return fmt.Errorf("need to set config path by -config option")
	}
	return nil
}
