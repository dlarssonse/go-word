// Go-Word public
package goword

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// ZipArchive ...
type ZipArchive struct {
	Files []ZipFile

	Reader *zip.ReadCloser
}

// ZipFile ...
type ZipFile struct {
	File     *zip.File
	Filename string
	Replace  bool
}

// Open ...
func Open(filename string) (*ZipArchive, error) {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	archive := ZipArchive{Reader: reader}
	for _, file := range archive.Reader.File {
		archive.Files = append(archive.Files, ZipFile{File: file})
	}

	return &archive, nil
}

// GetFileFromReader ...
func (archive *ZipArchive) GetFileFromReader(name string) (file *zip.File, err error) {
	for _, f := range archive.Files {
		if f.File.Name == name {
			return f.File, nil
		}
	}
	return nil, fmt.Errorf("no such file '%s'", name)
}

// ReplaceFile ...
func (archive *ZipArchive) ReplaceFile(sourceFilename string, destinationFilename string) error {

	for index, f := range archive.Files {
		if f.File.Name == sourceFilename {

			archive.Files[index].Replace = true
			archive.Files[index].Filename = destinationFilename

			return nil
		}
	}
	return fmt.Errorf("no such file '%s'", sourceFilename)
}

// SaveAs ...
func (archive *ZipArchive) SaveAs(filename string) error {

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := zip.NewWriter(output)

	for _, file := range archive.Files {
		fout, err := writer.Create(file.File.Name)
		if err != nil {
			return err
		}

		if file.Replace {
			srcFile, err := os.Open(file.Filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(fout, srcFile)
			if err != nil {
				return err
			}
		} else {
			fin, err := file.File.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(fout, fin)
			if err != nil {
				return err
			}
		}
	}

	return writer.Close()
}
