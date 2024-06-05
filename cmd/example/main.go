package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/aaronland/go-http-leaflet"
	"github.com/aaronland/go-http-leaflet/templates/html"
)

type ExampleVars struct {
	TileURL          string
	EnableHash       bool
	EnableFullscreen bool
	EnableDraw       bool
}

func ExampleHandler(templates *template.Template, example_vars *ExampleVars) (http.Handler, error) {

	t := templates.Lookup("example")

	if t == nil {
		return nil, errors.New("Missing 'example' template")
	}

	fn := func(rsp http.ResponseWriter, req *http.Request) {

		err := t.Execute(rsp, example_vars)

		if err != nil {
			http.Error(rsp, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	return http.HandlerFunc(fn), nil
}

func main() {

	host := flag.String("host", "localhost", "The host name to listen for requests on.")
	port := flag.Int("port", 8080, "The host port to listen for requests on.")

	tile_url := flag.String("tile-url", "", "A valid Leaflet layer tile URL")

	enable_hash := flag.Bool("enable-hash", false, "Enable the Leaflet Hash plugin")
	enable_fullscreen := flag.Bool("enable-fullscreen", false, "Enable the Leaflet Fullscreen plugin")
	enable_draw := flag.Bool("enable-draw", false, "Enable the Leaflet Draw plugin")

	js_eof := flag.Bool("javascript-at-eof", false, "A boolean flag to include JavaScript resources (<script> tags) at the end of HTML output.")
	rollup_assets := flag.Bool("rollup-assets", false, "Rollup (minify and bundle) JavaScript and CSS assets.")

	flag.Parse()

	t, err := template.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	mux := http.NewServeMux()

	leaflet_opts := leaflet.DefaultLeafletOptions()

	leaflet_opts.AppendJavaScriptAtEOF = *js_eof
	leaflet_opts.RollupAssets = *rollup_assets

	if *enable_hash {
		leaflet_opts.EnableHash()
	}

	if *enable_fullscreen {
		leaflet_opts.EnableFullscreen()
	}

	if *enable_draw {
		leaflet_opts.EnableDraw()
	}

	err = leaflet.AppendAssetHandlers(mux, leaflet_opts)

	if err != nil {
		log.Fatalf("Failed to append Leaflet assets handler, %v", err)
	}

	example_vars := &ExampleVars{
		TileURL:          *tile_url,
		EnableHash:       *enable_hash,
		EnableFullscreen: *enable_fullscreen,
		EnableDraw:       *enable_draw,
	}

	example_handler, err := ExampleHandler(t, example_vars)

	if err != nil {
		log.Fatalf("Failed to create example handler, %v", err)
	}

	example_handler = leaflet.AppendResourcesHandler(example_handler, leaflet_opts)

	mux.Handle("/", example_handler)

	endpoint := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", endpoint)

	err = http.ListenAndServe(endpoint, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}

}
