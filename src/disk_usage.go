package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
)

// Dir - A data structure to hold a dir/size pair.
type Dir struct {
	Name string
	Size uint64
}

// File struct containing the specific file info we are looking for
type File struct {
	ModTime time.Time
	Name    string
	Path    string
	Size    uint64
}

// SortedDirs - A slice of Dirs that implements sort.Interface to sort by Size.
type SortedDirs []Dir

// A function to turn a map into a PairList, then sort and return it.
func sortDirsBySize(m map[string]uint64) SortedDirs {
	d := make(SortedDirs, 0, len(m))
	for k, v := range m {
		d = append(d, Dir{k, v})
	}

	sort.Slice(d, func(i, j int) bool { return d[i].Size > d[j].Size })
	return d
}

// PartitionSpace returns total and free bytes available in a directory, e.g. `/`.
func PartitionSpace(path string) (m map[string]uint64, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return
	}
	m = make(map[string]uint64)
	m["total"] = uint64(int(s.Bsize) * int(s.Blocks))
	m["free"] = uint64(int(s.Bsize) * int(s.Bfree))
	m["inodes"] = uint64(s.Files)
	m["inodesFree"] = uint64(s.Ffree)
	return
}

// Percent calculate [number1] is what percent of [number2
func Percent(current uint64, all uint64) float64 {
	percent := (float64(current) * float64(100)) / float64(all)
	return percent
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must provide a single location to start from.")
		os.Exit(1)
	}
	start := os.Args[1]
	// Check if the `start` dir actually exists:
	fileInfo, err := os.Stat(start)
	if os.IsNotExist(err) {
		fmt.Printf("%s: No such file or directory\n", start)
		os.Exit(1)
	}
	// Check if it's a directory:
	if !fileInfo.IsDir() {
		fmt.Println("Start location must be a directory (not a file)")
		os.Exit(1)
	}
	totalFiles := 0
	files := make([]File, 0, 1024)
	dirs := make(map[string]uint64)
	filepath.Walk(start, func(file_path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs[file_path] = 0
		} else {
			dirs[path.Dir(file_path)] += uint64(info.Size())
			// Verify is not a Symlink
			if info.Mode()&os.ModeSymlink == 0 {
				totalFiles++
				file := File{
					ModTime: info.ModTime(),
					Name:    info.Name(),
					Path:    file_path,
					Size:    uint64(info.Size()),
				}
				// Keep files array to length of 20, re-sort of file add to full array
				if len(files) < 20 {
					files = append(files, file)
				} else {
					if file.Size > files[len(files)-1].Size {
						files[len(files)-1] = file
						sort.Slice(files, func(i, j int) bool { return files[i].Size > files[j].Size })
					}
				}
			}
		}
		return nil
	})
	p, err := PartitionSpace(start)
	if err != nil {
		fmt.Println(err)
	}

	sortedDirs := sortDirsBySize(dirs)
	if len(dirs) > 10 {
		sortedDirs = sortedDirs[:10]
	}

	// Output to stdout
	fmt.Printf("%.2f%% available disk space on %v\n", Percent(p["free"], p["total"]), start)
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", humanize.Bytes(p["total"]), humanize.Bytes((p["total"] - p["free"])), humanize.Bytes(p["free"]))

	fmt.Printf("\n%.2f%% of total inodes are free.\n", Percent(p["inodesFree"], p["inodes"]))
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", p["inodes"], (p["inodes"] - p["inodesFree"]), p["inodesFree"])

	fmt.Printf("\nTotal directory count of %d\n", len(dirs))
	fmt.Printf("The %v largest directories are:\n", len(sortedDirs))
	fmt.Println("Size   Directory")
	for _, dir := range sortedDirs {
		fmt.Printf("%-6v %v\n", humanize.Bytes(dir.Size), dir.Name)
	}

	fmt.Printf("\nTotal file count of %d\n", totalFiles)
	fmt.Printf("The %v largest files are:\n", len(files))
	fmt.Println("Size   Modified                      File")
	for _, file := range files {
		fmt.Printf("%-6v %v %v\n", humanize.Bytes(file.Size), file.ModTime, file.Path)
	}
}
