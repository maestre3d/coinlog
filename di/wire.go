//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/maestre3d/coinlog"
	"github.com/maestre3d/coinlog/domain/contact"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/storage/sql"
	"github.com/maestre3d/coinlog/transport/http"
)

var kernelCfgSet = wire.NewSet(
	coinlog.NewConfig,
	http.NewConfig,
	sql.NewConfig,
	wire.Struct(new(coinlogHTTPConfig), "*"),
)

var userSet = wire.NewSet(
	wire.Bind(new(user.Repository), new(sql.UserStorage)),
	sql.NewUserStorage,
	user.NewService,
	http.NewUserController,
)

var contactSet = wire.NewSet(
	wire.Bind(new(contact.Repository), new(sql.ContactStorage)),
	sql.NewContactStorage,
	contact.NewService,
	http.NewContactController,
)

// Holds all controllers for HTTP protocol, wire auto-binds inner deps.
type httpCtrl struct {
	//Liveness controller.LivenessHTTP
	Healthcheck http.HealthcheckController
	User        http.UserController
	Contact     http.ContactController
}

func provideHttpRoutes(cfg coinlogHTTPConfig, ctrls httpCtrl) *http.ControllerMapper {
	mapper := http.NewControllerMapper(cfg.Application, cfg.Server)
	// Add desired controllers here in single liner -method accepts variadic-.
	//
	// e.g. mux.Add(ctrls.Report, ctrls.Foo, ctrls.Bar)
	mapper.Add(ctrls.Healthcheck, ctrls.User, ctrls.Contact)
	return mapper
}

func NewCoinlogHTTP() (*CoinlogHTTP, func(), error) {
	wire.Build(
		kernelCfgSet,
		sql.NewEntClient,
		userSet,
		contactSet,
		http.NewHealthcheckController,
		wire.Struct(new(httpCtrl), "*"),
		provideHttpRoutes,
		http.NewEcho,
		wire.Struct(new(CoinlogHTTP), "*"),
	)
	return nil, nil, nil
}
