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

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x41\x8f\xdb\x36\x13\x3d\x4b\xbf\x62\x3e\x05\x1f\x2a\x03\xb6\x6c\xef\x36\x9b\xc4\xb2\x0c\x34\x45\xd0\x04\xd8\xf6\x10\x6c\x7a\x68\x51\x2c\x68\x71\x6c\xb1\x4b\x91\x04\x45\x79\xed\x35\xfc\xdf\x0b\x92\x92\x4c\xbb\x69\xd0\x43\x2f\x89\x28\xcd\xbc\x37\xf3\xe6\x0d\xd7\xc7\x23\x50\xdc\x30\x81\x90\xd4\x44\x25\x70\x3a\xc5\xcb\xca\xd4\x7c\x15\x03\x00\x2c\x2b\x24\x74\x15\x47\x4b\xc3\x0c\xc7\x15\xc5\x75\xbb\x5d\x4e\xfd\x21\x8e\x96\x8d\x39\x70\x04\x73\x50\x58\x24\x06\xf7\x66\x5a\x36\x4d\xb2\x8a\x23\x58\x4b\x7a\x80\x63\x1c\xd5\x44\x6f\x99\x58\xc0\x4c\xed\xf3\x38\x52\x84\x52\x26\xb6\xdd\xf1\x14\xc7\x19\x47\xb2\xe1\x68\x26\x86\x71\xb4\xf1\x6b\xa9\x29\xea\x05\xcc\xd5\x1e\x1a\xc9\x19\x85\x35\x6f\x31\x8f\xa3\x35\x29\x9f\xb6\x5a\xb6\x82\x4e\x4a\xc9\xa5\x5e\x28\x26\x9e\x1c\xc6\xab\x9a\x28\x9b\xfa\xcc\xa8\xa9\x16\xf3\xd9\xec\xff\x79\x1c\x55\xc8\xb6\x95\x59\xbc\x9d\xed\x2a\x1f\x54\xb6\x5a\xa3\x30\x13\x83\xb5\xe2\xc4\x38\xb2\x8d\x14\x66\xb2\x21\x35\xe3\x87\x45\x2d\x85\x6c\x14\x29\x2d\x17\x65\x8d\xe2\xe4\xb0\x58\x73\x59\x3e\xe5\x71\x64\x3b\x9b\x10\xce\xb6\x62\x51\xa2\x30\xa8\xf3\xbe\xb1\x89\x91\x6a\x91\xdd\xbc\xd6\x58\x5f\xd1\x74\xfd\xfc\x57\x14\x9e\x21\xf2\xad\xbf\xba\xbb\xbb\xeb\xe8\xa4\x30\x5a\xf2\xc6\x52\xf5\xe2\xf6\xb1\x5e\x0f\x78\x3b\xdb\x3d\xe7\xc1\x20\x80\xb4\x46\x5e\xa2\xdf\xd8\x84\x28\xba\x44\x64\x42\xb5\xe6\xac\x6b\x8f\xe3\x1a\x6a\xd8\x0b\x2e\xe6\x3d\xd1\xdf\xea\x8c\x86\x77\x6b\x69\x8c\xac\xfb\x92\x2e\xf0\xd7\xad\x31\x52\x0c\x1a\x39\xc8\xec\xad\x47\xec\x5c\xe0\x0d\x60\x2a\x26\x42\xef\x64\x77\x21\xed\x62\xa6\xf6\xf0\x3f\x56\x2b\xa9\x0d\x11\xc6\x91\x44\xcb\xa9\x33\x66\x67\xe1\xa9\xf7\xb0\x7b\x76\xbe\xa4\xc4\x10\x37\xa1\x49\xab\x79\x91\x1c\x8f\x90\x3d\x30\x8e\x5f\x3e\xdf\xc3\xe9\x64\xed\xbb\xa4\x6c\x07\x8c\x16\x6e\x23\x56\xcb\x29\x65\x3b\xe7\x76\x45\x84\x7b\x7d\x6d\xa6\x64\x95\x65\x59\xb6\x9c\xda\x80\xaf\x07\x32\x8e\x16\xc8\x05\x44\x01\x41\x2f\x87\x5b\x1a\x5b\xa0\x97\xfd\xbc\x51\x89\x0b\x6b\x35\x7f\x1c\xc8\x40\x90\x1a\xaf\xdf\xed\x08\x6f\xb1\x48\x12\x50\x9c\x94\x58\x49\x4e\x51\x17\xc9\x07\x6b\x25\x20\x02\x08\x37\xa8\x85\xf5\xbd\xed\x72\x58\x82\x0a\x35\x26\x30\xed\xc9\xbb\x99\x78\xf6\xa6\x5d\xd7\xac\xe7\x57\xd4\xb5\xf9\xc5\xfd\xbf\x9c\xfa\x40\xdb\xaa\xd7\xc6\xeb\x6c\xc5\xed\x9e\x9b\x52\x33\x65\x56\xb1\x3b\xc1\x8e\x68\xa8\x89\xca\xcf\x27\x4e\x0e\xd6\xe3\xee\xec\xff\xb5\x5b\x5c\xc0\x7d\x56\x13\x95\x3a\xdd\x47\x59\x83\xe6\x57\x86\xcf\xe9\xef\xb7\x6f\xb2\xbb\xf9\x9b\x77\xef\xc6\x30\x99\xdf\xdc\x64\xb7\x6f\x66\xef\xbe\xbf\xfd\x63\x0c\xf3\xdb\x51\x87\x21\xf0\x19\xee\xb3\x8f\xa4\xa9\xd2\x9a\xa8\x51\x1e\x0f\xa0\x99\x14\x69\x52\x72\x56\x3e\x25\x63\xd8\xb4\xa2\x34\x4c\x8a\x14\x47\xc7\x38\x8e\x7c\x29\x1c\x0a\xc0\x8c\x13\xc3\xc5\x36\xef\x5e\xbe\x48\x59\x43\xe1\xf2\xb7\x68\x7e\x93\xb2\x4e\x2d\xa8\xff\x58\x4a\xa9\x69\xd3\x7d\x56\x5a\xfe\x89\xa5\x49\x39\x1f\xbb\xac\x51\x46\xd9\x8e\x51\x7c\x7f\x48\x6f\x5e\xdf\x8d\xb2\x0d\x97\x52\xdb\xe4\xa8\xcb\xcb\x5e\xa0\x70\x91\x03\x9e\x75\xc7\x63\xab\x6d\x1d\x4e\x17\x4b\xe9\x0c\xa9\x79\xea\x73\x02\x6e\xef\xa8\x47\x77\xc1\x14\x40\x65\xd9\xd6\x28\x8c\x4d\xf9\xc0\xd1\x3e\xbe\x3f\x7c\xa2\xe9\xa5\xf3\xce\xe9\x9c\x89\xa7\x30\xad\xd4\x48\x0c\x76\x99\x69\x42\x12\x57\xa8\x8d\xb2\xea\xff\x60\x8c\x66\xeb\xd6\x60\x9a\x54\x1a\x37\xc9\x78\x28\xf5\x9f\xc2\x0c\xd1\x5b\x34\xc9\x18\x92\xc7\x9e\xb9\x8b\x23\x4a\xa1\xa0\x3f\x56\x8c\xd3\xf4\x8a\xfe\x01\xf7\xe6\x17\x49\x31\x1d\xd0\x5d\x9a\x55\x2c\xe8\x36\x63\x42\xa0\xfe\xf8\xf0\xf3\x3d\x14\x90\x24\xf9\xf5\xe7\x90\xc0\x32\xf6\xd6\x38\x0d\x6e\xb0\x02\x74\x16\xff\x86\x72\x9d\xd7\x47\x81\x5b\x87\x7d\xf9\x56\x5a\xb8\x8e\x61\x72\x57\xe4\xbf\x1a\xd6\x75\xfe\xd5\xc2\x3c\x12\xbd\xb5\xbe\x3b\xc6\x11\xd4\x4c\x58\x5b\x2e\x60\x3e\xb6\x27\xb2\xf7\xa7\x9b\xd9\xb8\x6b\x3b\xec\xba\x41\xf3\xe8\x10\xa0\x38\xef\x40\x58\xb1\x5f\x07\xb6\x81\xd4\x85\x8d\x8e\xfe\x46\x70\x0e\xd7\x58\xcb\x1d\xde\xdb\xf7\xdd\x57\xab\xfd\xc9\x4f\xa8\x47\xbd\xcf\xec\x10\x7c\x50\x08\x3c\x0e\x2a\xf7\x66\x70\x0e\x27\x94\x3e\xc8\x6e\x55\xc3\x49\xfb\x21\x5b\x43\x40\x01\x21\x4e\x3e\x74\x15\x08\xe3\x87\x99\x49\xe1\xf6\x3b\xec\xcd\x35\x10\xe6\x43\x31\x4c\x31\x73\x77\xa5\xe5\x1d\x64\xb9\xd4\x22\xff\x8a\x82\xc1\x8e\x0e\x63\xb4\xf7\x9d\x9d\x65\xe0\xff\x8b\x3f\x2d\xc9\xe0\x3c\x2b\xec\x60\x6e\x5b\x5a\x08\x87\xa5\xa4\xf8\xe5\xf3\xa7\x73\x44\x1e\x4d\xa7\x20\xb0\xc4\xa6\x21\xfa\x00\x6b\x2c\x49\xdb\x20\xc8\x0d\xfc\x24\xbf\x6b\x00\x9b\x92\x28\x26\xb6\xd0\x98\x76\xb3\xb9\xe8\x23\xd8\x4f\xdf\x83\x55\xf7\xac\xd8\x72\xda\x5f\xcc\xcb\xa9\xff\x91\x77\x3c\x02\x0a\x6a\x7f\xf5\xfd\x15\x00\x00\xff\xff\x5c\x39\x3a\x71\x09\x0a\x00\x00")

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

	info := bindataFileInfo{name: "templates/html/index.html", size: 2569, mode: os.FileMode(420), modTime: time.Unix(1567708212, 0)}
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
