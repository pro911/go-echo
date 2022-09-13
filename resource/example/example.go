// Code generated for package resource by go-bindata DO NOT EDIT. (@generated)
// sources:
// resource/example/client.json
// resource/example/web.json
package resource

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _resourceExampleClientJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xed\x6f\x14\xc7\x19\xff\x1e\x29\xff\xc3\x74\xf3\x25\x48\xe0\x17\x5e\x2a\xd5\xc8\xaa\x52\x88\x0a\x6a\x4b\x08\x76\xda\x0f\x59\x74\x1a\xef\x8e\xf7\xb6\xde\x9b\x59\x66\xe6\x6c\x5c\x84\x44\xa9\x83\x0d\x31\x3e\x57\x98\x97\x38\x34\x09\x14\x0a\x0a\xae\x6d\xd4\x16\x9c\xb3\x81\x3f\x26\xb7\x7b\xe7\x4f\xfc\x0b\xd5\xec\xec\xee\xcd\xee\xce\x1e\x76\xe4\x44\x54\xea\x17\xb8\x7b\xde\xe6\x99\xdf\xf3\x32\xcf\xcc\xf9\xd3\x77\xdf\x01\xe0\x82\xf8\x07\x00\xc3\x87\x14\x61\x5e\x71\x6d\x63\x08\x18\x03\xc6\xfe\x98\xcc\x21\x75\x10\xaf\xf0\x69\x1f\x09\x86\x43\x7d\x2b\xe5\x61\x58\x8b\x88\xed\x07\xcd\xd6\xcb\xcf\x33\x2c\x46\x28\x37\x86\xc0\x60\xf2\x7d\x12\x51\xe6\x12\xac\x92\x6a\x90\x4e\x08\x6d\x1b\x4d\x22\x8f\xf8\x2e\x76\xba\xda\x1c\xf2\x3a\x53\x85\x7d\x4a\x38\x11\x94\xd8\x5d\x00\x8c\x3a\x43\xb4\x2f\xa2\xab\xe4\x44\xf6\x34\xe4\x55\x61\xfd\xd8\x90\x69\x7e\xc2\x10\x65\xa6\xe9\x53\xf2\x8b\x41\xd3\x3c\x4e\xa6\xb0\x47\xa0\xcd\x4c\x53\x31\xb1\x3f\x6f\xe0\x18\xc1\x1c\x61\xb1\x09\xa3\xbf\xbf\xb5\x71\xa9\xb5\xf1\x6d\x0d\x31\x06\x1d\xd4\x7e\xba\x19\x7c\xf5\x79\xb0\xba\xdc\xfa\xee\x6a\xd8\x68\x74\x5e\xad\xb7\x9a\x57\xa4\x44\x67\xed\x79\xf8\xf4\x72\x38\x77\x2b\xb8\x71\x3d\x68\x2e\xb5\x97\x67\xc2\x67\x73\xe1\x9f\xd7\xc2\x6f\xb6\x82\xad\xc6\xeb\xad\xf9\xa0\xb1\xd6\xda\x7c\x18\xcc\x7f\x16\x2c\x3e\x09\x1e\x2c\xb7\x1f\x5d\x8f\xad\xfd\xf3\x76\xb8\xfa\x1f\x93\x9a\xb8\xbf\x5f\x7e\x0e\x16\xaf\xb7\x97\x1e\x07\xeb\x8d\xe0\xca\x17\xaf\xb7\xe6\x3b\x2f\x56\x5a\x1b\x4d\x87\x84\xb7\x66\x5b\x9b\xcf\x82\xc5\x85\xce\xec\xb7\xc1\xb5\xc7\x41\xe3\x4e\x38\xb7\x18\x3c\x78\x24\xa5\x62\x0f\x83\xf5\xa7\xed\x2f\x37\x82\x17\x37\x5a\xcd\x85\xf6\xe6\x8d\xf0\xab\x99\xd6\x8b\x1b\xc2\x78\xbc\x40\xe4\x7a\xfb\xea\x5c\x78\x77\x25\x68\x3c\x17\x44\x36\x8d\x39\x3c\x0f\x86\x81\x29\x77\x7f\xc8\x34\x8e\x16\x14\x96\xbe\x0e\xe7\x16\xdb\xcb\x33\x91\xfb\xd7\xdb\x8f\xd6\xc1\x10\x08\xe7\x67\x83\xd5\x65\xc9\x92\x70\x8e\xf5\x39\x44\x4a\xb4\x36\x9a\x91\x31\x61\xc4\x87\xd6\x04\x74\x10\x88\x08\x8a\xe9\xce\xcb\x1b\xc1\x6c\xb3\xf3\x7c\x2d\x78\x39\x23\x28\xc4\xe7\x2e\xc1\xc0\x21\x95\x44\x43\xf8\xd4\xd7\xaf\xf1\x27\x58\xbd\x1f\xce\x3d\x6f\x3f\x59\x8b\x51\x8f\x30\x16\x02\x31\x0a\x40\xc4\xfd\x0c\x3a\x57\x47\x8c\x5f\x10\x74\x00\x12\x55\xa9\x11\x34\x2e\x87\x37\xd7\x25\x87\x71\xea\x62\x07\x88\x94\x06\xc3\x60\x30\x5a\xec\xa2\xba\x5a\x78\xf7\x7a\x70\xed\x5e\xfb\xc9\x9a\x0c\x6d\xd9\x6a\xcc\x27\x98\xa1\xdc\x72\x52\x45\x5d\xce\xc5\xfc\xd0\x41\xe0\xda\xe9\x5a\x79\x0f\x0e\x1e\x55\x05\x25\x0e\x87\x8e\xa6\x66\xa3\x14\x69\xbd\x5a\xdd\x7e\xb8\xde\x5e\xf9\x47\x42\xa6\xc8\x47\x90\x23\x1b\x74\xee\x3d\x6e\x3f\x68\x06\x8d\x35\x91\x1f\x37\xd7\xdb\x9b\x33\xaf\xb7\xe6\xdb\x4f\x37\x5b\x5b\x5b\xad\xe6\x42\x30\x37\xdb\xbe\x3a\x2b\x13\x4f\xaa\xa6\x8a\xb1\x13\x55\x32\x36\x36\x0d\x86\xc1\xe1\x2e\x0e\xea\x56\x4f\xc1\x1a\xda\x29\xb0\x85\x9d\x96\x58\xdb\x21\x70\xa5\x71\x2a\xc6\x2a\x5c\x78\x18\x34\xfe\xfe\xfd\xa5\xcb\x69\x2d\x48\xca\xeb\xad\x79\x86\xe8\xa4\x6b\xa1\x94\x11\x6c\x5d\x0a\x1a\x7f\xed\xac\xff\xa5\xbd\xf4\x58\x14\x6d\xa4\x1f\xd5\x85\x14\x8c\x82\x7b\x12\x8f\x93\x11\xf9\x3d\x75\x53\x5a\x0c\xae\x7c\x26\xb4\x6e\x7d\x17\xfe\xeb\x66\x39\x1e\xad\x8d\xa6\x92\x91\xfb\xd5\xbd\xa5\x2c\x89\x42\x1c\x14\xdf\x02\xbf\x46\x3c\x59\x1a\xbc\xaf\x68\xef\x03\x14\xf1\x3a\xc5\x2c\xa1\x4a\xc5\x7d\xe0\x42\x0a\x46\xc6\x80\xc0\x18\xbc\xaf\xc4\x4d\x31\xa0\xe2\x9f\x18\x50\x20\x8d\x2a\xd6\x02\x07\x0e\x38\xa4\x42\xea\x7c\xd8\xf7\xea\x8e\x8b\xd9\xb0\xe8\xf7\x43\x7d\xa0\xdb\x43\xa5\x78\x78\xf7\x6a\x78\xfb\x59\x70\xe9\x8b\xf0\xea\xa3\xce\xbd\xf9\xd6\xc6\xb5\xed\xbf\xdd\x6f\x6d\x2d\x87\xd7\x1e\x6e\x2f\x89\x1e\x15\xde\x59\x0b\xbe\xfc\xa6\xb5\xd1\x94\x0d\x68\xfb\xf6\xea\xf6\xfd\x3b\xdf\x5f\xba\xac\x5f\x4d\x5d\x20\xd3\xa4\xe3\xc8\xb0\x6c\xf3\x4f\xba\x77\x5f\x2e\x60\x79\x29\x00\x0c\x05\xda\x22\x37\xb1\x23\xda\xbf\xde\x81\xac\x1b\x42\x4e\xbf\x70\x51\xa5\x4e\xbd\xe8\x28\x25\x7d\xd0\x77\x7d\xc2\x78\x9f\x85\x35\x62\xdc\x53\x4f\x40\x85\x51\x43\xbc\x4a\xa2\x63\x5a\xdd\x42\x51\x8e\xca\x48\xeb\x36\x07\x80\x51\x45\xd0\x46\xd4\x18\x02\x9f\x9e\x2d\xa8\x02\x60\x8c\x11\x7b\x5a\xaf\x29\x2c\xc3\x29\xb1\xfc\x05\x13\x9b\xdc\x8c\xc6\x00\xd3\x18\x32\x8d\xee\x76\x4c\xc3\xc4\x17\x8b\x2e\x81\x78\xd0\x80\x35\xc4\xe5\xda\x3a\x09\xa0\x5f\x35\x52\x9e\x40\xc2\x2b\x39\x79\x68\xcd\x47\x52\x93\xd0\xab\x47\x21\xe9\x05\x70\x2a\x6d\x23\x66\x51\x37\x3a\x77\x84\x4e\x0f\x49\x4c\x78\x05\xd7\x3d\x4f\x1b\x97\x54\x6a\xdc\x45\x9e\x9d\x4e\x4c\x23\x51\xc3\xea\x61\x34\x11\x1c\x45\xe7\x79\x0f\x31\x97\x55\xac\x2a\xb2\x26\x90\x88\xfc\xa0\x5e\xec\xa2\x8e\x7c\xb6\x48\x2c\xc8\x5d\xd4\x65\x8f\xec\x08\x25\xe9\xc3\xea\x96\x85\x58\xa1\xfa\xba\xea\x32\x45\x76\x92\x03\x3a\x07\x75\x29\x89\x28\x25\xf4\xc7\x5a\xb0\x80\x48\x96\x90\x77\xc8\x50\xfa\xea\xff\x78\xf7\x38\xa5\xad\xa5\x9f\xa2\x7b\x50\xd9\x3f\x5c\x3b\xea\x1e\x83\xa6\x11\x9d\x3c\x7b\x16\xc1\xff\xe7\xf4\x1b\x72\x5a\xfd\xaa\x7c\x49\x3f\xc6\x1f\x12\xc7\x75\x29\x51\x68\x9e\xb1\x8a\xf8\x4f\xea\xed\xee\x7e\x39\x85\xc6\x18\xb1\x26\x10\x2f\xb9\x64\x16\xf9\x7b\x7d\xd3\xec\xd6\xc7\x19\x38\x65\xf4\xda\x7a\x5a\x03\x99\xcb\xa7\x36\x28\xdd\xd8\x1b\xe7\xea\x88\x4e\xef\x56\x29\x2e\xfa\x29\xc6\x86\xfa\xfb\x33\xa5\xdf\x5f\x40\xa4\x3c\x28\x5d\xc8\x84\xfc\x31\x82\xc7\x5d\x27\xb3\x23\x17\x8f\x13\x5a\x83\x42\x6f\xc4\xfd\x93\x80\xfd\x48\xd7\x28\x45\x16\xc1\x18\x59\xfc\x54\xbd\x66\x0c\x81\x01\x0d\x67\xd4\x8d\x62\x75\x64\x60\x40\xe1\xb2\x2a\x9c\x40\x27\x20\xb6\x59\x72\x27\xef\x97\x0e\xf4\xb9\x4a\x4b\x54\xc4\x84\x95\x8f\xea\x3c\xbb\x86\x54\xf9\x70\x12\x61\x1e\xf7\x5c\xc3\xc8\x73\x4f\x92\x88\xff\x5b\x97\x71\x84\x11\x65\xd9\xf6\x94\x0a\xfd\x3e\xcd\x10\x63\xf2\x50\x1e\x9b\xf8\x52\xa2\xda\x4f\x68\xa3\x9a\x83\x7a\x67\x91\xf9\xa1\xd5\x60\x93\xb2\xc7\x96\xf0\xd6\x6c\x78\x77\xe5\x47\x2b\x82\x1d\x54\xfa\x7b\xef\x81\x70\xe5\x7e\xe7\xd5\x42\xeb\xc5\xab\xf6\xd2\xe3\x0f\x7c\xf7\x34\x61\xdc\xc4\x27\x90\xe7\x11\xf0\x07\x42\x3d\xfb\x67\x79\x70\xa3\x19\x06\x62\x27\x9a\x61\x0e\x0c\xfe\x50\x58\xa0\xef\x96\xc1\x12\x5d\xc7\x8c\xbd\xef\x05\xa7\x3f\x1a\x19\xed\xe6\x03\xb1\x22\x1b\x17\x2e\x66\x48\x95\x38\x17\xfa\xf3\xc9\x51\xe5\xdc\x2f\xa6\xc7\x2f\x3f\x16\x9d\x60\x98\x0b\x9c\x7b\xe1\xbe\x5b\x23\x3b\x0d\x95\x22\x0f\xeb\x51\x61\xaa\x2d\x29\xc1\x7a\x0c\x32\xd7\xca\x5c\xbc\x26\x26\x0b\x57\xae\x78\x2a\xcf\x9e\x53\xdd\x29\xdc\x50\xce\x16\xd5\xd2\x18\x82\x34\xdf\x40\x15\x6b\x65\x5a\x91\x47\x79\x25\x31\x66\x25\xb9\xc0\x49\x2d\xe7\x8a\x0f\x19\x9b\x22\x34\x0a\xe5\xe0\xc1\x43\x87\x8f\xfc\xbc\xcc\xba\xed\x3a\x9a\xa9\x27\x63\xbe\x87\xed\x1c\x8b\x22\xe8\xd5\x34\x74\x4c\xb0\xa5\x33\x05\x3d\x87\x50\x97\x57\x75\x3a\xe7\x88\xaf\xb3\x64\x69\x88\x56\x99\x7d\xe2\xc3\x73\xbd\x43\x52\x85\x53\x13\x85\xcd\x8b\xf4\x38\xa9\xdb\x9f\x60\xfc\x46\x1b\xfb\x5e\x3b\x11\x58\xee\x02\x14\x74\x9e\x53\x78\x1c\x72\xa8\x5b\xc6\xd7\x81\x62\x23\x0f\x39\x50\x7f\x9f\x33\xb8\x5b\x43\x8c\xc3\x9a\x4e\xd1\xc5\x96\x57\xb7\xd1\x69\x38\xed\x11\x68\x9f\x80\xac\x9a\xf6\x29\x0d\x56\x70\x8a\x4d\x1e\x2e\x82\x15\xcd\x90\x7a\x54\x18\xb2\x28\xe2\x7a\x1e\x45\x8e\xde\x61\xe5\x8e\x50\xe0\x30\xd1\xde\x46\xc9\x04\xd2\x29\x42\xdb\xfe\xa0\xce\xab\x02\xba\x51\xf2\x71\x3c\x76\x94\xee\x06\xf3\x28\x55\xf7\x22\xed\x6d\x52\x83\xae\xce\xa3\x29\x42\x27\x44\x97\xd5\x6f\xd4\x76\x19\x1c\xf3\xd0\x19\xc4\xe9\xf4\x99\xb4\x11\x96\xf9\x8b\x6c\x07\x39\x34\x3a\x29\x74\x01\x28\x03\xc5\xf2\x5c\x84\x79\x6f\xee\x48\x14\xa6\x5d\xa4\x68\xaf\x9c\x1a\x83\x0c\x7d\x72\xc6\xd5\x70\xe4\xfc\xc8\x46\xc9\x88\xeb\xe0\x5e\x45\x49\x44\x9d\x0d\x16\x36\x6a\x11\xcc\xea\x35\x44\xf5\xf9\x94\x70\x4b\x37\xc3\x5c\x07\x43\x5e\xa7\xe8\x77\xe9\x29\x57\x4c\xa0\x0f\x6b\x3e\x9f\x3e\x2d\xc6\xd3\xae\x9f\x07\x06\xb5\x55\xf3\x2b\x62\x4f\xa7\x25\x53\x30\x94\xd8\x38\x91\x0c\xcd\x79\x99\xb2\x5e\xd9\x3d\xc3\x8d\xc1\xbe\x81\x3d\x08\xc8\x24\xa2\xee\xb8\xab\x6d\x41\x16\xf4\xbc\x31\x28\x0f\xf7\xbc\x45\x91\x31\xa5\x58\xf2\x34\x9f\x94\x10\x6a\xe6\xf8\xe2\xbd\xd8\xa8\x11\x3b\xda\x81\x98\xbc\x0f\xd8\xa2\xcd\xed\x2f\xbb\x19\xa8\x4b\xe6\x5f\x33\xdf\xf0\x88\xb5\xa3\x47\xa9\xdc\x1b\x53\x9e\x1b\x9f\xcb\xd1\x1e\xf2\xbc\xf4\xac\xcf\x4e\x22\x69\x98\x4a\xdf\xcd\xb2\xcf\x5f\x2a\x4f\xb9\x91\xaa\x6f\x0c\xba\x4b\xb1\xa0\x55\x04\x54\x65\xf7\x27\x34\x29\x7f\xca\xcb\xfe\x4c\x88\x2a\x12\x33\x61\x4e\x94\x0b\xf1\x50\x9f\x47\x9c\xf7\x4d\x63\xfb\xfe\x8c\x7c\xbe\xee\xcc\x2c\x87\x77\x57\x86\x74\x93\xae\xb1\xef\x68\xc6\x07\x2e\xfb\x55\xce\x52\xb0\xb8\xb0\x23\x4b\x1a\xa7\xdf\x7c\xbf\x7c\x3b\x12\x22\x9a\xf8\xcb\x33\x62\x34\x3f\x8d\xed\x45\x42\x68\xe0\x7a\xe3\xc5\xfa\xed\x40\x4b\x1e\xc4\x3f\x71\xfd\x68\xe0\xb2\x08\x99\x70\xd1\x6e\x1f\x22\x28\x62\xe3\x75\x6f\x67\x5a\xf2\x83\xf2\x7a\x54\x7c\x7c\xd3\x3f\xb9\x69\x4b\x3c\xbb\x48\x66\x12\x38\xef\x23\xab\x38\xb2\x27\x73\x4b\x38\xb7\x18\x5c\xfb\x3a\x3f\xe9\xb1\xe3\x68\x1c\xd6\x3d\x5e\x80\xd4\xb0\x64\x43\x3e\xa8\xbc\x61\xc4\xf4\xe8\x0f\x02\x92\x57\x80\x3f\x32\x82\x75\x27\xcb\x74\x22\xc0\xac\x2a\xaa\xc1\x9c\x48\x72\x75\xcc\x9f\xc5\x52\x36\xfb\x8e\xa7\x3b\x40\x34\xef\x85\x7b\x8b\x56\xf0\xe0\x69\xe7\xdf\x0f\x7b\xa0\x95\x3f\xb7\x63\xb8\x0e\x0f\x1c\x7e\x6b\xe0\xca\x25\x5e\x52\x96\x15\x8a\x58\x36\xe2\x86\x4a\xaf\x48\x64\xb2\x17\x1d\xa1\x8c\xce\xc3\x9a\xef\x21\x75\xef\x82\xec\x91\xb8\xd2\xbb\x54\x84\xc5\x00\x5b\x81\x6e\x25\x45\xb9\xc0\x14\xe3\x3c\xa2\x95\x78\x67\x5d\xb6\xaa\xd4\xad\x0f\x4c\x30\xaa\xd4\x20\xaf\xc6\xcc\xf8\x69\x44\xb9\xbb\x7b\x6e\x74\xec\x24\x48\xf4\x7c\x6b\x79\xf7\x9d\xb3\xff\x0d\x00\x00\xff\xff\x67\x2c\x7e\xa8\xf4\x23\x00\x00")

func resourceExampleClientJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourceExampleClientJson,
		"resource/example/client.json",
	)
}

func resourceExampleClientJson() (*asset, error) {
	bytes, err := resourceExampleClientJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/example/client.json", size: 9204, mode: os.FileMode(438), modTime: time.Unix(1662732672, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourceExampleWebJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x73\xd4\x46\xf6\x7f\x4f\x55\xbe\x43\xff\x95\x07\x9b\x7f\xc0\xf6\x00\xd9\xca\xda\xeb\xdd\xca\x26\xa9\x25\xb5\xbb\x09\x01\x67\xf3\x90\x49\x4d\xb5\xa5\xb6\x46\x6b\x8d\x5a\xa8\x7b\x6c\xbc\x14\x55\x06\x1c\x5f\x88\x6f\xbb\x06\x63\x1c\x93\x8b\x63\x02\x15\x83\xc7\x54\x2e\x4c\xc6\x06\x7f\x97\x44\xd2\xc8\x4f\x7c\x85\xad\x56\x4b\x33\xba\xb4\x06\x9b\x35\x29\x52\xe5\x17\x98\xe9\x73\xe9\x73\x7e\xe7\xf4\xe9\xd3\xc7\xf3\xd1\xcb\x2f\x01\x70\x81\xfd\x03\x80\x64\x42\x0b\x19\xb4\xa0\x29\x52\x37\x90\xba\xa4\xa3\xc1\x32\x85\x96\x8a\x68\x81\x8e\x98\x88\x11\x54\xcb\x94\x1b\x34\x03\x96\xfc\xc5\xfa\x5a\xcd\x7e\xfc\x69\x8c\x44\xb0\x45\xa5\x6e\x90\x0b\xbf\x0f\x21\x8b\x68\xd8\x88\x2e\x95\xa0\x35\xc8\xa4\x15\x34\x84\x74\x6c\x6a\x86\xda\x94\xa6\x90\x96\x49\x94\xd9\xb4\x30\xc5\x6c\x25\x30\x17\x00\xa9\x4c\x90\xd5\xe1\xaf\x47\x97\x43\xde\xd3\x90\x16\x99\xf6\x37\xbb\xf3\xf9\x0f\x08\xb2\x48\x3e\x6f\x5a\xf8\xf7\xb9\x7c\xfe\x2d\x3c\x6c\xe8\x18\x2a\x24\x9f\x8f\xa8\x38\x9a\x54\xf0\x26\x36\x28\x32\x98\x13\x52\x67\xa7\x5d\x1d\xb5\xab\xdf\x96\x10\x21\x50\x45\xf5\x07\x5b\xce\xe7\x9f\x3a\x1b\xcb\xf6\x4f\x53\xee\xdc\x9c\xb7\xb3\x69\xd7\xc6\x39\x87\x57\x79\xe8\x3e\xb8\xec\x4e\x2e\x3a\x0b\x33\x4e\xed\x5a\x7d\x79\xcc\xfd\x71\xd2\xbd\x54\x71\xbf\xdc\x76\xb6\xe7\x9e\x6c\x4f\x3b\x73\x15\x7b\xeb\xb6\x33\xfd\x89\x33\xbf\xee\xac\x2d\xd7\xef\xcc\x04\xda\xee\xdf\x70\x37\x7e\xc8\x5b\x79\xa3\xb3\x93\x7f\x76\xe6\x67\xea\xd7\xee\x3a\x9b\x73\xce\xf8\xcd\x27\xdb\xd3\xde\xa3\x7b\x76\xb5\xa6\x62\x77\x71\xc2\xde\xfa\xd1\x99\x9f\xf5\x26\xbe\x75\xae\xde\x75\xe6\x96\xdc\xc9\x79\x67\xed\x0e\xe7\x0a\x2c\x74\x36\x1f\xd4\x3f\xab\x3a\x8f\x16\xec\xda\x6c\x7d\x6b\xc1\xfd\x7c\xcc\x7e\xb4\xc0\x94\x07\x1b\xf8\xa6\xd7\xa7\x26\xdd\x95\x7b\xce\xdc\x43\xb6\x48\x46\x0c\x0a\xcf\x83\x5e\x90\xe7\xde\x9f\xc8\x4b\x3d\x29\x81\x6b\x5f\xb8\x93\xf3\xf5\xe5\x31\xdf\xfc\x99\xfa\x9d\x4d\xd0\x0d\xdc\xe9\x09\x67\x63\x99\x93\x38\x9c\xfd\x1d\x2a\xe6\x1c\x76\xb5\xe6\x2b\x63\x4a\x4c\x28\x0f\x42\x15\x01\x7f\x21\xa2\xda\x7b\xbc\xe0\x4c\xd4\xbc\x87\x15\xe7\xf1\x18\x5b\xc1\x26\xd5\xb0\x01\x54\x5c\x08\x25\x98\x4d\x1d\x9d\x02\x7b\x9c\x8d\x55\x77\xf2\x61\x7d\xbd\x12\xa0\xee\x63\xcc\x18\x02\x14\x00\x8b\xfb\x19\x74\xae\x8c\x08\xbd\xc0\xd6\x01\x08\x45\xb9\x84\x33\x77\xd9\xbd\xbe\xc9\x29\x84\x5a\x9a\xa1\x02\x96\xd2\xa0\x17\xe4\xfc\xcd\x2e\x46\x77\x73\x57\x66\x9c\xab\x5f\xd5\xd7\x2b\x3c\xb4\x59\xbb\x11\x13\x1b\x04\x25\xb6\xe3\x22\xd1\xed\x34\x83\x9e\x38\x0e\x34\xa5\xb1\x57\xd2\x82\xe3\x3d\x51\x46\x8e\xc3\x89\x9e\x86\x5a\x3f\x45\xec\x9d\x8d\xdd\xdb\x9b\xf5\x7b\xdf\x84\xcb\x16\x32\x11\xa4\x48\x01\xde\x57\x77\xeb\x6b\x35\x67\xae\xc2\xf2\xe3\xfa\x66\x7d\x6b\xec\xc9\xf6\x74\xfd\xc1\x96\xbd\xbd\x6d\xd7\x66\x9d\xc9\x89\xfa\xd4\x04\x4f\x3c\x2e\xda\x10\x0c\x8c\x28\xe2\xfe\xfe\x11\xd0\x0b\x4e\x36\x71\x88\xba\xfa\x2e\x2c\xa1\xbd\x02\x9b\xf2\x34\x43\xdb\x1e\x81\xcb\x8c\x53\x3a\x56\xee\xec\x6d\x67\xee\xeb\x5f\x46\x2f\x37\xce\x02\x5f\x79\xb2\x3d\x4d\x90\x35\xa4\xc9\xa8\x41\x70\xb6\x47\x9d\xb9\x7f\x7b\x9b\x57\xea\xd7\xee\xb2\x43\xeb\xcb\xfb\xe7\x82\x33\xfa\xc1\x7d\xc7\x18\xc0\x67\xf9\xf7\x86\x99\x5c\xa3\x33\xfe\x09\x93\x5a\xfc\xc9\xfd\xee\x7a\x36\x1e\x76\xb5\x16\xc9\xc8\xa3\x51\xdf\x1a\x24\x8e\x42\x10\x14\x53\x06\x7f\x41\x34\xdc\x1a\xb4\x47\xa4\x8f\x00\x0b\xd1\xb2\x65\x90\x70\x95\x0b\x1e\x01\x17\x1a\x60\xc4\x14\x30\x8c\x41\x7b\x24\x6e\x11\x05\x51\xfc\x43\x05\x11\x48\xfd\x13\x2b\x83\x63\xc7\x54\x5c\xc0\x65\xda\x6b\xea\x65\x55\x33\x48\x2f\xab\xf7\xdd\x1d\xa0\x59\x43\x39\xbb\xbb\x32\xe5\xde\xf8\xd1\x19\xbd\xe9\x4e\xdd\xf1\xbe\x9a\xb6\xab\x57\x77\x6f\xad\xda\xdb\xcb\xee\xd5\xdb\xbb\xd7\x58\x8d\x72\x97\x2a\xce\x67\x5f\xda\xd5\x1a\x2f\x40\xbb\x37\x36\x76\x57\x97\x7e\x19\xbd\x2c\xde\x2d\xba\x41\xac\x48\x07\x91\x21\xf1\xe2\x1f\x56\xef\x8e\x44\xc0\x92\x5c\x00\x48\x11\x68\xd3\xd4\x50\x0f\x2b\xff\x62\x03\xe2\x66\x30\x3e\xf1\xc6\x69\x91\xb2\xa5\xfb\x57\x29\xee\x80\xa6\x66\x62\x42\x3b\x64\x43\xc0\x46\xf5\xe8\x0d\x18\x21\x94\x10\x2d\x62\xff\x9a\x8e\xba\x90\xe6\xb3\x78\xa4\x45\xce\x01\x20\x15\x11\x54\x90\x25\x75\x83\x8f\x3e\x4e\x89\x02\x20\xf5\x63\x65\x44\x2c\xc9\x34\xc3\x61\xb6\xfd\x85\xbc\x91\xa7\x79\xbf\x0d\xc8\x4b\xdd\x79\xa9\xe9\x4e\x5e\xca\x1b\x17\xd3\x26\x81\xa0\xd1\x80\x25\x44\xf9\xde\x22\x0e\x20\xde\xd5\x17\x1e\x44\xcc\x2a\xde\x79\x08\xd5\xfb\x5c\x43\x50\x2f\xfb\x21\x69\x05\x70\x83\x5b\x41\x44\xb6\x34\xff\xde\x61\x32\x2d\x38\x0d\x4c\x0b\x46\x59\xd7\x85\x71\x69\x70\x0d\x68\x48\x57\x1a\x1d\xd3\x59\xbf\x60\xb5\x50\x1a\x32\xf6\xa1\xf3\xb4\x05\x9b\x46\x0a\x72\x11\xc9\x83\x88\x45\x3e\x27\x66\xbb\x28\x5a\xfe\x38\xbd\x98\xe2\xbb\x28\xca\x1e\x5e\x11\x32\xd2\x87\x94\x65\x19\x91\xd4\xe9\x6b\x8a\xf3\x14\xd9\x4b\x0e\x88\x0c\x14\xa5\x24\xb2\x2c\x6c\x3d\xaf\x0d\x53\x88\xc4\x17\x92\x06\x49\x91\xba\xfa\x1b\xaf\x1e\xef\x0a\xcf\xd2\xaf\x51\x3d\x2c\x5e\x3f\x34\xc5\xaf\x1e\xb9\xbc\xe4\xdf\x3c\x07\x16\xc1\xc3\x9c\x7e\x4a\x4e\x47\xbf\x46\xbe\x34\x3e\x06\x1f\x42\xc3\x45\x29\x91\x2a\x9e\x81\x08\xfb\x8f\xcb\xed\xef\x7d\x39\x8c\xfa\x09\x96\x07\x11\xcd\x78\x64\xa6\xe9\x07\xfd\xd2\x6c\x9e\x8f\x33\x70\x58\x6a\xe5\x7a\xe3\x0c\xc4\x1e\x9f\xc2\xa0\x34\x63\x2f\x9d\x2b\x23\x6b\x64\xbf\x42\xc1\xa1\x1f\x26\xa4\xbb\xb3\x33\x76\xf4\x3b\x53\x88\x64\x07\xa5\x09\x19\xe3\x7f\x13\x1b\x03\x9a\x1a\xf3\x48\x33\x06\xb0\x55\x82\x4c\xee\xac\xf6\x2f\x06\xfb\x6b\x4d\xa5\x16\x92\xb1\x61\x20\x99\xbe\x5b\x2e\x49\xdd\xa0\x4b\x40\xe9\xd3\xfc\x58\xbd\xd6\xd5\x15\xa1\x92\x22\x1c\x44\xa7\xa0\xa1\x90\xf0\x4d\xde\xc9\x0d\xe8\xd0\x22\x25\x31\xc2\xc6\xb4\xbc\x57\xa6\xf1\x3d\xb8\xc8\xdb\x43\xc8\xa0\x41\xcd\x95\xa4\x24\xf5\x1d\xec\xd3\xff\xa6\x11\x8a\x0c\x64\x91\x78\x79\x6a\x30\xfd\xa3\x91\x21\xd2\xd0\x89\x24\x36\xc1\xa3\x24\xaa\x3f\x5c\xeb\x13\x5c\xd4\x7b\x8b\xcc\xb3\x9e\x06\x05\x67\x0d\x5b\xdc\xc5\x09\x77\xe5\xde\x73\x3b\x04\x7b\x38\xe9\xaf\xbc\x02\xdc\x7b\xab\xde\xce\xac\xfd\x68\xa7\x7e\xed\xee\x1b\xa6\x76\x1a\x13\x9a\x37\x4e\x21\x5d\xc7\xe0\x43\x6c\xe9\xca\xff\x25\xc1\xf5\x7b\x18\x68\xa8\x7e\x0f\x73\x2c\xf7\xac\xb0\x40\x53\xcb\x82\xc5\x7f\x8e\x49\x07\x5f\x0b\x4e\xbf\x77\xb6\xaf\x99\x0f\x58\xf6\x75\x5c\xb8\x18\x5b\x2a\x04\xb9\xd0\x99\x4c\x8e\x22\xa5\x66\x3a\x3d\xfe\xf4\x3e\xab\x04\xbd\x94\xe1\xdc\x0a\xf7\xfd\x2a\xd9\x6b\xa8\x22\xfc\xb0\xec\x1f\xcc\x68\x49\x0a\xb1\xee\x87\x44\x93\x63\x0f\xaf\xc1\xa1\xd4\x93\x2b\xe8\xca\xe3\xf7\x54\xb3\x0b\x97\x22\x77\x4b\x54\x53\x3f\x82\x56\xb2\x80\x46\xb4\x65\x49\xf9\x16\x25\x85\x58\x9b\x15\xe6\x02\xc5\xa5\x84\x29\x26\x24\x64\x18\x5b\x7e\x28\x73\xc7\x4f\x9c\x7c\xed\x77\x59\xda\x15\x4d\x15\x74\x3d\x31\xf5\x2d\x74\x27\x48\x16\x82\x7a\x49\xb0\x6e\x60\x43\x16\xa9\x82\xba\x8a\x2d\x8d\x16\x45\x32\xe7\xb0\x29\xd2\x24\x0b\x16\xe5\x2c\xfd\xd8\x84\xe7\x5a\x87\xa4\x08\x87\x07\x53\xce\xb3\xf4\x78\x47\xe4\x1f\x23\xfc\x55\x18\xfb\x56\x9e\x30\x2c\xf7\x01\x0a\x3a\x4f\x2d\xf8\x16\xa4\x50\xb4\x8d\x29\x02\x45\x41\x3a\x52\xa1\xf8\x3d\x27\x51\xad\x84\x08\x85\x25\x91\xa0\x66\xc8\x7a\x59\x41\xa7\xe1\x88\x8e\xa1\x72\x0a\x92\x62\xa3\x4e\x09\xb0\x82\xc3\x64\xe8\x64\x1a\x2c\xbf\x87\x14\xa3\x42\x90\x6c\x21\x2a\xa6\x59\x48\x15\x1b\x1c\x79\x23\xa4\x28\x84\x95\xb7\x3e\x3c\x88\x44\x82\x50\x51\xde\x28\xd3\x22\x83\xae\x0f\xbf\x1f\xb4\x1d\x99\xde\x18\xd4\x4f\xd5\x83\x48\x7b\x05\x97\xa0\x26\xb2\x68\x18\x5b\x83\xac\xca\x8a\x1d\x55\x34\x02\xfb\x75\x74\x06\x51\x6b\xe4\x4c\xa3\x10\x66\xd9\x8b\x14\x15\xa9\x96\x7f\x53\x88\x02\x90\x05\x8a\xac\x6b\xc8\xa0\xad\xa9\x67\xfd\x30\xed\x23\x45\x5b\xe5\x54\x3f\x24\xe8\x83\x33\x9a\x80\xc2\xfb\x47\xd2\x87\xcf\x6a\xaa\xd1\xea\x50\x62\x76\xce\x72\x29\x47\x65\x6c\x90\x72\x09\x59\xe2\x7c\x0a\xa9\x99\xce\x10\x4d\x35\x20\x2d\x5b\xe8\xef\x8d\x5b\x2e\x9d\x40\x6f\x97\x4c\x3a\x72\x9a\xb5\xa7\x4d\x3b\x8f\xe5\x84\xa7\xe6\xcf\x58\x19\x69\x1c\x99\x94\xa2\x50\xc7\xa9\xb0\x69\x4e\xf2\x64\xd5\xca\xe6\x1d\x2e\xe5\x3a\xba\x0e\x20\x20\x43\xc8\xd2\x06\x34\x61\x09\x92\xa1\xae\xf7\x43\x7e\xb9\x27\x35\xb2\x8c\xc9\xc4\x92\x36\xf2\x29\x12\x42\x41\x1f\x9f\x7e\x17\x4b\x25\xac\xf8\x1e\xb0\xce\xfb\x98\xc2\xca\xdc\xd1\xac\x97\x41\x74\xcb\xe4\x34\xf3\x29\x43\xac\x3d\x0d\xa5\x12\x33\xa6\x24\x35\xb8\x97\x7d\x1f\x92\xb4\xc6\x5d\x1f\xef\x44\x1a\x61\xca\x9c\x9b\xc5\xc7\x5f\x51\x5a\xe4\x45\x1a\x9d\x31\x88\x1e\xc5\x6c\xad\xc0\xa0\xca\x7a\x3f\xa1\x21\xfe\xa7\xbc\xf8\x9f\x09\x51\x81\x63\xc6\xd4\xb1\xe3\x82\x75\xd4\xa1\x63\xb5\x3d\x2f\xed\xae\x8e\xf1\xf1\xb5\x37\xb6\xec\xae\xdc\xeb\x16\x75\xba\xd2\x91\x9e\x98\x0d\x94\xd7\xab\x84\x26\x67\x7e\x76\x4f\x9a\x04\x46\x3f\xfd\x7d\xf9\x62\x24\x84\xdf\xf1\x67\x67\x44\x5f\xb2\x1b\x3b\x88\x84\x10\xc0\xf5\xd4\x87\xf5\x8b\x81\x16\xbf\x88\x7f\xe5\xf3\x23\x80\x4b\xc6\x78\x50\x43\xfb\x1d\x44\x58\x88\x0c\x94\xf5\xbd\x49\xf1\x0f\x91\xe9\x51\x7a\xf8\x26\x1e\xb9\x09\x8f\x78\x7c\x93\x58\x27\x70\xde\x44\x72\xba\x65\x0f\xfb\x16\x77\x72\xde\xb9\xfa\x45\xb2\xd3\x23\x6f\xa1\x01\x58\xd6\x69\x0a\x52\x49\xe6\x05\xf9\x78\x64\x86\x11\xac\xfb\x3f\x08\x08\xa7\x00\xff\x24\xd8\x10\xdd\x2c\x23\x21\x03\x91\x8b\xa8\x04\x13\x2c\xe1\xd3\x31\x79\x17\x73\xde\xf8\x1c\x4f\x74\x81\x08\xe6\x85\x07\x8b\x96\xb3\xf6\xc0\xfb\xfe\x76\x0b\xb4\x92\xf7\x76\x00\xd7\xc9\xae\x93\x2f\x0c\x5c\x89\xc4\x0b\x8f\x65\xc1\x42\x24\x1e\x71\x29\xba\x5e\xe0\xc8\xc4\x1f\x3a\x4c\x18\x9d\x87\x25\x53\x47\x51\xdf\xd9\xb2\x8e\x83\x93\xde\x5c\x45\x06\x6b\x60\x0b\x50\x2b\x34\x50\x4e\x11\x59\x3b\x8f\xac\x42\xe0\x59\x93\x1c\x15\x6a\x9e\x0f\x03\x1b\xa8\x50\x82\xb4\x18\x10\x83\xd1\x48\xe4\xed\xae\x6b\xfe\xb5\x13\x22\xf1\x3f\xcf\x5a\x98\x90\x81\xf4\x02\x94\x35\xaa\x0d\x69\x74\x84\x37\x27\x65\x43\xe9\x2f\xab\xfb\x99\xc8\xfc\x3c\x7a\xcb\x9d\x99\xf2\x6e\xae\xff\x3c\xba\xe2\xac\xdd\x71\xaa\x93\xce\xca\xdd\x7a\x6d\xc7\xdb\xbc\xe2\x55\xae\x3b\x2b\x9b\xce\xad\xd1\xe4\x00\xeb\xe4\x6f\x62\x72\xf3\xec\xe3\x9a\x16\x33\xda\x3d\x4d\x63\x0c\xff\x29\xf0\x1b\x1b\xc7\xb4\x7a\x38\x66\x28\x3d\x9c\xc2\x1c\x4e\x61\x0e\xa7\x30\xd1\x50\x1d\x4e\x61\x9a\x18\x1d\x4e\x61\x0e\xa7\x30\xcf\x69\x0a\xc3\xfa\xad\x3d\x36\xd2\xcf\x67\x0e\xd1\xd9\xe9\xde\xff\xc6\xbd\xf5\xd8\x59\xbc\x65\xd7\xc6\x73\xaf\xd7\x57\x6f\xec\x7e\xf2\x1f\xbb\x5a\xb3\xb7\x46\xed\x9f\x26\xbd\x9d\x25\x77\xa9\xe2\x4e\x2d\xd8\xd5\x19\xf7\xfe\x96\x37\x7a\xc9\xbb\x39\x1b\xfe\x56\x4f\xd6\x21\x21\xe0\x0d\x43\x2b\x41\x1d\x04\x3f\x51\x64\x09\x48\xad\xb2\x4c\xb1\xd5\xce\x6a\xc8\x51\x60\xe2\x61\x64\x1d\x09\xe9\x00\xd0\xa2\x46\x3a\x82\xdf\x53\xb2\xff\x7a\x62\x04\x9f\x1b\xf4\x72\xa9\x38\x89\xd5\x10\x26\xd4\xe6\x4d\x55\x9c\xdb\x8f\xda\x38\x31\xfe\xf3\xc1\xa6\x59\xa7\x91\x45\xb0\x01\xd0\x79\x8a\x0c\x65\xff\x56\x92\xb2\x89\xe2\xb4\xc8\x76\xfc\xa3\x62\x69\xc6\xe0\x87\x9a\x81\xda\x87\x91\xa6\x16\x69\xd2\xc7\xd0\x95\xc8\x97\x57\x01\x67\x05\xff\x0f\x72\x5d\x3d\x29\x8d\x03\x8c\xd6\x0e\x7d\x53\x23\xda\xb4\x01\xd0\x1e\x51\xf2\x47\xc0\x39\x3a\x92\x36\x73\xbf\xc2\x59\x50\x13\xe7\x57\x41\x1b\x7f\x94\xba\x53\x0b\xee\xfd\xad\x36\xf0\x6a\xa8\x21\x20\xe7\xa5\xa3\xde\xea\xba\xb3\x79\xc9\xae\xd6\x9c\xaf\x37\x9c\xaf\xd7\x9f\x6c\x5f\xca\x4b\x47\x7a\x9a\x9a\x5b\x44\x20\x70\x81\x7d\x40\x3a\x41\x49\x73\xff\xb0\x37\x73\xdb\xf8\xc6\xcc\xb8\x98\xe5\xde\xc3\x8a\xb7\x33\xe1\xde\xac\xec\x2e\x7d\xe7\x8c\x4f\x3e\xd9\x9e\xf6\x56\xd7\xeb\xd3\x9f\x79\x37\x67\xdd\xab\x0b\xde\x0f\x9b\x76\x6d\xbc\xed\x48\xa6\xa1\xab\xeb\xce\xfc\x15\xc6\x22\xb6\x35\xcb\x18\x77\x72\xc9\x5d\xbc\xe1\xae\x7c\xeb\x4c\x8e\x7b\x57\x56\xbc\xef\xbf\xc8\xde\xc3\x9d\x9f\x73\x17\x6f\xd8\xd5\xfb\xc9\x3d\xe2\xa9\xe9\xe7\x1c\x18\x2e\x13\x6c\xa8\x2c\xf5\xd1\x70\x90\xa3\xed\x6d\xfc\x08\xb6\x1d\x05\x6d\xb9\xae\xae\x36\x0e\x3b\x67\xa7\x9a\xaa\xfa\x49\xc4\xd8\x79\x0e\xb7\xb7\x39\x6b\x77\xf8\x41\x64\x12\xc7\x1b\x12\x5c\x75\x47\x33\x2d\x73\xaf\xc7\x08\x3c\xbb\xb8\xc6\xe6\x1e\xa1\xcb\x01\x93\xef\x55\xc6\xa4\xf1\xd9\x86\x86\x07\xf6\xa3\x94\xc3\x09\xd2\x8b\x38\x41\xba\x20\xfc\x25\xd7\xe1\x04\x69\x5f\x70\x25\x12\xef\x70\x82\x14\x4e\x90\x5e\x7e\xe9\xe3\xff\x06\x00\x00\xff\xff\xc5\x28\x07\xbd\x36\x36\x00\x00")

func resourceExampleWebJsonBytes() ([]byte, error) {
	return bindataRead(
		_resourceExampleWebJson,
		"resource/example/web.json",
	)
}

func resourceExampleWebJson() (*asset, error) {
	bytes, err := resourceExampleWebJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resource/example/web.json", size: 13878, mode: os.FileMode(438), modTime: time.Unix(1663060963, 0)}
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
	"resource/example/client.json": resourceExampleClientJson,
	"resource/example/web.json":    resourceExampleWebJson,
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
	"resource": &bintree{nil, map[string]*bintree{
		"example": &bintree{nil, map[string]*bintree{
			"client.json": &bintree{resourceExampleClientJson, map[string]*bintree{}},
			"web.json":    &bintree{resourceExampleWebJson, map[string]*bintree{}},
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
