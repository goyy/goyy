// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package files implements file utility functions.
package files

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/uuids"
)

// IsExist reports whether the specified file exists.
// Returns true if the file exists, false if it does not exist.
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// Read reads the file named by filename and returns the contents.
// A successful call returns err == nil, not err == EOF. Because Read
// reads the whole file, it does not treat an EOF from Read as an error
// to be reported.
func Read(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Write writes data to a file named by filename.
// If the file does not exist, Write creates it with permissions perm;
// otherwise Write truncates it before writing.
func Write(filename string, data string, perm os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), perm)
}

// Create creates the named file mode 0666 (before umask), truncating
// it if it already exists.  If successful, methods on the returned
// File can be used for I/O; the associated file descriptor has mode
// O_RDWR.
// If there is an error, it will be of type *PathError.
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Copy copies from src to dst until either EOF is reached on src or an error occurs.
func Copy(dstfile string, srcfile string, perm os.FileMode) error {
	data, rerr := ioutil.ReadFile(srcfile)
	if rerr != nil {
		return rerr
	}
	werr := ioutil.WriteFile(dstfile, data, perm)
	if werr != nil {
		return werr
	}
	return nil
}

// Mkdir creates a new directory with the specified name and permission bits.
// If there is an error, it will be of type *PathError.
func Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

// MkdirAll creates a directory named path, along with any necessary parents,
// and returns nil, or else returns an error.
// The permission bits perm are used for all directories that MkdirAll creates.
// If path is already a directory, MkdirAll does nothing and returns nil.
func MkdirAll(name string, perm os.FileMode) error {
	return os.MkdirAll(name, perm)
}

// Rename renames (moves) oldpath to newpath.
// If newpath already exists, Rename replaces it.
// OS-specific restrictions may apply when oldpath and newpath are in different directories.
// If there is an error, it will be of type *LinkError.
func Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Remove removes the named file or directory.
// If there is an error, it will be of type *PathError.
func Remove(name string) error {
	return os.Remove(name)
}

// LookPath searches for an executable binary named file
// in the directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
func Dir(file string) string {
	return filepath.Dir(file)
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func AbsDir(file string) (string, error) {
	dir, err := filepath.Abs(filepath.Dir(file))
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}

// Abs returns an absolute representation of path.
// If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path. The absolute
// path name for a given file is not guaranteed to be unique.
// Abs calls Clean on the result.
func Abs(path string) (string, error) {
	return filepath.Abs(path)
}

// Extension returns the <a href="http://en.wikipedia.org/wiki/Filename_extension">file
// extension</a> for the given file name, or the empty string if the file has
// no extension.  The result does not include the '{@code .}'.
func Extension(fileName string) string {
	fname := strings.Replace(fileName, "\\", "/", -1)
	ext := strings.AfterLast(fname, ".")
	first := strings.Left(ext, 1)
	if first == "/" {
		return ""
	}
	return ext
}

// ModTime returns the file modification time
func ModTime(file string) (out time.Time, err error) {
	f, err := os.Stat(file)
	if err != nil {
		return
	}
	return f.ModTime(), nil
}

// ModTimeUnix returns the file modification time
func ModTimeUnix(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// Size returns the length in bytes for regular files
func Size(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Join joins any number of path elements into a single path, adding
// a Separator if necessary. Join calls Clean on the result; in particular,
// all empty strings are ignored.
// On Windows, the result is a UNC path if and only if the first path
// element is a UNC path.
func Join(elem ...string) string {
	return filepath.Join(elem...)
}

// Zip compresses and archives a files.
func Zip(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// Unzip decompresses and unarchives a ZIP archive.
func Unzip(zipFile, destDir string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	MkdirAll(destDir, 0755)
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		filename := filepath.Join(destDir, file.Name)
		if file.FileInfo().IsDir() {
			err = MkdirAll(filename, file.Mode())
			if err != nil {
				rc.Close()
				return err
			}
		} else {
			w, err := os.Create(filename)
			if err != nil {
				rc.Close()
				return err
			}
			_, err = io.Copy(w, rc)
			if err != nil {
				w.Close()
				rc.Close()
				return err
			}
			w.Close()
		}
		rc.Close()
	}
	return nil
}

// Upload uploads a file.
// eg.
// confdir : /assets
// filedir : /upl
// field : art/library
// out : /assets/upl/art/library/8826e07b8d2111e6ab7854e873b1560c.png
func Upload(w http.ResponseWriter, r *http.Request, field, confdir, filedir string) (out string, err error) {
	if r.Method != "POST" {
		err = errors.New("status method tot allowed")
		logger.Error(err.Error())
		return
	}
	file, handler, err := r.FormFile(field)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	dir := parseDirLR(confdir) + parseDirRight(filedir)
	if !IsExist(dir) {
		if err = MkdirAll(dir, 0751); err != nil {
			logger.Error(err.Error())
			return
		}
	}
	filename := uuids.New() + "." + Extension(handler.Filename)
	filepath := dir + filename
	err = ioutil.WriteFile(filepath, data, 0755)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	out = parseDirLR(filedir) + filename
	return
}

// Uploads uploads a multipart file.
// eg.
// confdir : /assets
// filedir : /upl
// field : art/library
// out : /assets/upl/art/library/8826e07b8d2111e6ab7854e873b1560c.png
func Uploads(w http.ResponseWriter, r *http.Request, field, confdir, filedir string) (out []string, err error) {
	if r.Method != "POST" {
		err = errors.New("status method tot allowed")
		logger.Error(err.Error())
		return
	}
	//parse the multipart form in the request
	err = r.ParseMultipartForm(100000)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	//get a ref to the parsed multipart form
	m := r.MultipartForm

	dir := parseDirLR(confdir) + parseDirRight(filedir)
	//get the *fileheaders
	files := m.File[field]
	for i := range files {
		//for each fileheader, get a handle to the actual file
		file, ferr := files[i].Open()
		defer file.Close()
		if err != nil {
			err = ferr
			logger.Error(err.Error())
			return
		}
		if !IsExist(dir) {
			if err = MkdirAll(dir, 0751); err != nil {
				logger.Error(err.Error())
				return
			}
		}
		filename := uuids.New() + "." + Extension(files[i].Filename)
		filepath := dir + filename
		//create destination file making sure the path is writeable.
		dst, derr := os.Create(filepath)
		defer dst.Close()
		if err != nil {
			err = derr
			logger.Error(err.Error())
			return
		}
		//copy the uploaded file to the destination file
		if _, err = io.Copy(dst, file); err != nil {
			logger.Error(err.Error())
			return
		}
		out = append(out, parseDirLR(filedir)+filename)
	}
	return
}

func parseDirLeft(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	if strings.Left(dir, 1) != "/" {
		dir = "/" + dir
	}
	if strings.Right(dir, 1) == "/" {
		dir = dir[:len(dir)-1]
	}
	return dir
}

func parseDirRight(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	if strings.Left(dir, 1) == "/" {
		dir = dir[1:]
	}
	if strings.Right(dir, 1) != "/" {
		dir = dir + "/"
	}
	return dir
}

func parseDirLR(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	if strings.Left(dir, 1) != "/" {
		dir = "/" + dir
	}
	if strings.Right(dir, 1) != "/" {
		dir = dir + "/"
	}
	return dir
}
