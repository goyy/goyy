// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package files_test

import (
	"gopkg.in/goyy/goyy.v0/util/files"
	"os"
	"testing"
)

func TestIsExist(t *testing.T) {
	s := []struct {
		file     string
		expected bool
	}{
		{"./example.txt", true},
		{"./README", false},
	}
	for _, v := range s {
		if out := files.IsExist(v.file); out != v.expected {
			t.Errorf(`IsExist("%s") = %v, want %v`, v.file, out, v.expected)
		}
	}
}

func TestRead(t *testing.T) {
	filename := "rumpelstilzchen"
	contents, err := files.Read(filename)
	if err == nil {
		t.Fatalf("Read %s: error expected, none found", filename)
	}

	filename = "files.go"
	contents, err = files.Read(filename)
	if err != nil {
		t.Fatalf("Read %s: %v", filename, err)
	}

	checkSize(t, filename, int64(len(contents)))
}

func TestWrite(t *testing.T) {
	filename := "./example.txt"
	data := "Programming today is a race between software engineers striving to " +
		"build bigger and better idiot-proof programs, and the Universe trying " +
		"to produce bigger and better idiots. So far, the Universe is winning."

	if err := files.Write(filename, data, 0644); err != nil {
		t.Fatalf("Write %s: %v", filename, err)
	}

	contents, err := files.Read(filename)
	if err != nil {
		t.Fatalf("Read %s: %v", filename, err)
	}

	if contents != data {
		t.Fatalf("contents = %q\nexpected = %q", contents, data)
	}

	// recover
	files.Write(filename, "Hello world!", 0644)
}

func TestCopy(t *testing.T) {
	dst := "./copy.txt"
	src := "./example.txt"
	err := files.Copy(dst, src, 0644)
	if err != nil {
		t.Fatalf("Copy %s %s: %v", dst, src, err)
	}
	if files.IsExist(dst) == false {
		t.Errorf(`Copy("%s", "%s", %v) error`, dst, src, 0644)
	}
	files.Remove(dst)
}

func TestRename(t *testing.T) {
	from, to := "renamefrom", "renameto"
	files.Remove(to) // Just in case.
	file, err := files.Create(from)
	if err != nil {
		t.Fatalf("open %q failed: %v", to, err)
	}
	if err = file.Close(); err != nil {
		t.Errorf("close %q failed: %v", to, err)
	}
	err = files.Rename(from, to)
	if err != nil {
		t.Fatalf("rename %q, %q failed: %v", to, from, err)
	}
	defer files.Remove(to)
	_, err = os.Stat(to)
	if err != nil {
		t.Errorf("stat %q failed: %v", to, err)
	}
}

func TestRemove(t *testing.T) {
	filename := "./remove.txt"
	expected := false
	if err := files.Write(filename, "remove", 0644); err != nil {
		t.Fatalf("Remove:Write %s: %v", filename, err)
	}
	err := files.Remove(filename)
	if err != nil {
		t.Fatalf("Remove(%#q): %v", filename, err)
	}
	if out := files.IsExist(filename); out != expected {
		t.Errorf(`Remove("%s") = %v, want %v`, filename, out, expected)
	}
}

func TestGetExtension(t *testing.T) {
	s := []struct {
		file     string
		expected string
	}{
		{"./example.txt", "txt"},
		{"./sys/file.example.txt", "txt"},
		{"./README", ""},
	}
	for _, v := range s {
		if out := files.GetExtension(v.file); out != v.expected {
			t.Errorf(`GetExtension("%s") = %v, want %v`, v.file, out, v.expected)
		}
	}
}

func TestModTime(t *testing.T) {
	filename := "./README.md"
	if out, _ := files.ModTime(filename); out <= 0 {
		t.Errorf(`ModTime("%s") = %v, want %s`, filename, out, ">0")
	}
}

func TestSize(t *testing.T) {
	filename := "./README.md"
	expected := int64(750)
	if out, _ := files.Size(filename); out != expected {
		t.Errorf(`Size("%s") = %v, want %v`, filename, out, expected)
	}
}

func checkSize(t *testing.T, path string, size int64) {
	dir, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", path, size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, dir.Size(), size)
	}
}
