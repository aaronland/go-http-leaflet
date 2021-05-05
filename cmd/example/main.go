package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-leaflet/templates/html"
	"html/template"
	"log"
	"net/http"
)

type MapVars struct {
	TileURL string
}

func MapHandler(templates *template.Template, map_vars *MapVars) (http.Handler, error) {

	t := templates.Lookup("map")

	if t == nil {
		return nil, errors.New("Missing 'map' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		err := t.Execute(rsp, map_vars)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn), nil
}

func main() {

	host := flag.String("host", "localhost", "...")
	port := flag.Int("port", 8080, "...")

	tile_url := flag.String("tile-url", "", "A valid Leaflet layer tile URL")

	flag.Parse()

	t, err := template.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	map_vars := new(MapVars)

	if *tile_url != "" {
		map_vars.TileURL = *tile_url
	}

	mux := http.NewServeMux()

	map_handler, err := MapHandler(t, map_vars)

	if err != nil {
		log.Fatalf("Failed to create map handler, %v", err)
	}

	leaflet_opts := leaflet.DefaultLeafletOptions()
	leaflet_opts.EnableHash()
	leaflet_opts.EnableFullscreen()
	
	map_handler = leaflet.AppendResourcesHandler(map_handler, leaflet_opts)

	mux.Handle("/", map_handler)

	err = leaflet.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Leaflet asset handler, %v", err)
	}

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", endpoint)

	err = http.ListenAndServe(endpoint, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}

}
