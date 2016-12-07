package main

import (
	"flag"
	"fmt"
)

/* Command line parameters */
type ParamsStruct struct {
	rootPath string
	intPort  int
	debug    bool
}

var params ParamsStruct

func init() {
	sep := "\n    \t"
	flag.StringVar(&params.rootPath, "path", "", "The directory to serve."+sep+"(default: current directory)")
	flag.BoolVar(&params.debug, "debug", false, "Show debug output.")
	flag.IntVar(&params.intPort, "port", 5555, "Start local webserver that serves files at -path.")
}

func setupFlags(f *flag.FlagSet) {
	f.Usage = func() {
		f.VisitAll(func(f *flag.Flag) {
			s := fmt.Sprintf("  -%s", f.Name) // Two spaces before -; see next two comments.
			name, usage := flag.UnquoteUsage(f)
			if len(name) > 0 {
				s += " " + name
			}
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"

			s += usage
			if !isZeroValue(f.DefValue) {
				s += fmt.Sprintf("\n    \t(default %v)", f.DefValue)
			}
			fmt.Println(s)
		})
	}
}

// isZeroValue guesses whether the string represents the zero
// value for a flag. It is not accurate but in practice works OK.
func isZeroValue(value string) bool {
	switch value {
	case "false":
		return true
	case "":
		return true
	case "0":
		return true
	}
	return false
}
