package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BANKEX/plasma-research/src/node/cmd"
	"github.com/BANKEX/plasma-research/src/node/verifier"
	"golang.org/x/sync/errgroup"
)

func main() {
	cmd.NewCmd(run).Execute()
}

func run(app cmd.AppContext) error {
	cfg, err := verifier.NewConfig()
	if err != nil {
		log.Printf("failed parse configs: %s", err)
		os.Exit(1)
	}

	v, err := verifier.NewVerifier(cfg)
	if err != nil {
		log.Printf("failed to build verifier instance: %s", err)
		os.Exit(1)
	}

	ctx := context.Background()

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		return cmd.WaitInterrupted(ctx)
	})
	wg.Go(func() error {
		return v.Serve(ctx)
	})

	if err := wg.Wait(); err != nil {
		return fmt.Errorf("termination: %s", err)
	}

	return nil
}
