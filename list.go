/*
 * This file is part of PathsHelper library.
 *
 * Copyright 2018 Arduino AG (http://www.arduino.cc/)
 *
 * PropertiesMap library is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 */

package paths

import (
	"strings"
)

// PathList is a list of Path
type PathList []*Path

// NewPathList creates a new PathList with the given paths
func NewPathList(paths ...string) PathList {
	res := PathList{}
	for _, path := range paths {
		res = append(res, New(path))
	}
	return res
}

// Clone returns a copy of the current PathList
func (p *PathList) Clone() PathList {
	res := PathList{}
	for _, path := range *p {
		res.Add(path.Clone())
	}
	return res
}

// AsStrings return this path list as a string array
func (p *PathList) AsStrings() []string {
	res := []string{}
	for _, path := range *p {
		res = append(res, path.String())
	}
	return res
}

// FilterDirs remove all entries except directories
func (p *PathList) FilterDirs() {
	res := (*p)[:0]
	for _, path := range *p {
		if isDir, _ := path.IsDir(); isDir {
			res = append(res, path)
		}
	}
	*p = res
}

// FilterOutHiddenFiles remove all hidden files (files with the name
// starting with ".")
func (p *PathList) FilterOutHiddenFiles() {
	p.FilterOutPrefix(".")
}

// FilterOutPrefix remove all entries having the specified prefix
func (p *PathList) FilterOutPrefix(prefix string) {
	res := (*p)[:0]
	for _, path := range *p {
		if !strings.HasPrefix(path.Base(), prefix) {
			res = append(res, path)
		}
	}
	*p = res
}

// Add adds a Path to the PathList
func (p *PathList) Add(path *Path) {
	*p = append(*p, path)
}

// AddAll adds all Paths in the list passed as argument
func (p *PathList) AddAll(paths PathList) {
	for _, path := range paths {
		*p = append(*p, path)
	}
}

// AddIfMissing adds a Path to the PathList if the path is not already
// in the list
func (p *PathList) AddIfMissing(path *Path) {
	if (*p).Contains(path) {
		return
	}
	(*p).Add(path)
}

// AddAllMissing adds all paths to the PathList excluding the path already
// in the list
func (p *PathList) AddAllMissing(paths PathList) {
	for _, path := range *p {
		(*p).AddIfMissing(path)
	}
}

// ToAbs calls Path.ToAbs() method on each path of the list.
// It stops at the first error and returns it. If all ToAbs calls
// are successful nil is returned.
func (p *PathList) ToAbs() error {
	for _, path := range *p {
		if err := path.ToAbs(); err != nil {
			return err
		}
	}
	return nil
}

// Contains check if the list contains a path that match
// exactly (EqualsTo) to the specified path
func (p *PathList) Contains(path *Path) bool {
	for _, path := range *p {
		if path.EqualsTo(path) {
			return true
		}
	}
	return false
}

// ContainsEquivalentTo check if the list contains a path
// that is equivalent (EquivalentTo) to the specified path
func (p *PathList) ContainsEquivalentTo(path *Path) bool {
	for _, path := range *p {
		if path.EquivalentTo(path) {
			return true
		}
	}
	return false
}
