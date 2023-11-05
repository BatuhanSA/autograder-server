package main

import (
    "fmt"
    "os"

    "github.com/alecthomas/kong"
    "github.com/rs/zerolog/log"

    "github.com/eriq-augustine/autograder/config"
    "github.com/eriq-augustine/autograder/db"
)

var args struct {
    config.ConfigArgs
    Course string `help:"ID of the course." arg:""`
    Email string `help:"Email of the user to fetch." arg:""`
}

func main() {
    kong.Parse(&args,
        kong.Description("Fetch users for a specific LMS course."),
    );

    err := config.HandleConfigArgs(args.ConfigArgs);
    if (err != nil) {
        log.Fatal().Err(err).Msg("Could not load config options.");
    }

    db.MustOpen();
    defer db.MustClose();

    course := db.MustGetCourse(args.Course);
    if (course.GetLMSAdapter() == nil) {
        fmt.Println("Course has no LMS info associated with it.");
        os.Exit(2);
    }

    user, err := course.GetLMSAdapter().FetchUser(args.Email);
    if (err != nil) {
        log.Fatal().Err(err).Msg("Could not fetch user.");
    }

    if (user == nil) {
        fmt.Println("No user found.");
        return;
    }

    fmt.Println("id\temail\tname\trole");
    fmt.Printf("%s\t%s\t%s\t%s\n", user.ID, user.Email, user.Name, user.Role.String());
}
