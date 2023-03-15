package arangodb

import (
        "fmt"
	"context"

        "github.com/arangodb/go-driver/http"
        driver "github.com/arangodb/go-driver"

)


type Sid struct {
	Srv6Sid string `json:"srv6_sid"`
}


func Newclient() error {

	conndb, err := http.NewConnection(http.ConnectionConfig{
                Endpoints: []string{"http://10.200.99.27:32748/"},
        })

        if err != nil {
		return err
        }

        fmt.Println("conndb: ", conndb)

	client, err := driver.NewClient(driver.ClientConfig{
                Connection: conndb,
		Authentication: driver.BasicAuthentication("root", "jalapeno"),

        })

        fmt.Println("client:", client)

	if err != nil {
		return err
        }

	db, err := client.Database(context.TODO(), "jalapeno")

        fmt.Println("db: %s", db.Name())

	if err != nil {
		return err
        }

	col, err := db.Collection(context.TODO(), "ls_srv6_sid")

	if err != nil {
		return err
        }

	fmt.Printf("collection name: %s", col.Name())
	srv6_query := "for s in " + col.Name() + " return s"
	cursor, err := db.Query(context.TODO(), srv6_query, nil)
	if err != nil {
    		return err
	}
	defer cursor.Close()
	docs := make([]Sid,0)
	for {
  		var doc Sid
		meta, err := cursor.ReadDocument(context.TODO(), &doc)
  		if driver.IsNoMoreDocuments(err) {
     		break
  		} else if err != nil {
			return err
 		}
  		fmt.Printf("Got doc with key '%s' from query\n", meta.Key)
  		fmt.Printf("Got doc with key '%s' from query\n", doc)
		docs = append(docs,doc)
	}
	for _, s :=range docs {
		fmt.Printf("Sid: %s %T\n", s.Srv6Sid, s.Srv6Sid)
	}
	return nil
}
