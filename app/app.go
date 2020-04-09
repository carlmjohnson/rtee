package app

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/carlmjohnson/flagext"
	"github.com/peterbourgon/ff"
)

const AppName = "rtee"

func CLI(args []string) error {
	var app appEnv
	err := app.ParseArgs(args)
	if err == nil {
		err = app.Exec()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	return err
}

func (app *appEnv) ParseArgs(args []string) error {
	fl := flag.NewFlagSet(AppName, flag.ContinueOnError)
	dst := flagext.FileWriter(flagext.StdIO)
	app.dst = dst
	fl.Var(dst, "dst", "secondary output file or URL")
	fl.Usage = func() {
		fmt.Fprintf(fl.Output(),
			`rtee - Like tee but with automatic process substitution.`,
		)
		fl.PrintDefaults()
		fmt.Fprintln(fl.Output(), "")
	}
	if err := ff.Parse(fl, args, ff.WithEnvVarPrefix("RTEE")); err != nil {
		return err
	}
	app.args = fl.Args()

	return flagext.MustHaveArgs(fl, 1, -1)
}

type appEnv struct {
	dst  io.Writer
	args []string
}

func (app *appEnv) Exec() (err error) {
	stdin := io.TeeReader(os.Stdin, app.dst)
	cmd := exec.Command(app.args[0], app.args[1:]...)
	cmd.Stdin = stdin
	cmd.Stdout = os.Stderr
	return cmd.Run()
}
