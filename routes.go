package main

import (
	"github.com/k8guard/k8guard-report/views"
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func start_http_router() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// When a client closes their connection midway through a request, the
	// http.CloseNotifier will cancel the request context (ctx).
	r.Use(middleware.CloseNotify)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.ThrottleBacklog(20, 5, 20*time.Second))

	r.Route("/", func(r chi.Router) {
		r.Get("/", views.Index)

	})

	r.Route("/last", func(r chi.Router) {
		r.Get("/", views.Last)

	})

	r.Route("/recent", func(r chi.Router) {
		r.Get("/", views.Recent)
		r.Post("/", views.Recent)

	})

	http.ListenAndServe(":3001", r)
}
