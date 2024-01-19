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
var cfgPath = flag.String("config", "", "dataru -config <path>")

func main() {
	defer pprof.StopCPUProfile()
	if err := parseFlags(); err != nil {
		log.Fatal(err)
	}

	if err := config.LoadConfig(*cfgPath); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	host, err := config.Host()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host)

	port, err := config.Port()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(port)

	logger, err := config.Logger()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(logger)
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
	if *cfgPath == "" {
		return fmt.Errorf("need to set config path by -config option")
	}
	return nil
}
