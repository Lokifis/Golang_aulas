package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// Gerar retorna aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []*cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereço na internet",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(c *cli.Context) error {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}
