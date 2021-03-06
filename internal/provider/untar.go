package provider

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Untar extracts a gzipped tarfile into the local filesystem
// In the event of an error, some files may still be written to disk, so any required cleanup should still take place
func Untar(dest string, reader io.Reader) error {
	absDest, err := filepath.Abs(dest)
	if err != nil {
		return err
	}
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()

		// We finished reading the file
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		} else if header == nil {
			continue
		}

		target, err := filepath.Abs(filepath.Join(absDest, header.Name))
		if err != nil {
			return nil
		}
		if !strings.HasPrefix(target, absDest) {
			return fmt.Errorf("tarball attempting to unpack to %s, which is outside of unpack dir", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			_, err := os.Stat(target)
			if os.IsNotExist(err) {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(file, tarReader); err != nil {
				return err
			}

			_ = file.Close()
		}
	}
}
