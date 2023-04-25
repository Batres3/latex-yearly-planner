package app

import (
	"context"
	"fmt"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/commanders"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/filewriters"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/mos/mostitles"
	"github.com/kudrykv/latex-yearly-planner/internal/adapters/texindexer"
	"github.com/kudrykv/latex-yearly-planner/internal/core/plannerbuilders"
	"github.com/kudrykv/latex-yearly-planner/internal/core/planners"
	"github.com/urfave/cli/v2"
	"io"
)

type App struct {
	app cli.App
}

func New(reader io.Reader, writer, errWriter io.Writer) App {
	return App{
		app: cli.App{
			Reader:    reader,
			Writer:    writer,
			ErrWriter: errWriter,

			Commands: cli.Commands{
				{
					Name: "generate",
					Subcommands: cli.Commands{
						{
							Name: "mos",
							Action: func(cliContext *cli.Context) error {
								fileWriter := filewriters.New("./out")
								cmder := commanders.New("./out")

								indexer, err := texindexer.New("")
								if err != nil {
									return fmt.Errorf("new indexer: %w", err)
								}

								titleSection := mostitles.New(mostitles.TitleParameters{})

								builder := plannerbuilders.New(indexer, plannerbuilders.Sections{titleSection})

								planners.New(builder, fileWriter, cmder)

								panic("not implemented")
							},
						},
					},
				},
			},
		},
	}
}

func (r App) Run(ctx context.Context, arguments []string) error {
	if err := r.app.RunContext(ctx, arguments); err != nil {
		return fmt.Errorf("run context: %w", err)
	}

	return nil
}
