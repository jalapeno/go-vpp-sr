package arangodb

import (
        "fmt"

        "github.com/arangodb/go-driver/http"
        driver "github.com/arangodb/go-driver"

)



func newclient() error {

	conndb, err := http.NewConnection(http.ConnectionConfig{
                Endpoints: []string{"http://10.200.99.27:30852"},
        })

        if err != nil {
		return err
        }

        fmt.Println("conndb: ", conndb)

	client, err := driver.NewClient(driver.ClientConfig{
                Connection: conndb,
        })

        fmt.Println("client:", client)

	if err != nil {
		return err
        }
	return nil
}
