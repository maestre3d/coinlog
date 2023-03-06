//go:build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/maestre3d/coinlog"
	"github.com/maestre3d/coinlog/domain/card"
	"github.com/maestre3d/coinlog/domain/contact"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/messaging"
	"github.com/maestre3d/coinlog/messaging/kafka"
	"github.com/maestre3d/coinlog/storage/sql"
	"github.com/maestre3d/coinlog/transport/http"
	"github.com/maestre3d/coinlog/transport/stream"
)

var kernelCfgSet = wire.NewSet(
	coinlog.NewConfig,
	http.NewConfig,
	sql.NewConfig,
	kafka.NewConfig,
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

var finAccountSet = wire.NewSet(
	wire.Bind(new(financialaccount.Repository), new(sql.FinancialAccountStorage)),
	sql.NewFinancialAccountStorage,
	financialaccount.NewService,
	http.NewFinancialController,
)

var cardSet = wire.NewSet(
	wire.Bind(new(card.Repository), new(sql.CardStorage)),
	sql.NewCardStorage,
	card.NewService,
	http.NewCardController,
	stream.NewUserController,
)

// Holds all controllers for HTTP protocol, wire auto-binds inner deps.
type httpCtrl struct {
	//Liveness controller.LivenessHTTP
	Healthcheck http.HealthcheckController
	User        http.UserController
	Contact     http.ContactController
	FinAccount  http.FinancialAccountController
	Card        http.CardController
}

func provideHttpRoutes(cfg coinlogHTTPConfig, ctrls httpCtrl) *http.ControllerMapper {
	mapper := http.NewControllerMapper(cfg.Application, cfg.Server)
	// Add desired controllers here in single liner -method accepts variadic-.
	//
	// e.g. mux.Add(ctrls.Report, ctrls.Foo, ctrls.Bar)
	mapper.Add(
		ctrls.Healthcheck,
		ctrls.User,
		ctrls.Contact,
		ctrls.FinAccount,
		ctrls.Card,
	)
	return mapper
}

// Holds all controllers for streams, wire auto-binds inner deps.
type streamCtrl struct {
	User stream.UserController
}

func provideStreamSubscribers(ctrls streamCtrl) *stream.ControllerMapper {
	mapper := stream.NewControllerMapper()
	// Add desired controllers here in single liner -method accepts variadic-.
	//
	// e.g. mux.Add(ctrls.Report, ctrls.Foo, ctrls.Bar)
	mapper.Add(
		ctrls.User,
	)
	return mapper
}

func NewCoinlogHTTP() (*CoinlogHTTP, func(), error) {
	wire.Build(
		kernelCfgSet,
		sql.NewEntClientWithAutoMigrate,
		kafka.NewWriter,
		wire.Bind(new(messaging.Writer), new(kafka.Writer)),
		kafka.NewReader,
		wire.Bind(new(messaging.Reader), new(kafka.Reader)),
		userSet,
		contactSet,
		finAccountSet,
		cardSet,
		http.NewHealthcheckController,
		wire.Struct(new(httpCtrl), "*"),
		provideHttpRoutes,
		wire.Struct(new(streamCtrl), "*"),
		provideStreamSubscribers,
		http.NewEcho,
		stream.NewBus,
		wire.Struct(new(CoinlogHTTP), "*"),
	)
	return nil, nil, nil
}
