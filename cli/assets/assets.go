// Code generated by go-bindata.
// sources:
// assets/unversioned/console.html
// assets/v1.0-alpha/console.html
// assets/v1.0/console.html
// DO NOT EDIT!

package assets

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

var _assetsUnversionedConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5d\xab\xe3\x36\x10\x7d\xdf\x5f\x21\x54\xca\xbe\xd4\x52\x76\xbb\xd0\xc5\xd7\xbe\x50\x5a\x4a\x4b\x5b\xd8\x87\x4d\x5f\x97\x89\x3c\xb6\x27\x2b\x4b\xae\x46\x49\xae\x7b\xf1\x7f\x2f\xfe\x4a\x9c\xa4\x50\x7a\xe9\x92\x17\xcd\x99\xf1\x9c\x39\x23\x71\x92\xd5\xb1\xb1\xc2\x82\xab\x72\x89\x2e\x39\xb0\x7c\x7c\x25\x44\x56\x23\x14\xc3\x41\x88\xcc\x92\xfb\x2c\x02\xda\x5c\x92\xf1\x4e\x8a\xd8\xb5\x98\x4b\x6a\xa0\x42\xdd\xba\x4a\x8a\x3a\x60\x99\xcb\x3a\xc6\x96\x53\xad\x39\xfa\x00\x15\xaa\xca\xfb\xca\x22\xb4\xc4\xca\xf8\x46\xd7\xc0\x87\x00\x49\x15\xa0\xad\xff\xb4\x09\xba\x8a\x1c\x6a\xe3\x1d\x7b\x8b\x1a\x98\x31\xb2\x2e\xe1\x38\x70\xa8\xb1\xad\x9e\xf9\xd9\x04\x6a\xe3\x14\x88\x13\xb9\xc2\x9f\xd4\xa7\x4f\xe8\x8e\x22\x17\xcf\x13\x2a\x04\xb4\xf4\xb3\xe7\x98\x8a\xe7\x67\x35\x9f\xfb\xfe\x9b\x55\xf6\x83\x0f\x31\x15\x72\x4a\x0f\x41\xdf\xcb\x73\xde\x58\xfa\x03\x03\x93\x77\x63\x83\x4b\xb8\xea\x51\x40\x84\xef\x5b\xda\x06\x3b\xd6\x5c\xc2\xfb\x9a\x75\xaf\x6b\x68\x3d\x93\x31\xc8\xfc\x2b\x76\xd3\xcc\x4b\xb4\xaa\x38\x04\xfb\x21\x60\x49\x4f\xa9\x90\x7a\x35\xec\xb4\xb3\xdf\x7d\x81\xa9\x90\xc6\xd2\x95\x8e\xed\xf6\x97\x1f\x17\x11\xc3\x79\xd5\x0f\x1d\xec\x2c\x7e\x44\x8b\x0d\xc6\x30\xf1\xde\x60\x7d\x3f\x15\xf7\x0f\xd3\xee\xf5\x65\xf9\x99\x5e\xde\x44\xb6\xf3\x45\xb7\x5c\x4e\xec\x2c\x4e\x67\xd5\x00\xb9\x1f\xbc\x8b\xe8\xe2\xf9\x62\x0a\xe2\xd6\x42\x97\x8a\xd7\xce\x3b\x7c\xfd\x30\xc3\xbe\x05\x43\xb1\x4b\xc5\x66\x41\x62\x00\xc7\x14\xc7\xb5\xcd\x59\xa1\xde\x6e\x58\x58\x72\x08\x61\x2a\xeb\xef\x88\x14\xd7\xfe\xf4\x0f\x6c\x3b\xeb\xcd\xe7\x7b\xba\x37\x2f\xa0\xcb\xf4\x2c\x72\x8a\x0a\x3a\x0a\x2a\x72\x69\x3d\x14\xe4\x2a\x39\x3f\xcc\x29\x61\x2c\x30\xe7\xb2\x85\x0a\x93\xa5\x40\x8c\x9f\xe7\x72\xae\x6b\xc8\x25\x35\x52\x55\xc7\x54\xbc\xd9\x6c\x8e\xf5\x32\xd2\x89\x8a\x58\x8f\xd8\xd7\x0f\xb7\x7a\x4a\x8b\x4f\x0b\x08\x96\x2a\x97\x50\xc4\x86\x53\x61\xd0\x45\x0c\x4b\xaa\xf4\x2e\x26\x25\x34\x64\xbb\x54\x30\x38\x4e\x18\x03\x95\x4b\x7a\x7f\xe0\x48\x65\x97\x98\x69\x77\xb7\x5f\x9f\xa5\x0c\xf7\xda\x82\x5b\xd4\xdc\x2a\x98\x79\x98\xfe\xc2\x54\xbc\xc5\xe6\xe1\x8c\x37\x10\x2a\x72\x49\xf4\x6d\x2a\x92\x6f\xd7\x19\xe3\xad\x0f\xa9\xf8\xea\xfd\xbb\xe1\x77\xc1\x57\x9c\xbf\x4d\xfb\x52\x4a\x2d\x1b\xd5\xc3\x14\xe7\xfd\xea\x82\x8e\xf3\xab\x5b\x1d\x97\xeb\x98\x35\xc9\x65\xe8\xd5\x1b\x91\x8f\xeb\x0f\x2e\x76\x36\xaa\xe2\x1a\x31\xde\x7a\x98\x29\xdc\x9e\x95\xb1\xfe\x50\x94\x16\x02\x8e\x0e\x06\x7b\x78\xd2\x96\x76\xac\x47\xf9\x70\x42\xf6\x0d\xea\x77\xea\x3b\xb5\xd1\x86\xaf\x61\xd5\x90\x53\x86\x59\xea\xff\x40\xfb\x22\xeb\x1c\xcc\x63\x74\xcf\xb3\xc5\xe8\x41\xfb\x48\x2e\x4c\x0d\x81\x31\xe6\x72\xfb\xf1\xa7\xe4\xbd\xbc\xb6\x54\xc1\xc1\xfc\xff\xe4\x47\x74\x85\x0f\x6a\x7f\xcf\xfe\xb8\xb6\x93\x2f\x3c\xc5\xb8\x82\x7f\x9b\x21\xd3\x93\x93\x65\x7a\xf8\x13\x7c\x7c\xf5\x77\x00\x00\x00\xff\xff\xe7\x5e\x19\x48\x0c\x07\x00\x00")

func assetsUnversionedConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsUnversionedConsoleHtml,
		"assets/unversioned/console.html",
	)
}

func assetsUnversionedConsoleHtml() (*asset, error) {
	bytes, err := assetsUnversionedConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/unversioned/console.html", size: 1804, mode: os.FileMode(420), modTime: time.Unix(1547104628, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10AlphaConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5d\xab\xe3\x36\x10\x7d\xdf\x5f\x21\x54\xca\xbe\xd4\x52\x76\xbb\xd0\xc5\xd7\xbe\x50\x5a\x4a\x4b\x5b\xd8\x87\x4d\x5f\x97\x89\x3c\xb6\x27\x2b\x4b\xae\x46\x49\xae\x7b\xf1\x7f\x2f\xfe\x4a\x9c\xa4\x50\x7a\xe9\x92\x17\xcd\x99\xf1\x9c\x39\x23\x71\x92\xd5\xb1\xb1\xc2\x82\xab\x72\x89\x2e\x39\xb0\x7c\x7c\x25\x44\x56\x23\x14\xc3\x41\x88\xcc\x92\xfb\x2c\x02\xda\x5c\x92\xf1\x4e\x8a\xd8\xb5\x98\x4b\x6a\xa0\x42\xdd\xba\x4a\x8a\x3a\x60\x99\xcb\x3a\xc6\x96\x53\xad\x39\xfa\x00\x15\xaa\xca\xfb\xca\x22\xb4\xc4\xca\xf8\x46\xd7\xc0\x87\x00\x49\x15\xa0\xad\xff\xb4\x09\xba\x8a\x1c\x6a\xe3\x1d\x7b\x8b\x1a\x98\x31\xb2\x2e\xe1\x38\x70\xa8\xb1\xad\x9e\xf9\xd9\x04\x6a\xe3\x14\x88\x13\xb9\xc2\x9f\xd4\xa7\x4f\xe8\x8e\x22\x17\xcf\x13\x2a\x04\xb4\xf4\xb3\xe7\x98\x8a\xe7\x67\x35\x9f\xfb\xfe\x9b\x55\xf6\x83\x0f\x31\x15\x72\x4a\x0f\x41\xdf\xcb\x73\xde\x58\xfa\x03\x03\x93\x77\x63\x83\x4b\xb8\xea\x51\x40\x84\xef\x5b\xda\x06\x3b\xd6\x5c\xc2\xfb\x9a\x75\xaf\x6b\x68\x3d\x93\x31\xc8\xfc\x2b\x76\xd3\xcc\x4b\xb4\xaa\x38\x04\xfb\x21\x60\x49\x4f\xa9\x90\x7a\x35\xec\xb4\xb3\xdf\x7d\x81\xa9\x90\xc6\xd2\x95\x8e\xed\xf6\x97\x1f\x17\x11\xc3\x79\xd5\x0f\x1d\xec\x2c\x7e\x44\x8b\x0d\xc6\x30\xf1\xde\x60\x7d\x3f\x15\xf7\x0f\xd3\xee\xf5\x65\xf9\x99\x5e\xde\x44\xb6\xf3\x45\xb7\x5c\x4e\xec\x2c\x4e\x67\xd5\x00\xb9\x1f\xbc\x8b\xe8\xe2\xf9\x62\x0a\xe2\xd6\x42\x97\x8a\xd7\xce\x3b\x7c\xfd\x30\xc3\xbe\x05\x43\xb1\x4b\xc5\x66\x41\x62\x00\xc7\x14\xc7\xb5\xcd\x59\xa1\xde\x6e\x58\x58\x72\x08\x61\x2a\xeb\xef\x88\x14\xd7\xfe\xf4\x0f\x6c\x3b\xeb\xcd\xe7\x7b\xba\x37\x2f\xa0\xcb\xf4\x2c\x72\x8a\x0a\x3a\x0a\x2a\x72\x69\x3d\x14\xe4\x2a\x39\x3f\xcc\x29\x61\x2c\x30\xe7\xb2\x85\x0a\x93\xa5\x40\x8c\x9f\xe7\x72\xae\x6b\xc8\x25\x35\x52\x55\xc7\x54\xbc\xd9\x6c\x8e\xf5\x32\xd2\x89\x8a\x58\x8f\xd8\xd7\x0f\xb7\x7a\x4a\x8b\x4f\x0b\x08\x96\x2a\x97\x50\xc4\x86\x53\x61\xd0\x45\x0c\x4b\xaa\xf4\x2e\x26\x25\x34\x64\xbb\x54\x30\x38\x4e\x18\x03\x95\x4b\x7a\x7f\xe0\x48\x65\x97\x98\x69\x77\xb7\x5f\x9f\xa5\x0c\xf7\xda\x82\x5b\xd4\xdc\x2a\x98\x79\x98\xfe\xc2\x54\xbc\xc5\xe6\xe1\x8c\x37\x10\x2a\x72\x49\xf4\x6d\x2a\x92\x6f\xd7\x19\xe3\xad\x0f\xa9\xf8\xea\xfd\xbb\xe1\x77\xc1\x57\x9c\xbf\x4d\xfb\x52\x4a\x2d\x1b\xd5\xc3\x14\xe7\xfd\xea\x82\x8e\xf3\xab\x5b\x1d\x97\xeb\x98\x35\xc9\x65\xe8\xd5\x1b\x91\x8f\xeb\x0f\x2e\x76\x36\xaa\xe2\x1a\x31\xde\x7a\x98\x29\xdc\x9e\x95\xb1\xfe\x50\x94\x16\x02\x8e\x0e\x06\x7b\x78\xd2\x96\x76\xac\x47\xf9\x70\x42\xf6\x0d\xea\x77\xea\x3b\xb5\xd1\x86\xaf\x61\xd5\x90\x53\x86\x59\xea\xff\x40\xfb\x22\xeb\x1c\xcc\x63\x74\xcf\xb3\xc5\xe8\x41\xfb\x48\x2e\x4c\x0d\x81\x31\xe6\x72\xfb\xf1\xa7\xe4\xbd\xbc\xb6\x54\xc1\xc1\xfc\xff\xe4\x47\x74\x85\x0f\x6a\x7f\xcf\xfe\xb8\xb6\x93\x2f\x3c\xc5\xb8\x82\x7f\x9b\x21\xd3\x93\x93\x65\x7a\xf8\x13\x7c\x7c\xf5\x77\x00\x00\x00\xff\xff\xe7\x5e\x19\x48\x0c\x07\x00\x00")

func assetsV10AlphaConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10AlphaConsoleHtml,
		"assets/v1.0-alpha/console.html",
	)
}

func assetsV10AlphaConsoleHtml() (*asset, error) {
	bytes, err := assetsV10AlphaConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0-alpha/console.html", size: 1804, mode: os.FileMode(420), modTime: time.Unix(1547104664, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10ConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x5f\x8f\xe3\x34\x10\x7f\xdf\x4f\x31\x32\x42\xf7\x42\xec\xde\x71\x12\xa7\x6c\xbb\x12\xe2\x84\x40\x80\x74\x12\xb7\xbc\x9e\x66\x9d\x49\x32\x7b\x8e\x1d\x3c\x6e\xbb\xa5\xca\x77\x47\xf9\xd7\x66\xbb\xc0\xc1\x02\xea\x8b\xe7\xff\xfc\x7e\xe3\x8c\xbb\xae\x53\xe3\xc0\xa1\xaf\x36\x8a\x7c\xb6\x15\x75\x73\x05\xb0\xae\x09\x8b\xfe\x00\xb0\x76\xec\x3f\x42\x24\xb7\x51\x6c\x83\x57\x90\x0e\x2d\x6d\x14\x37\x58\x91\x69\x7d\xa5\xa0\x8e\x54\x6e\x54\x9d\x52\x2b\xb9\x31\x92\x42\xc4\x8a\x74\x15\x42\xe5\x08\x5b\x16\x6d\x43\x63\x6a\x94\x6d\xc4\xac\x8a\xd8\xd6\xbf\xba\x8c\x7c\xc5\x9e\x8c\x0d\x5e\x82\x23\x83\x22\x94\xc4\x94\xb8\xeb\x6b\xe8\x21\xad\x99\xea\x8b\x8d\xdc\xa6\x51\x80\x3d\xfb\x22\xec\xf5\x87\x0f\xe4\x77\xb0\x81\xe3\xa8\x05\xc0\x96\xbf\x0b\x92\x72\x38\x1e\xf5\x74\xee\xba\x2f\x16\xd6\x77\x21\xa6\x1c\xd4\x68\xee\x85\xae\x53\x27\xbb\x75\xfc\x0b\x45\xe1\xe0\x87\x04\x67\x71\x91\xa3\xc0\x84\x5f\xb7\x7c\x1b\xdd\xe0\x73\x16\x9f\xfa\x2c\x73\x3d\x56\x2d\x7b\xb2\x96\x44\x7e\xa0\xc3\xd8\xf3\x2c\x2d\x3c\xb6\xd1\xbd\x8b\x54\xf2\x43\x0e\xca\x2c\x9a\x1d\x39\xfb\x29\x14\x94\x83\xb2\x8e\x1f\xe1\xb8\xbd\xfd\xfe\xed\x0c\xa2\x3f\x2f\xf2\x91\xc7\x3b\x47\xef\xc9\x51\x43\x29\x8e\x75\x2f\x74\x5d\x37\x3a\x77\xd7\x23\xf7\xe6\x4c\xfe\xda\xcc\x77\x62\x7d\x17\x8a\xc3\x3c\x9c\x74\x70\x34\x9e\x75\x83\xec\xbf\x09\x3e\x91\x4f\xa7\xc1\x14\x2c\xad\xc3\x43\x0e\x2f\x7c\xf0\xf4\xe2\x7a\x52\x87\x16\x2d\xa7\x43\x0e\xab\x59\x93\x22\x7a\xe1\x34\xd0\x36\x59\x41\xbf\x5a\x09\x38\xf6\x84\x71\x74\xeb\x9e\x14\xd2\x52\x87\xfd\x1f\x54\xbb\x73\xc1\x7e\x7c\x5a\xee\xe5\x33\xca\xad\xcd\x04\x72\x94\x0a\xde\x01\x17\x1b\xe5\x02\x16\xec\x2b\x35\x5d\xcc\xd1\x60\x1d\x8a\x6c\x54\x8b\x15\x65\xb3\x03\x0c\xe1\x1b\x35\xf9\x35\xec\xb3\x9a\xb8\xaa\x53\x0e\x2f\x57\xab\x5d\x3d\xb7\xb4\xe7\x22\xd5\x83\xee\xf3\xeb\x4b\x3c\xa5\xa3\x87\x59\x89\x8e\x2b\x9f\x71\xa2\x46\x72\xb0\xe4\x13\xc5\xd9\x54\x06\x9f\xb2\x12\x1b\x76\x87\x1c\x04\xbd\x64\x42\x91\xcb\xd9\x7c\xbf\x95\xc4\xe5\x21\xb3\x23\x77\x97\xd1\x27\x28\xfd\x5c\x5b\xf4\x33\x9a\x4b\x04\x53\x1d\xe1\xdf\x28\x87\x57\xd4\x5c\x9f\xf4\x0d\xc6\x8a\x7d\x96\x42\x9b\x43\xf6\xe5\xd2\x62\x83\x0b\x31\x87\xcf\xde\xbc\xee\x7f\x67\xfd\xa2\xe6\x8f\x23\x5f\x5a\xeb\x99\x51\xd3\x77\x71\xe2\xd7\x14\xbc\x9b\x6e\xdd\xe2\x38\x8f\x63\xc2\xa4\xe6\xa6\x17\x77\x44\xdd\x2c\x03\xce\xeb\x6c\x40\x25\x35\x51\xba\xdc\x61\xb6\xf0\xf7\xa2\xad\x0b\xdb\xa2\x74\x18\x69\xd8\x60\x78\x8f\x0f\xc6\xf1\x9d\x98\x01\x3e\xee\x49\x42\x43\xe6\xb5\xfe\x4a\xaf\x8c\x95\xc7\x6a\xdd\xb0\xd7\x56\x44\x99\xe9\xde\x1c\x8f\xc0\x25\xf4\x5f\xe5\xcf\x09\x13\xdb\xb7\x1c\xa1\xeb\xae\x3e\xdd\x93\x91\xc1\xdf\xf4\x80\x86\x8c\x60\x6b\x8c\x42\x69\xa3\x6e\xdf\x7f\x9b\xbd\x51\x8f\xf7\x24\x48\xb4\xe7\xa0\x1d\xf9\x22\x44\x7d\xff\x34\xea\x66\xf9\x6d\xff\x49\xf4\x50\xf2\xaf\x63\x67\x6c\xe4\x84\xfe\x1e\x9e\x7f\xf5\x4e\xf4\x9b\x72\x78\x2a\x4e\xfb\xf4\x9f\xf1\xf2\x5f\x17\x7f\x1e\xbf\xff\x0b\x05\x9f\x9e\x53\x3f\x26\x5f\x4c\x53\x5a\x9b\x71\x8b\xaf\x4d\xff\x07\xe0\xe6\xea\xf7\x00\x00\x00\xff\xff\xe7\xc6\x9b\x76\x08\x08\x00\x00")

func assetsV10ConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10ConsoleHtml,
		"assets/v1.0/console.html",
	)
}

func assetsV10ConsoleHtml() (*asset, error) {
	bytes, err := assetsV10ConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0/console.html", size: 2056, mode: os.FileMode(420), modTime: time.Unix(1547104967, 0)}
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
	"assets/unversioned/console.html": assetsUnversionedConsoleHtml,
	"assets/v1.0-alpha/console.html": assetsV10AlphaConsoleHtml,
	"assets/v1.0/console.html": assetsV10ConsoleHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"unversioned": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsUnversionedConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10ConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0-alpha": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10AlphaConsoleHtml, map[string]*bintree{}},
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

