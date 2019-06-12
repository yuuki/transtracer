// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../../CREDITS
package main

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

var _Credits = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\x51\x73\xea\xb8\x15\xc7\xdf\x35\xc3\x77\x38\x73\x5f\xf6\xa6\xe3\x85\x7b\xb7\xed\x4c\x7b\xb7\xed\xd4\xd8\x0a\x68\xd6\xd8\xd4\x16\xc9\xe6\xa9\x63\x6c\x81\xd5\x6b\x2c\x56\x12\x49\xe9\xa7\xef\x48\x36\xc1\x24\x64\x9b\x1b\x7b\xdb\xde\x1d\xf2\x12\x82\x2c\xe9\x9c\x63\xf9\xf7\x3f\xfc\xc9\x44\xc0\x7b\x5d\x30\x50\x3a\xad\xf2\x54\xe6\x50\xf2\xa5\x4c\xe5\xfe\x0a\x15\x5a\x6f\xd5\xa7\xd1\x68\x2d\xca\xb4\x5a\x0f\x85\x5c\x8f\xd0\xb7\x1d\x7f\x90\x27\xb6\x7b\xc9\xd7\x85\x86\xf7\xd9\x15\x7c\xf7\xe1\xc3\x1f\x81\x16\x0c\x26\x02\xdc\x9d\x2e\x84\x54\x43\x70\xcb\x12\xec\x25\x0a\x24\x53\x4c\xde\xb3\x7c\x88\x50\xcc\x72\xae\xb4\xe4\xcb\x9d\xe6\xa2\x82\xb4\xca\x61\xa7\x18\xf0\x0a\x94\xd8\xc9\x8c\xd9\x77\x96\xbc\x4a\xe5\x1e\x56\x42\x6e\x94\x03\x0f\x5c\x17\x20\xa4\xfd\x2d\x76\x1a\x6d\x44\xce\x57\x3c\x4b\xcd\x02\x0e\xa4\x92\xc1\x96\xc9\x0d\xd7\x9a\xe5\xb0\x95\xe2\x9e\xe7\x2c\x07\x5d\xa4\x1a\x4c\x41\x56\xa2\x2c\xc5\x03\xaf\xd6\x90\x89\x2a\xe7\x66\x92\x32\x93\xd0\x86\xe9\x4f\x08\x01\xc0\x6f\xe0\x34\x28\x05\x62\x75\x88\x26\x13\x39\x83\xcd\x4e\x69\x90\x4c\xa7\xbc\xb2\x4b\xa6\x4b\x71\x6f\x86\x9a\x12\xa0\x4a\x68\x9e\x31\x07\x74\xc1\x15\x94\x5c\x69\xb3\x40\x7b\xb7\x2a\x7f\x12\x4a\xce\x55\x56\xa6\x7c\xc3\xe4\xf0\x7c\x04\xbc\x6a\x17\xe1\x10\xc1\x56\x8a\x7c\x97\xb1\x63\x10\xe8\x31\x08\xe8\x12\x04\x6a\x12\xcb\x45\xb6\xdb\xb0\x4a\xa7\x87\x7b\x33\x12\x12\x84\x2e\x98\x84\x4d\xaa\x99\xe4\x69\xa9\x8e\x25\xb6\xf7\x45\x17\x0c\xb5\x43\x6f\xf2\x09\x19\xb7\xd3\xcc\xaa\x55\xba\x61\x26\x98\x89\x10\xeb\x92\x01\xa9\xb2\x21\x54\xe2\x38\x66\xeb\xcd\xb5\x42\x99\xa8\xea\x75\x84\x54\xb0\x49\xf7\xb0\x64\xe6\x70\xe4\xa0\x05\xb0\x2a\x17\x52\x31\x73\x0e\xb6\x52\x6c\x84\x66\x50\x57\x43\x2b\xc8\x99\xe4\xf7\x2c\x87\x95\x14\x1b\x64\xf3\x57\x62\xa5\x1f\xcc\xc9\x68\xce\x0c\xa8\x2d\xcb\xcc\xa1\x81\xad\xe4\xe6\x28\x49\x73\x5c\xaa\xfa\xe0\x28\x65\xe3\x46\x74\x4a\x12\x48\xa2\x6b\x7a\xeb\xc6\x18\x48\x02\xf3\x38\xba\x21\x3e\xf6\x61\x7c\x07\x74\x8a\xc1\x8b\xe6\x77\x31\x99\x4c\x29\x4c\xa3\xc0\xc7\x71\x02\x6e\xe8\x83\x17\x85\x34\x26\xe3\x05\x8d\xe2\x04\xbd\x73\x13\x20\xc9\x3b\x3b\xe0\x86\x77\x80\x7f\x9c\xc7\x38\x49\x20\x8a\x81\xcc\xe6\x01\xc1\x3e\xdc\xba\x71\xec\x86\x94\xe0\xc4\x01\x12\x7a\xc1\xc2\x27\xe1\xc4\x81\xf1\x82\x42\x18\x51\x14\x90\x19\xa1\xd8\x07\x1a\x39\x76\xd3\xe7\xd3\x20\xba\x86\x19\x8e\xbd\xa9\x1b\x52\x77\x4c\x02\x42\xef\xec\x7e\xd7\x84\x86\x66\xaf\xeb\x28\x46\x2e\xcc\xdd\x98\x12\x6f\x11\xb8\x31\xcc\x17\xf1\x3c\x4a\x30\x98\xb4\x7c\x92\x78\x81\x4b\x66\xd8\x1f\x02\x09\x21\x8c\x00\xdf\xe0\x90\x42\x32\x75\x83\xe0\x34\x4b\x14\xdd\x86\x38\x36\xa1\xb7\x53\x84\x31\x86\x80\xb8\xe3\x00\x9b\x8d\x6c\x92\x3e\x89\xb1\x47\x4d\x36\xc7\x57\x1e\xf1\x71\x48\xdd\xc0\x41\xc9\x1c\x7b\xc4\x0d\x1c\xc0\x3f\xe2\xd9\x3c\x70\xe3\x3b\xa7\x59\x33\xc1\x7f\x5b\xe0\x90\x12\x37\x00\xdf\x9d\xb9\x13\x9c\xc0\xfb\xff\x50\x91\x79\x1c\x79\x8b\x18\xcf\x4c\xc8\xd1\x35\x24\x8b\x71\x42\x09\x5d\x50\x0c\x93\x28\xf2\x6d\x9d\x13\x1c\xdf\x10\x0f\x27\xdf\x43\x10\x25\xb6\x58\x8b\x04\x3b\xc8\x77\xa9\x6b\x37\x9e\xc7\xd1\x35\xa1\xc9\xf7\xe6\xf5\x78\x91\x10\x5b\x33\x12\x52\x1c\xc7\x8b\x39\x25\x51\x78\x05\xd3\xe8\x16\xdf\xe0\x18\x3c\x77\x91\x60\xdf\x16\x37\x0a\x4d\xaa\x88\x4e\x71\x14\xdf\x99\x45\x4d\x0d\x6c\xed\x1d\xb8\x9d\x62\x3a\xc5\xb1\xa9\xa7\xad\x94\x6b\x4a\x90\xd0\x98\x78\xb4\x7d\x59\x14\x03\x8d\x62\x8a\x8e\x39\x42\x88\x27\x01\x99\xe0\xd0\xc3\x66\x34\x32\xab\xdc\x92\x04\x5f\x81\x1b\x93\xc4\x5c\x40\xec\xb6\x70\xeb\xde\x41\xb4\xb0\x29\x9b\x5b\xb4\x48\x30\xb2\x2f\x5b\x07\xd6\xb1\x37\x12\xc8\x35\xb8\xfe\x0d\x31\x61\x37\x17\xcf\xa3\x24\x21\xcd\x31\xb1\x25\xf3\xa6\x4d\xb9\x87\x08\xfd\xb9\xe3\x0f\x42\x6b\xae\x8b\xdd\x72\x98\x89\xcd\xa8\xe4\xcb\xd1\xf6\xa7\xa3\xe0\x3c\x1b\xe9\x5d\x76\x3e\x7e\xfc\xf6\xbb\x0f\x1f\x7f\xeb\xc0\x37\xdb\x9f\xbe\x01\xaf\xc5\x0f\x34\x17\xb2\x06\x5f\x6b\x8e\x57\xcf\x81\x71\x99\x7e\x66\x30\xe3\xff\x62\x32\xad\xf6\x08\xcd\x1f\x31\x00\x5c\x41\xc1\x24\x5b\xee\x61\x2d\xd3\x4a\xb3\xdc\x81\x95\x64\x96\x5f\x59\x91\xca\xb5\x21\xac\x80\xb4\xda\x1b\x76\x28\x51\x81\x58\x1a\x5d\x30\x3c\x4d\xad\x24\x98\x2b\x4f\x19\x64\xc8\x9b\x2a\x25\x32\x9e\x1a\x89\x3a\x85\xec\x8a\x97\x4c\xd5\xd2\xfd\x2e\x69\x66\xbc\xbb\xb2\x9b\xe4\x2c\x2d\xa1\x21\x73\xf2\x14\x68\x92\x19\xe6\x66\xb5\x06\xf2\x2a\x2b\x77\xb9\x89\xe1\x30\x5c\xf2\x0d\x6f\x76\x30\xd3\x1b\x1d\xd6\xc2\xf0\xd4\xb1\x71\x3a\x60\x75\xd4\xfc\x66\x36\xad\xed\x6e\x59\x72\x55\x38\xf0\x88\x73\xe6\x80\x32\x6f\x66\xac\x32\xb3\x1a\x3d\x50\xac\x2c\xcd\x0a\xbc\x06\x77\x3b\x3a\xa7\x56\x19\xd1\x08\x72\x53\x22\xbb\xef\x43\x21\x36\xa7\x99\x70\x05\xab\x9d\xac\xb8\x2a\x6a\xc2\xe7\x02\x94\xb0\x3b\xfe\x83\x65\xda\xbc\xf3\x92\x7c\x7f\x42\x88\x3e\x97\xe1\x46\x01\x1b\xa1\xe3\xaa\x05\xf7\xc3\x90\x2a\xd2\xb2\x34\xaa\x52\x17\x8c\xe5\xa6\xbc\x69\x2b\x1d\x69\xb6\x37\x2d\x94\xe6\x69\x09\xdb\xc3\x11\x7a\x92\xa6\xd5\x0a\x7c\x5e\x2a\x1a\x01\x70\xe0\x96\xd0\xa9\x79\x64\x1b\x6e\xdb\x67\xcf\x3c\xca\x3f\x90\xd0\x77\xce\xe8\x42\x4b\x0c\x0e\xe4\x83\x23\xf9\xec\x93\xfc\xb3\x12\xe0\xb4\xf9\x0f\xe7\xf9\x1f\xfa\x10\x46\x21\x09\xaf\x63\x12\x4e\x2c\x44\x5f\x12\x01\x77\x41\xa7\x06\xf6\x16\xd4\x4f\x55\xef\xb9\x02\x58\x49\x71\x1e\x29\x7e\x60\xd9\x0b\x90\x74\x43\x70\x3d\x43\x5a\x93\xc6\x91\x98\x06\x8e\x27\x18\x74\x1e\x31\x78\x1d\x47\x33\xe7\x40\xc0\xe8\x40\xda\x10\xd7\xab\x98\x52\xc3\xc9\x1d\x31\xa8\xad\x39\x79\x8c\xc5\xc7\x6e\x40\xc2\x89\x61\xfd\xc9\xc5\x7d\x53\x70\xfb\x79\x3d\x62\x52\x1a\x06\x9d\x21\x61\x6b\xb4\x7f\x1a\xfe\xde\x01\x3f\xbd\x67\xe0\x15\xac\x62\x7b\xf8\x53\x9e\xde\xb3\xbf\x66\xf6\x8f\x61\xc5\xf4\x5f\xd0\xff\x79\x4f\x0e\x75\x4f\xde\xb5\x21\x3f\xe9\x85\x11\x7c\x79\x4b\x7e\x26\x82\x57\x36\xe4\xcf\x83\x40\xf0\xa6\x96\x1c\xce\xb5\xe4\x08\x5e\xdd\x94\xc3\x69\x53\xde\x47\x73\x7b\x60\x1b\x7a\x73\x73\x0b\x4f\x9a\x5b\xf4\xa6\xe6\xf6\x05\xb8\xc5\x18\x7d\x41\x73\xdb\x64\xf9\x72\x77\x8b\x5e\xd5\xdd\xc2\x6b\xba\x5b\xf4\x33\xdd\x2d\x7c\x59\x77\x8b\xce\x77\xb7\xf0\x86\xee\x16\x3d\xeb\x6e\xa1\x43\x77\x8b\x9a\xee\x16\x7e\xb5\xdd\xad\x2a\xb8\x14\xbb\xd1\x5a\x6c\xd5\x4e\xf3\xf2\x1c\xdc\x9f\x5e\xd2\x99\xf0\x87\x95\x4c\xb3\x74\x6c\xcb\x72\xd8\x55\x39\x93\x30\x4e\x7c\x68\xba\xb3\x23\x8b\x72\x58\xb2\x52\x3c\x0c\xd1\x73\x79\xf8\x9d\x03\xb7\xee\x0f\xee\x9d\x3b\x73\x21\xb1\xa1\xfe\x62\x82\x00\x27\x82\x80\x3a\x0a\x42\x67\x8b\xe6\xa9\x22\xbc\xc1\xa5\xe9\x57\x12\x7a\xd5\x84\xb7\x8a\xc2\x0b\x3e\x8d\x79\xfd\x78\xf0\xd2\xda\xce\x3b\x6b\xd8\x40\xdb\xb0\x31\x71\xbc\xd1\xb3\x81\xb3\x9e\x8d\x59\xf0\xbf\x64\xdb\x40\xcb\xb6\x41\xbd\x28\xdb\x61\x1a\xfa\x1f\x28\xdb\x2b\x6c\x1b\xd4\x8f\xb0\x1d\x1a\x7e\xd4\x5d\xd8\x8e\xb6\x0d\xea\x2a\x6c\xa7\xb6\x0d\xea\x28\x6c\xfd\xda\x36\xd0\x08\x1b\xea\x26\x6c\x07\x65\x41\xbc\xd2\x4c\x56\x69\x39\xca\xc4\x66\x23\xaa\x51\x0d\xa4\xe1\x5a\x1c\x58\xd1\x56\x10\xfb\x29\x37\xb7\x88\xa9\xf1\xdc\x3c\x7e\xf5\x37\x03\x23\x56\x65\x22\xe7\xd5\xfa\xb8\x88\xd9\xe9\x62\xf5\x5f\xac\xfe\x8b\xd5\x7f\xb1\xfa\x2f\x56\xff\xd7\x6d\xf5\xf7\xfa\x59\x68\xbf\xdb\x7d\xe6\xa3\x52\xe9\xd5\xb9\x8f\x41\xad\xd1\xfe\x3d\xae\x3f\xc0\xfe\xef\x66\xfd\x01\x1a\xa0\x19\xa1\x10\xd4\x1f\x80\xcc\x9f\xbd\xd8\xf8\x03\xd4\xd1\xc7\x1f\xa0\x1e\x8c\xfc\x01\xea\xc1\xc9\x1f\xa0\xee\x56\xfe\x00\xf5\xe2\xe5\x0f\xd0\x8b\x66\xbe\xb9\x73\xdd\xec\xfc\x01\xea\xe8\xe7\xdb\x10\xba\x39\xfa\x03\xd4\xdd\xd2\x1f\xa0\x37\x79\xfa\x03\xd4\x8f\xa9\x3f\x40\xfd\xb8\xfa\x03\xd4\x8b\xad\x3f\x40\x1d\x7d\xfd\x41\xcf\x06\x50\x8d\xb5\xcd\x4a\xcb\x34\x63\xf2\x65\xf0\x3d\x5e\xf1\xcb\xc1\x0f\xb5\xd1\xd7\xd3\xf7\x97\x5d\xb9\xd7\x07\xf6\xfa\xa0\x5e\x0f\xd0\xeb\x87\x79\x2f\x23\xaf\xeb\xf7\x97\x5d\x79\xd7\xf9\xfb\xcb\x1e\x60\xf7\x36\xd6\xf5\x84\xba\x9e\x48\xd7\x0f\xe8\xba\x72\xae\x17\xcc\x1d\xff\x39\xf0\x9f\x23\xb5\x57\xe7\xfe\x6b\xb0\x1e\xe8\x1d\x6a\x17\x3f\xe1\xe2\x27\x5c\xfc\x84\x8b\x9f\x70\xf1\x13\xbe\x2a\x3f\xa1\x07\xd1\xf9\x77\x00\x00\x00\xff\xff\x3f\x65\x33\xed\xb4\x2e\x00\x00")

func CreditsBytes() ([]byte, error) {
	return bindataRead(
		_Credits,
		"../../CREDITS",
	)
}

func Credits() (*asset, error) {
	bytes, err := CreditsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../CREDITS", size: 11956, mode: os.FileMode(420), modTime: time.Unix(1560316799, 0)}
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
	"../../CREDITS": Credits,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"CREDITS": &bintree{Credits, map[string]*bintree{}},
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
