// Package dns provides a DNS registration service for autodiscovery of core network nodes.
package dns

import (
	"github.com/geiqin/go-micro"
	log "github.com/geiqin/go-micro/logger"
	"github.com/micro/cli/v2"

	"github.com/geiqin/micro/service/network/dns/handler"
	dns "github.com/geiqin/micro/service/network/dns/proto/dns"
	"github.com/geiqin/micro/service/network/dns/provider/cloudflare"
)

// Run is the entrypoint for network/dns
func Run(c *cli.Context) {

	if c.String("provider") != "cloudflare" {
		log.Fatal("The only implemented DNS provider is cloudflare")
	}

	dnsService := micro.NewService(
		micro.Name("go.micro.network.dns"),
	)

	// Create handler
	provider, err := cloudflare.New(c.String("api-token"), c.String("zone-id"))
	if err != nil {
		log.Fatal(err)
	}
	h := handler.New(
		provider,
		c.String("token"),
	)

	// Register Handler
	dns.RegisterDnsHandler(dnsService.Server(), h)

	// Run service
	if err := dnsService.Run(); err != nil {
		log.Fatal(err)
	}

}
