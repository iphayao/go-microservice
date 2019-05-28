package main

import (
	"fmt"
	"os"
	"context"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	proto "github.com/iphayao/go-microservice/proto"
)

type Greeter struct {
	
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	res.Greeting = "Hello " + req.Name
	return nil
}

// setup the client
func runClient(service micro.Service) {
	// create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())
	// call greeter
	res, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	// prince response
	fmt.Println(res.Greeting)
}

func main() {
	// create service
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
		micro.Flags(cli.BoolFlag{
			Name: "run_client",
			Usage: "Launch the client",
		}),
	)
	
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}