package device

import (
	"aoc2022/pkg/io"
	"strconv"
	"strings"
)

// Magic value to identify directory entries (i.e. size to be calculated)
const SIZE_DIRECTORY = -1

type FileSystemEntry struct {
	name     string
	size     int
	contents map[string]*FileSystemEntry
	parent   *FileSystemEntry
}

type FileSystem struct {
	working_directory *FileSystemEntry
	root              FileSystemEntry
}

func (fse *FileSystemEntry) GetPath() string {
	if fse.name == "" {
		return "/"
	}
	if fse.size == SIZE_DIRECTORY {
		return fse.parent.GetPath() + fse.name + "/"
	}
	return fse.parent.GetPath() + fse.name
}

func (fse *FileSystemEntry) GetSize() int {
	if fse.size != SIZE_DIRECTORY {
		return fse.size
	}
	size := 0
	for _, e := range fse.contents {
		size += e.GetSize()
	}
	return size
}

func (fs *FileSystem) GetSize() int {
	return fs.root.GetSize()
}

func (fse *FileSystemEntry) GetDeletionCandidate(req int) *FileSystemEntry {
	// Files aren't candidates
	if fse.size != SIZE_DIRECTORY {
		return nil
	}
	s := fse.GetSize()
	best := 0
	var best_e *FileSystemEntry
	// Directory as a whole is a valid candidate?
	if s >= req {
		best = s
		best_e = fse
	}
	for _, e := range fse.contents {
		// Ask children for their best candidate
		c := e.GetDeletionCandidate(req)
		if c == nil {
			continue
		}
		s := c.GetSize()
		// Is it better?
		if s >= req && (best == 0 || s < best) {
			best = s
			best_e = c
		}
	}
	return best_e
}

func (fs *FileSystem) GetDeletionCandidate(req int) *FileSystemEntry {
	return fs.root.GetDeletionCandidate(req)
}

func (fse *FileSystemEntry) GetSizeSum(limit int) int {
	if fse.size != SIZE_DIRECTORY {
		return 0
	}
	size1 := 0
	for _, e := range fse.contents {
		size1 += e.GetSizeSum(limit)
	}
	size2 := fse.GetSize()
	if size2 <= limit {
		return size1 + size2
	}
	return size1
}

func (fs *FileSystem) GetSizeSum(limit int) int {
	return fs.root.GetSizeSum(limit)
}

func NewFileSystem() *FileSystem {
	var fs FileSystem
	fs.root.size = SIZE_DIRECTORY
	fs.root.contents = make(map[string]*FileSystemEntry)
	return &fs
}

func ReadFileSystem(file string) *FileSystem {
	fs := NewFileSystem()
	lines := io.ReadLines(file)
	for i, line := range lines {
		switch line[0] {
		case '$': // Command
			if line[2:4] == "cd" {
				if line[5] == '/' {
					fs.working_directory = &fs.root
				} else if line[5] == '.' {
					fs.working_directory = fs.working_directory.parent
				} else {
					fs.working_directory = fs.working_directory.contents[line[5:]]
				}
			} else if line[2:4] == "ls" {
				for j := i + 1; j < len(lines) && lines[j][0] != '$'; j++ {
					data := strings.SplitN(lines[j], " ", 2)
					if data[0] == "dir" {
						fs.working_directory.contents[data[1]] = &FileSystemEntry{name: data[1], size: SIZE_DIRECTORY, contents: make(map[string]*FileSystemEntry), parent: fs.working_directory}
					} else {
						s, _ := strconv.Atoi(data[0])
						fs.working_directory.contents[data[1]] = &FileSystemEntry{name: data[1], size: s, parent: fs.working_directory}
					}
				}
			}
		}
	}
	return fs
}
