package leaflet

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/aaronland/go-http-leaflet/static"
	aa_static "github.com/aaronland/go-http-static"
	"github.com/sfomuseum/go-http-rollup"
)

var ROLLUP_ASSETS = false

// LeafletOptions provides a list of JavaScript and CSS link to include with HTML output.
type LeafletOptions struct {
	JS             []string
	CSS            []string
	DataAttributes map[string]string
	// AppendJavaScriptAtEOF is a boolean flag to append JavaScript markup at the end of an HTML document
	// rather than in the <head> HTML element. Default is false
	AppendJavaScriptAtEOF bool
}

// Append the Javascript and CSS URLs for the Leaflet.Fullscreen plugin.
func (opts *LeafletOptions) EnableFullscreen() {
	opts.CSS = append(opts.CSS, "/css/leaflet.fullscreen.css")
	opts.JS = append(opts.JS, "/javascript/leaflet.fullscreen.min.js")
}

// Append the Javascript and CSS URLs for the Leaflet.Hash plugin.
func (opts *LeafletOptions) EnableHash() {
	opts.JS = append(opts.JS, "/javascript/leaflet-hash.js")
}

// Append the Javascript and CSS URLs for the leaflet-geoman plugin.
// https://github.com/geoman-io/leaflet-geoman/
func (opts *LeafletOptions) EnableDraw() {
	opts.CSS = append(opts.CSS, "/css/leaflet-geoman.css")
	opts.JS = append(opts.JS, "/javascript/leaflet-geoman.min.js")
}

// Return a *LeafletOptions struct with default paths and URIs.
func DefaultLeafletOptions() *LeafletOptions {

	opts := &LeafletOptions{
		CSS: []string{
			"/css/leaflet.css",
		},
		JS: []string{
			"/javascript/leaflet.js",
		},
		DataAttributes: make(map[string]string),
	}

	return opts
}

// AppendResourcesHandler will rewrite any HTML produced by previous handler to include the necessary markup to load Leaflet JavaScript and CSS files and related assets.
func AppendResourcesHandler(next http.Handler, opts *LeafletOptions) http.Handler {
	return AppendResourcesHandlerWithPrefix(next, opts, "")
}

// AppendResourcesHandlerWithPrefix will rewrite any HTML produced by previous handler to include the necessary markup to load Leaflet JavaScript files and related assets ensuring that all URIs are prepended with a prefix.
func AppendResourcesHandlerWithPrefix(next http.Handler, opts *LeafletOptions, prefix string) http.Handler {

	static_opts := aa_static.DefaultResourcesOptions()
	static_opts.DataAttributes = opts.DataAttributes
	static_opts.AppendJavaScriptAtEOF = opts.AppendJavaScriptAtEOF

	if ROLLUP_ASSETS {

		static_opts.CSS = []string{
			"/css/leaflet.rollup.css",
		}

		static_opts.JS = []string{
			"/css/leaflet.rollup.js",
		}

	} else {

		static_opts.CSS = opts.CSS
		static_opts.JS = opts.JS
	}

	return aa_static.AppendResourcesHandlerWithPrefix(next, static_opts, prefix)
}

// Append all the files in the net/http FS instance containing the embedded Leaflet assets to an *http.ServeMux instance.
func AppendAssetHandlers(mux *http.ServeMux) error {

	return AppendAssetHandlersWithPrefix(mux, "")
}

// Append all the files in the net/http FS instance containing the embedded Leaflet assets to an *http.ServeMux instance ensuring that all URLs are prepended with prefix.
func AppendAssetHandlersWithPrefix(mux *http.ServeMux, prefix string) error {

	if !ROLLUP_ASSETS {
		return aa_static.AppendStaticAssetHandlersWithPrefix(mux, static.FS, prefix)
	}

	logger := log.Default()

	rollupjs_paths := map[string][]string{
		"leaflet.rollup.js": []string{},
	}

	rollupjs_opts := &rollup.RollupJSHandlerOptions{
		FS:     static.FS,
		Paths:  rollupjs_paths,
		Logger: logger,
	}

	rollupjs_handler, err := rollup.RollupJSHandler(rollupjs_opts)

	if err != nil {
		return fmt.Errorf("Failed to create rollup JS handler, %w", err)
	}

	rollupjs_uri := "/javascript/leaflet.rollup.js"

	if prefix != "" {

		u, err := url.JoinPath(prefix, rollupjs_uri)

		if err != nil {
			return fmt.Errorf("Failed to append prefix to %s, %w", rollupjs_uri, err)
		}

		rollupjs_uri = u
	}

	mux.Handle(rollupjs_uri, rollupjs_handler)

	// CSS

	rollupcss_paths := map[string][]string{
		"leaflet.rollup.css": []string{},
	}

	rollupcss_opts := &rollup.RollupCSSHandlerOptions{
		FS:     static.FS,
		Paths:  rollupcss_paths,
		Logger: logger,
	}

	rollupcss_handler, err := rollup.RollupCSSHandler(rollupcss_opts)

	if err != nil {
		return fmt.Errorf("Failed to create rollup CSS handler, %w", err)
	}

	rollupcss_uri := "/javascript/leaflet.rollup.css"

	if prefix != "" {

		u, err := url.JoinPath(prefix, rollupcss_uri)

		if err != nil {
			return fmt.Errorf("Failed to append prefix to %s, %w", rollupcss_uri, err)
		}

		rollupcss_uri = u
	}

	mux.Handle(rollupcss_uri, rollupcss_handler)
	return nil
}
