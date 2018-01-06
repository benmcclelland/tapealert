package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/benmcclelland/tapealert"
)

func main() {
	var severity int
	flag.IntVar(&severity, "severity", 0,
		"Define min severity level [0=all, 1=warning+critical, 2=critical only]")
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage(fmt.Errorf("no argument specified"))
	}

	ta, err := tapealert.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Tapealert init:", err)
		os.Exit(1)
	}

	i, err := strconv.ParseInt(flag.Args()[0], 0, 64)
	if err != nil {
		usage(err)
	}

	alerts := ta.FromFlags(i, severity)
	for i, alert := range alerts {
		fmt.Println(i + 1)
		fmt.Println(alert.Name)
		fmt.Println("-")
		fmt.Println("Severity:", alert.Severity)
		fmt.Println("-")
		fmt.Println(alert.Message)
		fmt.Println("-")
		fmt.Println("Cause:", alert.Cause)
		fmt.Println("--------")
	}
}

func usage(err error) {
	fmt.Fprintf(os.Stderr, "%s\n%s\n%s\n%s\n",
		"See -h for optional flags",
		"Arg must be hex, octal, or decimal.",
		"Hex must begin with 0x, and Octal must begin with 0", err)
	os.Exit(1)
}
