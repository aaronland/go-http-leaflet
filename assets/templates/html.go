// Code generated by go-bindata.
// sources:
// templates/html/index.html
// DO NOT EDIT!

package templates

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\xdf\x6f\xdb\x38\x12\x7e\x96\xff\x8a\x39\x15\x87\xc8\x80\x4d\xdb\xc9\x35\x6d\x65\x2b\xc0\xb5\xe8\x5d\x0b\x64\xf7\xa1\x48\xf7\x61\x17\x8b\x80\x16\xc7\x36\x37\x14\x49\x50\x94\x63\xd7\xf0\xff\xbe\x20\x29\x4b\xb4\x9b\x2d\xf6\x61\xfb\xd0\x88\xd4\xcc\x7c\xf3\xe3\xfb\x46\x3e\x1c\x80\xe1\x8a\x4b\x84\xb4\xa2\x3a\x85\xe3\x71\xb0\xd8\xd8\x4a\xdc\x0d\x00\x00\x16\x1b\xa4\xec\x6e\x90\x2c\x2c\xb7\x02\xef\x18\x2e\x9b\xf5\x62\x12\x0e\x83\x64\x51\xdb\xbd\x40\xb0\x7b\x8d\x45\x6a\x71\x67\x27\x65\x5d\xa7\x77\x83\x04\x96\x8a\xed\xe1\x30\x48\x12\xa8\xa8\x59\x73\x99\xc3\x54\xef\xe6\xee\xac\x29\x63\x5c\xae\x4f\x17\x70\x4c\x60\x30\x48\x80\x08\xa4\x2b\x81\x76\x6c\xb9\xc0\xe0\xb9\x54\x86\xa1\xc9\x61\xa6\x77\x50\x2b\xc1\x19\x2c\x45\x83\x3e\xc8\x92\x96\x4f\x6b\xa3\x1a\xc9\xc6\xa5\x12\xca\xe4\x9a\xcb\x27\x1f\x6d\x90\xc0\x20\x81\x57\x15\xd5\x21\xc8\x33\x67\x76\x93\xcf\xa6\xd3\x7f\x7b\xc7\x0d\xf2\xf5\xc6\xe6\x6f\xa7\xdb\x4d\x6c\x5e\x36\xc6\xa0\xb4\x63\x8b\x95\x16\xd4\xb6\x09\xac\x94\xb4\xe3\x15\xad\xb8\xd8\xe7\x95\x92\xaa\xd6\xb4\x0c\xf8\x8c\xd7\x5a\xd0\x7d\xbe\x14\xaa\x7c\xf2\x37\xae\xfa\x31\x15\x7c\x2d\xf3\x12\xa5\x45\x33\xef\x8b\x1f\x5b\xa5\x73\x72\xfd\xda\x60\xf5\x22\x68\x57\xf1\x3f\x0b\xd8\xe2\x25\x10\x5a\xf4\xea\xf6\xf6\xf6\x0c\x5e\x49\x6b\x94\xa8\x03\xf4\x69\x2c\xbd\x57\xe8\x1c\xbc\x9d\x6e\x9f\xe7\x67\x83\x04\xda\x58\x75\x89\x76\xed\xdc\x92\xe4\xa5\xf8\x5c\xea\xc6\xc6\xd3\xe8\x63\xfa\x82\x6b\xfe\x0d\xf3\x59\x0f\xfc\x5d\x0d\x49\x74\xbb\x54\xd6\xaa\xaa\x4b\xf3\x7b\xb4\x65\x63\xad\x92\x51\x3f\x7d\x78\xf2\xf6\x14\xbd\x65\x55\x20\x94\xdd\x70\x79\xce\x4a\x72\x7b\x9e\x46\x3e\xd5\x3b\xf8\x17\xaf\xb4\x32\x96\x4a\xdb\x43\x2e\x26\x9e\xfc\xad\x4c\x26\x41\x27\xfe\xd9\x73\x9f\x51\x4b\xfd\x64\xc7\x8d\x11\x45\x7a\x38\x00\x79\xe0\x02\xbf\x7e\xb9\x87\xe3\xd1\x49\x64\xc1\xf8\x16\x38\x2b\xbc\xea\xee\x16\x13\xc6\xb7\x5e\x51\x9a\x4a\x7f\x7d\x49\xc9\xf4\x8e\x10\x42\x16\x13\x67\xf0\xb2\x21\x17\xe8\x02\x79\x83\x24\x02\x38\xb5\xc6\x0b\xd3\x25\x18\x06\xd2\xab\x36\xf5\x66\x8d\x11\x8f\x1d\x18\x48\x5a\xe1\xe5\xdd\x96\x8a\x06\x8b\x34\x05\x2d\x68\x89\x1b\x25\x18\x9a\x22\xfd\xe8\xc8\x07\x54\x02\x15\x16\x8d\x74\xea\x71\x55\x76\x52\xda\xa0\xc1\x14\x26\x27\xf0\x76\x3e\x01\xbd\x6e\x96\x15\x3f\xe1\x6b\xe6\xcb\xfc\xea\xff\x2e\x26\xc1\xd0\x95\x1a\x7a\x13\xfa\xec\x9a\xdb\x3e\xd7\xa5\xe1\xda\x86\x03\x84\xff\xb7\xd4\x40\x45\xf5\xbc\x3f\x09\xba\x77\xd2\x88\x6c\xdc\x6a\x28\xe0\x9e\x54\x54\x67\xbe\xf9\x43\x52\xa3\xfd\x85\xe3\x73\xf6\xdb\xcd\x1b\x72\x3b\x7b\xf3\xee\xdd\x08\xc6\xb3\xeb\x6b\x72\xf3\x66\xfa\xee\x3f\x37\xbf\x8f\x60\x76\x33\x6c\x63\x48\x7c\x86\x7b\xf2\x89\xd6\x9b\xac\xa2\x7a\x38\x1f\x74\x41\x09\x65\xec\x43\x68\x75\x16\xac\xda\x13\xf9\x5f\x23\x44\x5d\x1a\x44\x99\x1d\x06\x09\xf8\xed\x99\xc3\x21\x78\x9e\xfe\x5d\xad\xa8\xa8\xf1\x2a\x87\x2b\x97\x0a\xf4\x3e\x57\xa3\x0b\x43\x6b\x1a\x6f\xf7\x71\xc7\x6d\x6c\xe7\xa9\xe9\x4d\x8e\xc3\xe1\x65\xc9\x44\xc9\x2c\x2d\x05\x2f\x9f\xd2\x11\xac\x1a\x59\x5a\xae\x64\x86\xc3\x83\x5b\xbe\xbe\x51\x02\x0a\x40\x22\xa8\x15\x72\x3d\x6f\x2f\xbf\x29\x55\x41\xe1\xfd\xd7\x68\x7f\x55\xaa\xca\x5c\xc9\xe1\x65\xa9\x94\x61\x75\xfb\x5a\x1b\xf5\x07\x96\x36\x13\x62\xe4\xbd\x86\x84\xf1\x2d\x67\xf8\x7e\x9f\x5d\xbf\xbe\x1d\x92\x95\x50\xca\x38\xe7\xa4\xf5\x23\xdf\xa0\xf0\x96\x5d\x3c\x47\xe0\xc7\xc6\xb8\x3c\xfc\xd4\x1c\xa4\xd7\x8c\x11\x59\xf0\x89\xb0\x03\xe9\x1f\xfd\xee\x2c\x80\xa9\xb2\xa9\x50\x5a\xe7\xf2\x51\xa0\x7b\x7c\xbf\xff\xcc\xb2\x73\x71\xf4\xee\x82\xcb\xa7\xd8\xad\x34\x48\x2d\xb6\x9e\x59\x4a\x53\x9f\xa8\xb3\x72\xdc\xf8\xaf\xb5\x86\x2f\x1b\x8b\x59\xba\x31\xb8\x4a\x47\x5d\xaa\x7f\x65\x66\xa9\x59\xa3\x4d\x47\x90\x3e\x9e\x90\x5b\x3b\xaa\x35\x4a\xf6\x61\xc3\x05\xcb\x2e\xe0\x1f\x70\x67\x7f\x56\x0c\xb3\x2e\xba\x77\x73\x1d\x8b\xaa\x25\x5c\x4a\x34\x9f\x1e\x7e\xba\x87\x02\xd2\x74\x7e\xf9\x3a\x06\x70\x88\x27\x26\x1c\x3b\xae\xba\x06\xb4\x2a\xfc\x41\xe7\x5a\x39\x0e\x23\x2d\x75\x92\xfe\x91\x5b\xbc\x31\x62\xe7\x36\xc9\xbf\x35\xac\x4b\xff\x0b\x39\x3f\x52\xb3\x76\xbc\x73\x62\xaa\xb8\x74\xb4\xcc\x61\x36\x72\x27\xba\x0b\xa7\xeb\x69\xab\x9a\x63\x5c\x75\x8d\xf6\xd1\x47\x80\xa2\xd7\x40\x9c\x71\x90\x03\x5f\x41\xe6\xcd\x86\x87\xb0\xb4\x3c\xc3\x0d\x56\x6a\x8b\xf7\xee\xbe\x7d\x1b\x7d\x7f\x4e\x51\xef\x89\x1b\x42\x30\x8a\x03\x8f\xa2\xcc\x03\x19\x3c\xc3\x29\x63\x0f\xaa\x5d\x24\xf1\xa4\xc3\x90\x1d\x21\xa0\x80\x38\xce\xbc\xab\x2a\x6a\x4c\x18\x26\x51\xd2\xeb\x3b\xae\xcd\x17\x10\xfb\x43\xd1\x4d\x91\xf8\x75\xee\x70\xbb\xb6\x9c\xf7\x62\xfe\x42\x07\x23\x8d\x76\x63\x74\x2b\xd9\xcd\x32\xe2\xff\xd9\xd7\x2f\xed\x98\xe7\x1a\xdb\x91\x3b\xac\xc2\x3e\x1c\x96\x8a\xe1\xd7\x2f\x9f\x7b\x8b\x79\x32\x99\x80\xc4\x12\xeb\x9a\x9a\x3d\x2c\xb1\xa4\x4d\x8d\xa0\x56\xf0\x7f\x75\x55\x03\xd6\x25\xd5\x5c\xae\xa1\xb6\xcd\x6a\x75\x56\x47\xa4\xcf\x50\x83\xeb\x6e\xdf\xb1\xc5\xe4\xf4\xed\x58\x4c\xc2\x6f\xdd\xc3\x01\x50\x32\xf7\xe3\xf7\xcf\x00\x00\x00\xff\xff\x56\x04\xd0\x91\x10\x0b\x00\x00")

func templatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIndexHtml,
		"templates/html/index.html",
	)
}

func templatesHtmlIndexHtml() (*asset, error) {
	bytes, err := templatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/index.html", size: 2832, mode: os.FileMode(420), modTime: time.Unix(1577735450, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/html/index.html": templatesHtmlIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

