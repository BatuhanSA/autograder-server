package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/edulinq/autograder/internal/common"
	"github.com/edulinq/autograder/internal/config"
	"github.com/edulinq/autograder/internal/db"
	"github.com/edulinq/autograder/internal/log"
	"github.com/edulinq/autograder/internal/procedures/courses"
	"github.com/edulinq/autograder/internal/util"
)

var args struct {
	config.ConfigArgs
	Source string `help:"The source to add a course from." arg:""`
	DryRun bool   `help:"Do not actually do the operation, just state what you would do." default:"false"`
}

func main() {
	kong.Parse(&args,
		kong.Description("Add a course to system from a source (FileSpec)."),
	)

	err := config.HandleConfigArgs(args.ConfigArgs)
	if err != nil {
		log.Fatal("Could not load config options.", err)
	}

	db.MustOpen()
	defer db.MustClose()

	spec, err := common.ParseFileSpec(args.Source)
	if err != nil {
		log.Fatal("Failed to parse FileSpec.", err)
	}

	options := courses.CourseUpsertOptions{
		ContextUser: db.MustGetRoot(),
		DryRun:      args.DryRun,
	}

	results, err := courses.UpsertFromFileSpec(spec, options)
	if err != nil {
		log.Fatal("Failed to add courses from FileSpec.", err)
	}

	fmt.Println(util.MustToJSONIndent(results))
}