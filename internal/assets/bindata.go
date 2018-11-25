// Code generated by go-bindata. DO NOT EDIT.
// sources:
// client.sh (698B)
// controller.sh (983B)
// node.sh (565B)
// registry.yaml (1.163kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _clientSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x31\x6f\x1b\x31\x0c\x85\x77\xfd\x8a\x57\xb7\xab\x7d\xb1\xe3\x18\x29\x9c\x76\x4b\x01\x4f\xcd\xd0\xa1\xab\x4e\x47\xfb\x88\xd3\x89\x82\x48\x1b\x2e\x82\xfc\xf7\x42\x97\xa6\x70\xda\x4e\xcd\x42\xe0\x01\x4f\x1f\x1f\x45\xbe\x7f\xd7\xb4\x9c\x9a\xd6\x6b\xef\x9c\x92\x61\x4e\xc7\xb3\x20\x73\xa6\xbd\xe7\xe8\x9c\x15\x9f\x31\x33\x39\x86\x1e\x4d\x2f\x6a\x4d\x2b\x62\x1a\x0a\x67\x9b\x77\x92\x68\x86\xfb\xef\xbb\x6f\xce\xed\x1e\x3e\x2d\x3f\xae\x16\xcb\xcd\xed\xe2\xe6\x6a\x71\xed\x38\xc3\x77\x5d\xa9\x05\x1f\x1e\x77\x0f\x4f\xcd\x6a\x8d\x8e\x4e\xa0\xa4\x6b\x57\x49\xc9\x8f\x14\x2c\x42\xc9\xe6\x2f\x1a\x21\x32\x25\x73\x2e\x78\xc3\xe7\x86\x2c\x34\x2d\x97\x6e\x2a\x8b\x20\x69\x8f\xbb\xbb\xfb\xaf\x5f\x5c\x91\xa3\x51\x01\xff\x62\x6f\x9d\xcb\x45\x4c\x82\x44\x0c\x54\x12\x45\x3c\x3a\x20\x53\x51\x56\xdb\x3a\x60\xa4\x72\x20\x64\x6f\xbd\x56\xa9\xc1\x27\x18\x8f\x84\xcd\x55\xd5\x3c\x66\x29\x86\x24\x89\xaa\xa4\xf3\x24\x7d\x8c\x5b\xf7\x74\xc1\xee\xe8\xc4\x81\x26\xf6\x6b\xc2\x2b\x13\x17\x0a\x36\x99\x38\x19\x95\xbd\x0f\x84\x59\x1d\x7b\x56\xd9\xa1\xa7\x30\x20\x72\x1a\x2e\x1a\xff\xd5\xa9\x3d\x64\x0c\xb7\x3a\x7a\xad\x73\x56\x56\x94\xe0\x23\xbc\x62\xb3\xbe\x59\x5e\xd7\xb7\xd9\xab\xf2\x69\x0a\x9c\x88\x0f\x7d\x2b\x05\x17\x3b\x58\xbe\x98\x57\xcf\x8e\xb3\xa1\x97\x0c\xa5\xb8\xff\xb3\xf3\xef\x89\x9f\x3f\xe0\x1f\x41\x92\x74\xf4\x7f\x31\x56\x6f\x88\x51\x57\xad\x3f\xd4\x68\xac\x87\x52\x48\xcd\x17\x43\x3d\x06\xf7\x33\x00\x00\xff\xff\x86\x8f\xe8\xe4\xba\x02\x00\x00"

func clientShBytes() ([]byte, error) {
	return bindataRead(
		_clientSh,
		"client.sh",
	)
}

func clientSh() (*asset, error) {
	bytes, err := clientShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1d, 0xff, 0x78, 0x58, 0xc, 0x94, 0x36, 0xb8, 0x14, 0x8, 0xfb, 0xf7, 0xb8, 0xb8, 0x84, 0xd3, 0x64, 0x30, 0xf9, 0x18, 0xca, 0x68, 0x9, 0xfb, 0xe6, 0x35, 0x28, 0x54, 0x45, 0xdf, 0xe0, 0xcd}}
	return a, nil
}

var _controllerSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6e\xdb\x30\x0c\xbd\xeb\x2b\x38\xad\x87\xed\x20\x29\xce\x8a\x6d\x30\xda\x02\x5d\x90\x0e\x45\x81\x14\x68\xb6\xde\x15\x8b\x89\x85\xc8\x92\x20\xd1\x5e\x8b\x61\xff\x3e\xc8\xf1\x1a\x74\xed\xa5\xba\xd8\x20\xf9\x1e\xf9\xc8\xf7\xfe\x9d\xda\x58\xaf\x36\x3a\xb7\x8c\x65\x24\x10\xd8\x3f\x04\x88\x36\xe2\x56\x5b\xc7\xd8\xfd\xf2\x6e\x7d\x7d\xbb\x3a\xe7\x43\x25\xab\xb9\x9c\x73\xc6\x1a\x4d\x70\xa1\xa8\x8b\x6a\xdf\x6f\x50\x9b\x4e\x36\xc1\x6f\xe1\xec\x6c\x79\x7b\xc5\x74\xb4\xf7\x98\xb2\x0d\xbe\x86\x7f\xe9\xfd\xd7\x2c\x6d\x50\x43\xa5\x5d\x6c\xf5\x27\xb6\xb7\xde\xd4\x70\xed\x2d\x2d\x82\xdf\xda\x5d\x9f\x34\xd9\xe0\xd9\x26\x04\xca\x94\x74\xfc\x11\xf6\xe8\x73\xcd\x04\x50\xf9\xab\x81\xcf\xc6\x27\x67\xff\x3d\xce\x00\x88\x5c\x0d\x7c\x7e\xda\xf2\xd2\x7c\xe9\x4d\x0c\xd6\x53\xcd\x00\xb4\x19\x30\x91\xcd\x78\x69\x4c\xc2\x9c\x6b\x38\xf9\xd0\xa2\x36\x20\x2a\x50\x6d\xc8\xa4\x6c\xfc\xc8\x7c\x30\x78\x87\x3b\x5b\x3a\x97\x31\x0a\xb2\x4c\xee\x90\x96\x0f\x94\xf4\x65\xda\xe5\x12\x03\x28\x95\xc2\xc6\xd7\x69\x84\x10\x6f\x10\xbf\x70\x7d\x26\x4c\xcf\xf5\x7b\xa4\x5f\x21\xed\xad\xdf\x95\x7e\x31\x98\x75\xbf\xf1\x48\x35\xf0\x6a\x26\x4f\xe7\x72\x26\x67\xaa\xfa\xcc\x59\xa1\x4e\x1e\x09\xf3\x53\x37\x7e\xf2\x7b\xba\xd4\x1f\xce\x9a\x03\xf9\x4a\x77\x58\x03\x1f\x6c\xa2\x5e\x17\xcc\xb8\xa0\x35\xa6\x01\xd3\x02\x13\xad\x2f\x57\xe3\x8e\x79\x35\xff\x52\xb8\x65\xc5\x5f\x55\xe1\x90\xc6\x0b\xdb\xdd\x51\xcc\x06\x49\x57\x93\x96\x9b\x43\xcd\x73\x2d\x09\x73\x70\x43\x89\xd5\xa0\x52\xef\x55\x7e\xcc\x84\x9d\x51\x87\x04\x4e\xdf\x91\x98\x15\xe3\x60\xd3\x06\xe0\xf8\x10\x43\x22\xb8\xf9\xf9\x6d\xb9\xb8\x5d\x5d\x5d\x7f\x3f\x57\x48\x8d\x3a\x2a\x56\xda\x74\xd6\x8f\x30\x0e\x17\x63\x32\xa6\xb0\xb5\x0e\xa5\x51\x65\xbc\xdc\xb2\x69\xf5\x60\xbd\x25\x10\xe2\x30\xfb\xf9\x0b\xc7\xb2\xb7\x34\x1b\x49\x1b\x72\x40\xda\x7a\x1a\xad\x90\x41\x08\xed\xdc\xc1\x16\x29\x38\x94\x47\x64\xd9\x52\xa7\xcb\x15\xc4\x13\x52\xc7\xe8\x1e\x41\x6c\x27\xd3\x68\x63\x82\xcf\xf2\x51\x77\x8e\x35\x11\x4e\x8e\x63\x4c\x05\x23\xee\xb0\x77\xea\x22\xeb\x86\x23\xf6\x79\xea\x45\x90\xfd\x0d\x00\x00\xff\xff\xbe\xd8\x64\xca\xd7\x03\x00\x00"

func controllerShBytes() ([]byte, error) {
	return bindataRead(
		_controllerSh,
		"controller.sh",
	)
}

func controllerSh() (*asset, error) {
	bytes, err := controllerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0xae, 0xfd, 0x76, 0x94, 0x28, 0x26, 0xf0, 0xde, 0xc3, 0xc5, 0x17, 0xbb, 0x45, 0x7a, 0x6b, 0xf7, 0xe9, 0xd1, 0x35, 0xc3, 0xbb, 0xe9, 0x78, 0x85, 0xd4, 0xcf, 0x3a, 0x7f, 0xb5, 0x37, 0x4b}}
	return a, nil
}

var _nodeSh = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x51\xab\xd3\x30\x14\xc0\xf1\xf7\x7c\x8a\x63\xef\xe0\xde\x3d\xa4\xdd\xd8\x10\x09\x9b\x50\x46\x45\x65\xae\xb2\xcd\xbd\x4a\x96\x9e\xb6\xc7\x76\x49\x48\xd2\x32\x51\xbf\xbb\xb4\x53\x91\xa9\xe0\x79\x4c\xfe\x84\x9c\xdf\xc3\xb3\xe4\x4c\x3a\x39\x4b\x5f\x33\xe6\x31\x00\xc7\xee\x6a\xc0\x92\xc5\x52\x52\xcb\xd8\x26\xdf\x1d\xf7\xf9\x76\x9b\xed\x3f\xbe\xce\x0f\xc7\x5d\xfa\x2e\x5b\x47\x93\xa7\xda\xf8\xa0\xe5\x05\xe1\x2b\xa8\x2e\x00\x2f\xe7\xc0\x8b\x47\xfe\x38\xe5\xca\xe8\xe0\x4c\xdb\xa2\x8b\x18\x53\x32\xc0\xcb\x24\x5c\x6c\xd2\x74\x67\x94\xc5\x25\x56\x46\x97\xb0\x5a\x65\xf9\x2b\x26\x2d\x9d\xd0\x79\x32\x5a\xc0\xcf\xeb\xe6\x85\x8f\xc9\x24\xfd\x5c\xb6\xb6\x96\x0b\xd6\x90\x2e\x04\xbc\x35\xa4\x37\x46\x97\x54\x75\x4e\x06\x32\x9a\x05\xd3\xa0\x16\x10\xcd\xc6\x89\x67\x77\x13\xb1\x82\xbc\x32\x3d\xba\xcf\xc7\xa1\xfc\xa0\xbd\x2c\xf1\xd0\x90\xdd\xa4\x27\x74\x54\x92\x1a\xdf\x11\x10\x5c\x87\x77\x71\xfa\xfe\xcd\x01\x5d\x8f\xce\x0b\xc6\x61\xf2\xe5\x2f\x04\xdf\xe2\xd6\x28\xd9\x8a\xe7\xcb\xe5\x82\x69\x53\xe0\x1e\x2b\xf2\xe1\xf6\x39\xc1\x60\xdc\xa7\xc5\x90\x5d\x83\x93\xa9\xab\xfc\x70\x06\x30\x94\x9c\xac\x80\xc9\x53\x8d\xb2\x00\x3e\x87\x64\x90\x4c\xc8\x4e\xd9\x40\xf2\x83\x01\x3e\x19\xd2\xc0\x07\xcb\x92\xaa\xf5\x1f\x80\xec\xe1\x2e\x1c\x39\xd6\xff\xc0\x00\xce\x7f\x2d\x78\x2b\x79\x37\x7a\x70\xdf\x90\xe5\x4a\xf2\xfe\x37\x92\xff\x58\xf8\x7b\x00\x00\x00\xff\xff\xe8\xbc\x4b\x30\x35\x02\x00\x00"

func nodeShBytes() ([]byte, error) {
	return bindataRead(
		_nodeSh,
		"node.sh",
	)
}

func nodeSh() (*asset, error) {
	bytes, err := nodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x5c, 0x7c, 0x73, 0x74, 0xbf, 0x12, 0xb0, 0x21, 0xb, 0xa3, 0x42, 0x32, 0xe, 0xdf, 0xff, 0x80, 0xc8, 0x2d, 0x83, 0x47, 0xa6, 0x22, 0x2d, 0x7e, 0x94, 0x14, 0x29, 0xb, 0x5c, 0xf6, 0x78, 0xb4}}
	return a, nil
}

var _registryYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x53\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x10\xbd\xbb\x71\x36\xec\xa2\x5b\xd1\x78\x5d\x80\x65\x09\x6c\x63\x40\x4e\x81\xe2\x10\x9d\x50\x7d\x8d\xa2\x8d\x19\xc3\xfe\xfb\xa0\x36\x4e\xe5\x3a\xdb\x65\x18\x4f\xc6\xa3\xde\x23\xf9\x4c\xe6\x79\x9e\x49\xaf\xbe\x22\x05\xe5\xac\x00\xfc\xc1\x68\xe3\x67\x58\xf4\xcb\x23\xb2\x5c\x66\x4f\xca\x9e\x04\xac\xd0\x6b\x37\x18\xb4\x9c\x19\x64\x79\x92\x2c\x45\x06\x60\xa5\x41\x01\x84\x8f\x2a\x30\x0d\x67\x20\x78\xd9\xa2\x80\xa7\xee\x88\x79\x18\x02\xa3\xc9\x00\xb4\x3c\xa2\x0e\x91\x03\x20\xbd\x4f\x48\xc1\x63\x1b\x71\x42\xaf\x55\x2b\x83\x80\x65\x06\xc0\x68\xbc\x96\x8c\x2f\x8c\xb4\x66\x8c\x54\x6d\xae\x18\x91\x51\x35\x86\x75\x27\xac\x51\x63\xcb\x8e\x5e\x39\x11\xcd\xc9\x69\xbc\x8d\x9d\x92\x45\xc6\x70\xab\xdc\xc2\xc8\xc0\x48\x02\x6e\x6e\xce\x4f\x5b\x67\x59\x2a\x8b\x74\x29\x98\xcf\x07\x7f\x09\x65\xe4\x63\x82\x8b\x77\x97\x0c\x61\x70\x1d\xb5\x98\x34\x0d\xa0\x95\x51\x3c\x41\x00\x5a\xdf\x09\x58\x16\x85\x99\xa0\x06\x8d\xa3\xe1\x39\xb1\x51\x49\x86\xf0\x7b\x87\xe1\x5f\x34\xd0\xf6\xaf\xe4\x71\xb0\xaa\x7c\x58\xd7\x4d\xb5\x3f\x7c\x6a\x9a\xdd\xe1\x6e\xb5\xaa\x12\xa5\x5e\xea\x0e\x05\x88\x0f\x45\x51\xfc\x99\x59\x37\xdb\xea\xee\xa1\x3c\x7c\x5c\x7f\x2e\xeb\x7d\xdd\x94\x9b\x43\xb5\xdd\x36\xab\x75\x55\xde\x37\xdb\x6a\x3f\x17\x5c\xf4\x92\x16\x5a\x1d\x17\x33\x5b\x7b\xa7\x3b\x83\x1b\xd7\xd9\x74\xd2\xb1\xe4\xb3\xe9\x79\x60\x47\x98\x88\x9a\xf8\x7a\x27\xf9\xdb\xdf\x84\xbd\xa3\xa9\xe2\xe5\x5f\xef\x1c\xb1\x80\xc9\x88\x57\xd6\xfd\xa2\x43\x8e\x5d\xeb\xb4\x80\xe6\x7e\x97\xa5\x4d\xcf\x76\xe6\x5a\xb7\x68\x3c\x0f\x2b\x45\x02\x7e\xfe\xca\xde\xde\x64\x3f\xde\x60\x8d\xd4\xab\x16\xff\xdb\x01\xf2\xe0\x51\xc0\x17\x77\xc2\x38\x7c\x96\xb8\x73\x75\xdf\xfd\xd4\x21\x7b\xe6\x09\x78\x5f\x8c\xd8\x5b\x5b\xc2\xe4\x08\xa7\x5d\xfc\x0e\x00\x00\xff\xff\x76\x26\x7f\x45\x8b\x04\x00\x00"

func registryYamlBytes() ([]byte, error) {
	return bindataRead(
		_registryYaml,
		"registry.yaml",
	)
}

func registryYaml() (*asset, error) {
	bytes, err := registryYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "registry.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x58, 0xd5, 0xef, 0xb9, 0x10, 0x6c, 0xdd, 0xa6, 0x39, 0x30, 0xc1, 0xc3, 0x77, 0xa4, 0x99, 0x9b, 0xd, 0x24, 0x34, 0x26, 0x40, 0x6, 0xa7, 0xe3, 0xab, 0x2d, 0x45, 0x56, 0x64, 0x28, 0x6d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"client.sh": clientSh,

	"controller.sh": controllerSh,

	"node.sh": nodeSh,

	"registry.yaml": registryYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"client.sh":     &bintree{clientSh, map[string]*bintree{}},
	"controller.sh": &bintree{controllerSh, map[string]*bintree{}},
	"node.sh":       &bintree{nodeSh, map[string]*bintree{}},
	"registry.yaml": &bintree{registryYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}