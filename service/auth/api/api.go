package api

import (
	"github.com/micro/cli/v2"
	"github.com/geiqin/go-micro"
	log "github.com/geiqin/go-micro/logger"

	pb "github.com/geiqin/micro/service/auth/api/proto"
)

var (
	// Name of the auth api
	Name = "go.micro.api.auth"
	// Address is the api address
	Address = ":8011"
)

// Run the micro auth api
func Run(ctx *cli.Context, srvOpts ...micro.Option) {
	log.Init(log.WithFields(map[string]interface{}{"service": "auth"}))

	service := micro.NewService(
		micro.Name(Name),
		micro.Address(Address),
	)

	pb.RegisterAuthHandler(service.Server(), NewHandler(service))

	if err := service.Run(); err != nil {
		log.Error(err)
	}
}
