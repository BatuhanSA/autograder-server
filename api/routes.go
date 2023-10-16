package api

// All routes handled by the server.

import (
    "github.com/eriq-augustine/autograder/api/core"
    "github.com/eriq-augustine/autograder/api/user"
)

var baseRoutes = []*core.Route{
    core.NewRedirect("GET", ``, `/static/index.html`),
    core.NewRedirect("GET", `/`, `/static/index.html`),
    core.NewRedirect("GET", `/index.html`, `/static/index.html`),

    core.NewRoute("GET", `/static`, handleStatic),
    core.NewRoute("GET", `/static/.*`, handleStatic),
}

func GetRoutes() *[]*core.Route {
    routes := make([]*core.Route, 0);

    routes = append(routes, baseRoutes...);
    routes = append(routes, *(user.GetRoutes())...);

    return &routes;
}