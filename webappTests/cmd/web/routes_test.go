package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"testing"
)

func Test_application_routes(t *testing.T) {
	tests := []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range tests {
		// check to see if the route exits
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s not registered", route.route)
		}

	}

}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, hnadler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})
	return found
}
