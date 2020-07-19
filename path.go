package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func separateFileName(filename string) (name, ext string) {
	ext = filepath.Ext(filename)
	base := filepath.Base(filename)
	name = strings.Replace(base, ext, "", 1)
	return
}

type path struct {
	dir       string
	name      string
	extention string
}

func newPath(v string) (path, error) {
	stat, err := os.Stat(v)
	if err != nil {
		return path{}, err
	}
	if stat.IsDir() {
		return path{dir: filepath.Dir(v)}, nil
	}
	if !stat.Mode().IsRegular() {
		return path{}, errBrokenFile
	}
	name, ext := separateFileName(v)
	return path{dir: filepath.Dir(v), name: name, extention: ext}, nil
}

func (p path) path() string {
	return p.dir + "/" + p.name + p.extention
}

func (p path) filename() string {
	return p.name + p.extention
}

func (p *path) copyFilename(name string) {
	p.name = name + "_duplicated"
}

func (p path) setName(n string) path {
	p.name = n
	return p
}

func (p path) addInt(i int) path {
	if i == 0 {
		return p
	}
	p.name += strconv.Itoa(i)
	return p
}

type src struct {
	path
}

func newSrc(path string) (src, error) {
	p, err := newPath(path)
	if err != nil {
		return src{}, err
	}
	return src{p}, nil
}

type dst struct {
	path
}

func newDst(s src, dir, filename string) (dst, error) {
	path, err := newPath(dir)
	if err != nil {
		return dst{}, err
	}
	if filename == "" || filename == s.filename() {
		path.copyFilename(s.name)
		path.extention = s.extention
		return dst{path}, nil
	}
	if filepath.Ext(filename) == "" {
		path.name = filename
		path.extention = s.extention
		return dst{path}, nil
	}
	name, ext := separateFileName(filename)
	path.name = name
	path.extention = ext
	return dst{path}, nil
}