package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/plasmatrip/muslib/internal/api/middleware/compress"
	"github.com/plasmatrip/muslib/internal/api/middleware/logger"
)

func NewRouter(deps deps.Dependencies, controller *controller.Controller) *chi.Mux {

	r := chi.NewRouter()

	balance := balance.NewBalanceService(deps)
	orders := orders.NewOrdersService(deps)
	info := info.NewInfoService(deps)

	r.Use(logger.WithLogging(deps.Logger))
	r.Use(compress.WithCompressed)

	r.Route("/api/user/orders", func(r chi.Router) {
		r.Use(auth.Validate)
		r.Post("/", orders.AddOrder)
		r.Get("/", orders.GetOrders)
	})

	r.Route("/api/user/balance", func(r chi.Router) {
		r.Use(auth.Validate)
		r.Get("/", balance.GetBalance)
	})

	r.Route("/api/user/balance/withdraw", func(r chi.Router) {
		r.Use(auth.Validate)
		r.Post("/", balance.Withdraw)
	})

	r.Route("/api/user/withdrawals", func(r chi.Router) {
		r.Use(auth.Validate)
		r.Get("/", balance.Withdrawals)
	})

	r.Route("/api/info", func(r chi.Router) {
		r.Get("/", info.Ping)
	})

	return r
}
