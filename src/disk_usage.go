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
func PartitionSpace(path string) (total uint64, free uint64, inodes uint64, inodesFree uint64, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return
	}
	total = uint64(int(s.Bsize) * int(s.Blocks))
	free = uint64(int(s.Bsize) * int(s.Bfree))
	inodes = uint64(s.Files)
	inodesFree = uint64(s.Ffree)
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

	files := make([]File, 0, 1024)
	dirs := make(map[string]uint64)
	filepath.Walk(start, func(file_path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs[file_path] = 0
		} else {
			dirs[path.Dir(file_path)] += uint64(info.Size())
			file := File{
				ModTime: info.ModTime(),
				Name:    info.Name(),
				Path:    file_path,
				Size:    uint64(info.Size()),
			}
			files = append(files, file)
		}
		return nil
	})
	total, free, inodes, inodesFree, error := PartitionSpace(start)
	if error != nil {
		fmt.Println(error)
	}

	sortedDirs := sortDirsBySize(dirs)
	if len(dirs) > 10 {
		sortedDirs = sortedDirs[:10]
	}
	// Sort the files array in descending order
	totalFiles := len(files)
	sort.Slice(files, func(i, j int) bool { return files[i].Size > files[j].Size })
	if len(files) > 20 {
		files = files[:20]
	}

	// Output to stdout
	fmt.Printf("%.2f%% available disk space on %v\n", Percent(free, total), start)
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", humanize.Bytes(total), humanize.Bytes((total - free)), humanize.Bytes(free))

	fmt.Printf("\n%.2f%% of total inodes are free.\n", Percent(inodesFree, inodes))
	fmt.Printf("Total: %v, Used: %v, Free: %v\n", inodes, (inodes - inodesFree), inodesFree)

	fmt.Printf("\nTotal directory count of %d\n", len(dirs))
	fmt.Println("The 10 largest directories are:")
	fmt.Println("Size   Directory")
	for _, dir := range sortedDirs {
		fmt.Printf("%-6v %v\n", humanize.Bytes(dir.Size), dir.Name)
	}

	fmt.Printf("\nTotal file count of %d\n", totalFiles)
	fmt.Printf("The %v largest files are:", len(files))
	fmt.Println("Size   Modified                      File")
	for _, file := range files {
		fmt.Printf("%-6v %v %v\n", humanize.Bytes(file.Size), file.ModTime, file.Path)
	}
}
