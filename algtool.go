/*
MIT License

Copyright (c) 2017 Rob Baruch

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"encoding/json"
	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"os"
	"fmt"
	"io/ioutil"
	"flag"
)


func main(){

	api_id := flag.String("id", "", "algolia api id string")
	api_key := flag.String("key", "", "algolia api key string")
	indexName := flag.String("index", "default-content", "algolia index name")
	verbose := flag.Bool("verbose", false, "enable verbosity")
	flag.Parse()

	if *api_id == "" || *api_key ==  "" || *indexName == "" {
		fmt.Printf("Error, must specify id and key\n")
		os.Exit(1)
	}

	if *verbose == true {
		fmt.Printf("\t\tAPI_ID = [%s]\n", *api_id)
		fmt.Printf("\t\tAPI_KEY= [%s]\n", *api_key)
		fmt.Printf("\t\tindex  = [%s]\n", *indexName)
	}

	client := algoliasearch.NewClient(*api_id, *api_key)
	index := client.InitIndex(*indexName)

	jsonData, err := ioutil.ReadFile("./public/index.json")
	if err != nil {
		fmt.Printf("Error reading file: %s", err)
		os.Exit(1)
	}

	var objects []algoliasearch.Object
	if err := json.Unmarshal(jsonData, &objects); err != nil {
		fmt.Printf("Error unmarshalling json file: %s", err)
		os.Exit(1)
	}

	res, err := index.UpdateObjects(objects)

	if err!= nil {
		fmt.Printf("Error adding objects: %s", err)
		os.Exit(1)
	} else {
		fmt.Printf("Res: %q\n", res)
	}

	os.Exit(0)
}
