package main

import (
    "context"
    "log"

    "terraform-provider-edgeadc/internal/provider"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
    opts := providerserver.ServeOpts{
        Address: "github.com/edgeNEXUS/terraform-provider-edgeadc",
    }

    err := providerserver.Serve(context.Background(), provider.New(), opts)
    if err != nil {
        log.Fatal(err.Error())
    }
}
