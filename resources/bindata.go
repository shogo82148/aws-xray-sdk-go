// Code generated by go-bindata.
// sources:
// resources/AWSWhitelist.json
// resources/DefaultSamplingRules.json
// resources/ExampleSamplingRules.json
// DO NOT EDIT!

package resources

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

var _resourcesAwswhitelistJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xcd\x72\xe2\x38\x10\xbe\xf3\x14\x2e\x9f\x73\xdb\x5b\x6e\x19\xf2\x53\xd4\x92\x09\x09\xcc\xe6\xb0\xb5\x45\xc9\x72\xe3\x68\x62\x4b\x8e\xd4\x62\xa0\xb6\xf2\xee\x5b\x92\x0c\x01\x63\x8c\xb0\x03\x21\x99\x3d\x4c\x4d\xb0\xda\xea\xaf\xbb\xbf\x6e\x75\xcb\xff\x76\x82\x20\x54\x20\xa7\x8c\x82\x0a\xcf\x03\xf3\x3b\x08\xc2\x78\xce\x49\x26\xe2\x68\xf9\x24\x08\x42\x91\x83\x24\xc8\x04\x57\x2b\x4f\x83\x20\xfc\x46\x90\x3e\xdd\x00\xf6\x10\xb2\xb5\x95\x20\x08\x25\xbc\x68\x50\x38\x8e\x41\x51\xc9\x72\x14\x52\x95\x44\x82\x20\x7c\x70\x42\xe6\xfd\xcd\xd5\x20\x08\x33\x92\x87\xe7\x01\x4a\x0d\x67\xe5\xa5\x04\x70\xfc\x0c\x73\xb5\x6d\x5d\x02\x27\x19\x8c\x51\x84\xe7\x41\x88\x24\x4a\x61\x6c\x1e\xa8\x70\x4d\xf0\x75\xe5\xd7\xeb\xd9\x3a\x7e\x95\x0b\xae\x60\x9c\x13\x49\x32\x40\xb0\xf8\xff\x5e\xc7\xdf\x15\x5c\xe9\x0c\xe2\x2e\xc9\x09\x65\x38\x5f\xdd\xfc\x9f\x4e\xc5\xc6\xce\x65\x8f\x92\x21\xfc\xef\xb4\x85\xd3\xd6\x61\x84\xc6\xb0\xae\x48\x53\xa0\x86\x72\xb7\x80\x92\x51\xb5\xdb\xb3\x5d\x09\x04\x61\x64\x40\x6f\x73\x6b\x1d\xaa\x9b\x54\x44\x24\x1d\x02\x15\x3c\x26\x72\xde\xe3\x31\xcc\x40\x95\xb1\xf5\x05\xdd\x2d\x34\x90\x62\xca\x14\x13\x1c\xe2\xd1\x93\x14\x3a\x79\xca\x35\x96\x85\x2c\xd0\xef\x24\x83\xdd\x96\x5d\x42\x0a\xf5\x7c\xa9\x33\xac\x5a\xd1\xc9\x84\xcd\x19\xd7\x38\x6c\x7b\xb9\xd1\x24\x55\x74\x14\x5d\x3b\x8a\xe2\x2e\x3f\x33\x85\xc0\xf1\x01\x48\x5c\xc1\xad\x9f\xce\xc3\x57\xb3\x5c\x82\x32\x34\xf3\xa3\xd6\x31\xaa\x5b\x9f\x29\xb4\xda\xcb\x85\xc9\xcb\xf2\xab\x19\x4d\xb5\x62\x53\x18\x22\x91\xf8\x66\x45\x39\x07\x59\xc6\xd0\xc3\xb4\xda\x22\xba\xdc\xbd\xb2\x84\xa6\x4c\x61\x5d\x0d\xa5\x42\xf3\xad\x02\x15\x45\xd4\xc9\x6f\x2f\xa2\x95\xce\x1c\xe8\xc6\x1c\x3a\xf1\x9c\xbf\xd7\x20\xe7\x4d\xec\xba\x40\x94\x2c\xd2\x08\x6a\x24\x6e\x60\xa3\xa6\xd6\xa7\x8e\xad\xd6\xdb\x09\xd5\x20\xd1\x86\x94\x70\xbb\xeb\xb5\x90\xbf\x88\xdc\xd0\x38\x04\xe3\x99\xd3\x49\x4f\x83\xf7\x2b\xb8\x1d\x92\x0c\xf8\xc6\xab\xbb\xbc\x5d\x5e\x10\x68\x8e\x72\xbb\x95\x3a\x70\x9e\x74\x6d\xfe\x57\xb0\x87\x43\xdc\x2d\xd7\x86\xea\xd8\xfd\xc8\x63\xf2\x65\x9b\x00\x67\x5c\xe3\x83\x79\xc9\xcf\x4b\x98\x30\xce\xdc\xa4\x72\xb6\xbb\xbf\x73\x6a\x0f\xd3\xbf\x75\x56\xff\x2f\x6c\x0d\xd5\x8b\xf2\x99\xab\x2e\xe2\x78\x00\x32\x63\x8e\xfd\x0d\x1c\xd2\x27\x11\xa4\x65\xc4\xf7\x1a\x34\xfc\x90\xa9\x47\x2b\xfd\x44\x78\x02\xb7\xa0\x14\x49\xe0\x2f\xa6\x58\xc4\x52\x13\xec\x06\x48\x96\x4a\x4b\x60\xde\x76\x1d\xb1\x0c\x84\xf6\xc8\x80\x2d\xa8\xec\x44\xd5\x0a\x5a\xbb\x94\xb8\x26\x2c\x85\xd8\x77\x3e\xb1\x4a\xdb\xd5\xe0\xca\xb8\xee\x33\x49\x14\x2e\x7c\x37\x9f\xed\x56\xf5\x59\x62\xe4\x40\x37\x8e\x91\xbf\x77\x6e\x00\xad\xf0\x4a\x54\x3f\xce\x3b\x2b\x20\xbc\x81\x1b\xad\x4d\x11\x57\x9d\xc6\x76\xe1\xee\x17\x07\x79\xf1\x38\xbc\xa0\xb6\x61\xee\xc5\x2d\x0d\xf3\x8f\x87\x19\x5c\x2e\x81\xc4\x7d\x40\x04\x39\x14\x5a\x52\x47\x83\x8f\x8c\xcb\x62\x1b\x8f\xb0\x18\xfc\x2d\xf1\x9a\xa8\x0c\x24\x4c\xd8\xac\xed\x6c\xf5\x86\xfb\xa0\xa3\xd5\x8b\x51\xd3\x74\xb4\x92\xc9\x31\xd2\xfc\x01\x28\xb0\x69\x9b\x82\xbb\xcc\x4d\x37\xab\x96\xb2\xe6\x96\xcc\xbe\xeb\x2c\x02\x79\x37\x29\x74\x6c\x8a\xb8\xe7\xf5\xfb\xf8\x9f\xd2\x25\x81\x47\xc2\xd0\x2c\xb9\xbe\xca\xa7\x8d\xae\x65\xce\xd2\x88\x83\x12\x27\x73\x5a\x9a\x51\xe7\x01\x32\x31\x85\x76\xed\x99\x3f\x81\x86\xc0\xe3\x16\xec\xb9\x84\x94\xcc\x17\xb1\xf1\x69\x07\xcf\xaa\x34\xf8\x84\x6c\xeb\x41\x66\x85\xde\xeb\x4a\x7a\x11\x39\xb2\xd0\x76\x80\xeb\xe9\xc2\xa0\x9e\x47\xbb\xb0\x12\x9d\x83\x76\x38\x1e\x61\xb8\xe2\x28\xd9\xa9\x24\x4e\x83\xcc\x2f\x3a\xb4\x83\xc2\x9f\x58\x1d\x95\xe8\xcb\xa3\xb9\xa6\x14\x94\x9a\xe8\x72\x93\xf3\xce\x88\xd4\x52\x4f\xb3\x62\x34\x3c\x5a\x1b\xe9\x41\xc1\xe3\x94\x00\xff\xd4\xef\x94\xff\x5a\x9f\xc5\x53\x92\x45\x31\xf1\x19\xc7\x7b\x7c\x2a\x9e\x1b\xd5\xdf\x6b\xcd\xed\x45\x48\x55\xc7\x6b\x76\xa5\x56\xdd\x68\x9e\x6f\x5e\x8e\x89\xa4\xea\xf1\xbd\x26\x29\x9b\x30\x90\x6d\x07\xa2\x02\xd8\x95\x94\x42\x6e\x5c\x4c\x21\x41\xad\xba\x22\xf6\x98\x28\x9d\x6f\x2e\xd4\x9c\xd3\xd6\x0e\x6a\x67\x92\x43\xbd\xd7\x6d\xcc\x1f\x3e\xd1\x77\x73\xfb\x37\x4d\x9f\x01\x1b\x7d\xaf\xad\x7c\xd5\xae\x4c\x49\x6a\x3b\xd0\xd0\xd0\x3e\xac\xe7\x7d\x64\x77\xb1\xa4\xdf\xb7\x4a\xb8\xa9\xf6\x33\x5b\x70\x03\x78\x17\xfd\x04\xba\x15\x7e\x1d\x31\xfe\x84\x79\xcb\xca\xf6\xe1\xf6\xbb\x08\xfe\x2e\x2e\xd8\x3f\xf7\x8b\xcb\x26\x22\x9f\x4b\x95\x71\xa7\x37\xb7\x1e\x98\x27\xee\x8d\x5a\x78\xce\xc0\x03\x77\x53\xb1\x53\x32\x16\xd6\x8f\xcd\xfa\x97\x3e\x53\xe8\x7c\x59\x11\x07\x1f\x4b\xab\x5f\x7e\x67\x4b\x8b\x38\x35\xb6\xf0\x0b\x33\xad\x2b\x38\xda\xef\x67\x07\x0d\x40\x1b\x8a\x0d\xf4\xef\x7d\x74\x0c\x74\x91\x61\x7d\x91\x24\x8c\x27\x27\xcf\xc2\x7a\x2b\x06\x22\x65\xb4\xd1\x97\xa8\xe2\xcd\xaf\x12\xce\x11\xf9\x6c\xe1\x5c\xef\xc1\x3b\xe6\xdf\x6b\xe7\xbf\x00\x00\x00\xff\xff\x2f\xee\xa4\x96\x86\x2a\x00\x00")

func resourcesAwswhitelistJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourcesAwswhitelistJson,
		"resources/AWSWhitelist.json",
	)
}

func resourcesAwswhitelistJson() (*asset, error) {
	bytes, err := resourcesAwswhitelistJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/AWSWhitelist.json", size: 10886, mode: os.FileMode(420), modTime: time.Unix(1504040357, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesDefaultsamplingrulesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x2a\x4b\x2d\x2a\xce\xcc\xcf\x53\xb2\x52\x30\xd4\x01\xf1\x53\x52\xd3\x12\x4b\x73\x4a\x94\xac\x14\x40\xd2\x0a\x0a\x4a\x69\x99\x15\xa9\x29\xf1\x25\x89\x45\xe9\xa9\x25\x30\x55\x0a\x0a\x4a\x45\x89\x25\xa9\x4a\x56\x0a\x06\x7a\x06\xa6\x5c\x0a\x0a\xb5\x60\xbd\x45\xa5\x39\xa9\xc5\x4a\x56\x0a\xd1\x5c\x0a\x0a\xb1\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\xc9\x98\x17\x60\x61\x00\x00\x00")

func resourcesDefaultsamplingrulesJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourcesDefaultsamplingrulesJson,
		"resources/DefaultSamplingRules.json",
	)
}

func resourcesDefaultsamplingrulesJson() (*asset, error) {
	bytes, err := resourcesDefaultsamplingrulesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/DefaultSamplingRules.json", size: 97, mode: os.FileMode(420), modTime: time.Unix(1501190086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesExamplesamplingrulesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x8e\xd4\x30\x10\x84\xef\x79\x8a\x92\x2f\x0b\xab\x6c\x98\x3d\x70\xc9\x8d\x03\x2f\x80\xb8\x21\x14\xf5\xc4\x9d\x49\x0b\xc7\xce\xda\xed\xd9\x45\x68\xde\x1d\xd9\xd9\x61\x7e\x38\xba\xab\xda\x5d\x5f\xfd\x69\x00\x73\xe4\x98\x24\x78\xd3\xe3\xb9\x2d\x6f\xcb\x13\x65\xa7\xa6\x47\x91\xeb\x20\x8d\x51\x56\xdd\x4c\xe6\x0b\xde\x1d\x88\xd9\x71\x0b\x4a\x10\x3f\xba\x6c\xd9\x62\xcf\x2e\xbc\xb6\x90\x84\xc8\x2f\x59\x22\x5b\x88\x07\xf9\xdf\x48\xb4\xac\x4e\xfc\xa1\x2e\x25\x4c\xe2\xb8\xc3\x87\xc4\xf1\x28\x23\x0f\x9e\x16\x6e\x31\xab\xae\xc3\xc2\x3a\x07\xdb\x82\xbc\x45\x8e\x6e\x58\x49\x67\x50\x64\x4c\xf2\xc6\x16\x1a\xf0\xf0\xf8\x80\x29\x44\xe8\x5c\xee\x64\xc7\xdd\x47\xd3\x6e\x51\xab\x67\x50\x8a\x07\xd6\x33\x10\x60\x22\x29\x9b\x1e\xbb\x6e\xf7\xb9\x01\x4e\x15\xb3\xe6\x30\x3d\x7e\x54\xcb\x86\xfa\x3f\xec\xd7\xb7\x92\x9b\x51\x52\x3c\xed\x29\xb1\xad\x17\x37\xd0\x0e\xdf\x2a\x4c\x49\xc7\x47\x72\x99\x74\x03\x16\xfb\x14\xa2\xe5\xd8\x42\x67\xbe\xa9\x0b\xaf\xe2\x1c\xf6\x8c\x5c\xbe\x92\x09\x3e\x78\xc6\x42\x3a\xce\xd5\x2b\x7e\x0c\x4b\xad\x89\x5f\x32\x27\xed\xf0\xbd\x50\x4a\x02\x6d\xfb\x1b\x38\x63\x9c\x79\xfc\x15\xb2\x62\xa5\x03\x77\xef\xfc\x80\x11\x5b\x62\x3f\x5f\x06\xd7\x15\x17\xe9\xf1\x22\x5d\xf5\x7d\xa7\x9c\x8b\x2f\xe3\x4f\xe7\x53\x17\xf9\xbe\xe6\xdd\x3f\xe5\xb6\x69\xe0\xd4\x00\x3f\x9b\xd3\xdf\x00\x00\x00\xff\xff\xc1\xd0\x29\x3f\x69\x02\x00\x00")

func resourcesExamplesamplingrulesJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourcesExamplesamplingrulesJson,
		"resources/ExampleSamplingRules.json",
	)
}

func resourcesExamplesamplingrulesJson() (*asset, error) {
	bytes, err := resourcesExamplesamplingrulesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/ExampleSamplingRules.json", size: 617, mode: os.FileMode(420), modTime: time.Unix(1503352180, 0)}
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
	"resources/AWSWhitelist.json": resourcesAwswhitelistJson,
	"resources/DefaultSamplingRules.json": resourcesDefaultsamplingrulesJson,
	"resources/ExampleSamplingRules.json": resourcesExamplesamplingrulesJson,
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
	"resources": &bintree{nil, map[string]*bintree{
		"AWSWhitelist.json": &bintree{resourcesAwswhitelistJson, map[string]*bintree{}},
		"DefaultSamplingRules.json": &bintree{resourcesDefaultsamplingrulesJson, map[string]*bintree{}},
		"ExampleSamplingRules.json": &bintree{resourcesExamplesamplingrulesJson, map[string]*bintree{}},
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

