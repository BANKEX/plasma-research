package main

import (
	"context"
	"fmt"

	"github.com/BANKEX/plasma-research/src/node/cmd"
	"github.com/BANKEX/plasma-research/src/node/operator"
	"golang.org/x/sync/errgroup"
)

func main() {
	cmd.NewCmd(run).Execute()
}

func run(app cmd.AppContext) error {
	cfg, err := operator.NewConfig()
	if err != nil {
		return fmt.Errorf("failed parse configs: %s", err)
	}

	o, err := operator.NewOperator(cfg)
	if err != nil {
		return fmt.Errorf("failed to build operator instance: %s", err)
	}

	ctx := context.Background()

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		return cmd.WaitInterrupted(ctx)
	})
	wg.Go(func() error {
		return o.Serve(ctx)
	})

	if err := wg.Wait(); err != nil {
		return fmt.Errorf("termination: %s", err)
	}

	return nil
}
