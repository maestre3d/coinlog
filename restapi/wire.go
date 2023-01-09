//go:build wireinject

package restapi

import (
	"github.com/google/wire"
	"github.com/maestre3d/coinlog/appservice"
	"github.com/maestre3d/coinlog/configuration"
	"github.com/maestre3d/coinlog/controller"
	"github.com/maestre3d/coinlog/persistence"
	"github.com/maestre3d/coinlog/repository"
)

var kernelCfgSet = wire.NewSet(
	configuration.NewApplication,
	configuration.NewServerHTTP,
	configuration.NewDatabaseSQL,
	wire.Struct(new(coinlogHTTPConfig), "*"),
)

var userSet = wire.NewSet(
	wire.Bind(new(repository.User), new(persistence.UserSQL)),
	persistence.NewUserSQL,
	appservice.NewUser,
	controller.NewUserHTTP,
)

var contactSet = wire.NewSet(
	wire.Bind(new(repository.Contact), new(persistence.ContactSQL)),
	persistence.NewContactSQL,
	appservice.NewContact,
	controller.NewContactHTTP,
)

// Holds all controllers for HTTP protocol, wire auto-binds inner deps.
type httpCtrl struct {
	Liveness controller.LivenessHTTP
	User     controller.UserHTTP
	Contact  controller.ContactHTTP
}

func provideHttpRoutes(cfg coinlogHTTPConfig, ctrls httpCtrl) *controller.MuxHTTP {
	mux := controller.NewMux(cfg.Application, cfg.Server)
	// Add desired controllers here in single liner -method accepts variadic-.
	//
	// e.g. mux.Add(ctrls.Report, ctrls.Foo, ctrls.Bar)
	mux.Add(ctrls.Liveness, ctrls.User, ctrls.Contact)
	return mux
}

func NewCoinlogHTTP() (*CoinlogHTTP, func(), error) {
	wire.Build(
		kernelCfgSet,
		persistence.NewEntClient,
		userSet,
		contactSet,
		controller.NewLivenessHTTP,
		wire.Struct(new(httpCtrl), "*"),
		provideHttpRoutes,
		controller.NewEcho,
		wire.Struct(new(CoinlogHTTP), "*"),
	)
	return nil, nil, nil
}
