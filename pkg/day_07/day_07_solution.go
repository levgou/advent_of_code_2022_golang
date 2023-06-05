package day_07

import (
	"advent_of_code/pkg/shared"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strconv"
	"strings"
)

const (
	DemoInput = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_07/demo_input.txt"
	Input     = "/Users/stavalaluf/lev_tmp/advent_of_code_2022_golang/pkg/day_07/input.txt"
)

type LsFile struct {
	IsDir bool
	Size  int
	Name  string
}

type FILE struct {
	IsDir    bool
	Size     int
	Children []FILE
	Name     string
	Parent   *FILE
}

type CMD struct {
	IsCD        bool
	Target      string
	FilesOutput []LsFile
}

func parseCMD(line string) CMD {
	lineParts := strings.Split(line, " ")
	if lineParts[1] == "cd" {
		return CMD{
			IsCD:        true,
			Target:      lineParts[2],
			FilesOutput: []LsFile{},
		}
	} else {
		return CMD{
			IsCD:        false,
			Target:      "",
			FilesOutput: []LsFile{},
		}
	}

}

func parseFile(line string) LsFile {
	lineParts := strings.Split(line, " ")
	if strings.Contains(line, "dir") {
		return LsFile{
			IsDir: true,
			Size:  0,
			Name:  lineParts[1],
		}
	} else {
		size, _ := strconv.ParseInt(lineParts[0], 10, 64)
		return LsFile{
			IsDir: false,
			Size:  int(size),
			Name:  lineParts[1],
		}
	}
}

func parse(lines []string) []CMD {
	cmds := []CMD{}

	for _, line := range lines {
		if line[0] == '$' {
			cmds = append(cmds, parseCMD(line))
		} else {
			// assume last cmd was ls
			ls := shared.LastRef(cmds)
			ls.FilesOutput = append(ls.FilesOutput, parseFile(line))
		}
	}

	return cmds
}

func deviseTree(cmds []CMD) *FILE {
	root := FILE{
		IsDir:    true,
		Size:     0,
		Children: []FILE{},
		Name:     "/",
		Parent:   nil,
	}

	cwd := &root

	for _, cmd := range cmds[1:] {
		if !cmd.IsCD {
			for _, lsfile := range cmd.FilesOutput {
				file := FILE{
					IsDir:    lsfile.IsDir,
					Size:     lsfile.Size,
					Children: []FILE{},
					Name:     lsfile.Name,
				}
				cwd.Children = append(cwd.Children, file)
			}
		} else {
			if cmd.Target == ".." {
				cwd = cwd.Parent
			} else {
				nexdCwdIndex := slices.IndexFunc(cwd.Children, func(f FILE) bool {
					return f.IsDir && f.Name == cmd.Target
				})

				if nexdCwdIndex == -1 {
					fmt.Println(cmd)
					fmt.Println(cwd)
					panic("no such directory " + cmd.Target)
				}

				cwd.Children[nexdCwdIndex].Parent = cwd
				cwd = &cwd.Children[nexdCwdIndex]
			}
		}
	}

	return &root
}

func printTree(tree *FILE, indent string) {
	suffix := fmt.Sprintf("(DIR, %d)", tree.Size)
	if !tree.IsDir {
		suffix = fmt.Sprintf("(FILE, %d)", tree.Size)
	}

	fmt.Println(indent, "-", tree.Name, suffix)
	for _, child := range tree.Children {
		printTree(&child, indent+"  ")
	}
}

func updateDirSizes(tree *FILE) int {
	if tree.IsDir {
		tree.Size = 0
		for i := range tree.Children {
			tree.Size += updateDirSizes(&tree.Children[i])
		}
	}

	return tree.Size
}

func sumDirectoriesBellow100k(tree *FILE) int {
	if !tree.IsDir {
		return 0
	}

	sum := 0
	if tree.Size <= 100000 {
		sum += tree.Size
	}

	for i := range tree.Children {
		sum += sumDirectoriesBellow100k(&tree.Children[i])
	}

	return sum
}

const (
	DeviceSpace = 70000000
	NeededSpace = 30000000
)

func Solution() {
	lines, err := shared.ReadLines(Input)
	if err != nil {
		log.Fatal(err)
		return
	}

	parsedCmds := parse(lines)
	tree := deviseTree(parsedCmds)
	updateDirSizes(tree)

	printTree(tree, "")

	sum := sumDirectoriesBellow100k(tree)
	fmt.Println(sum)

	rootSize := tree.Size
	freeSpace := DeviceSpace - rootSize
	needToFree := NeededSpace - freeSpace

	fmt.Printf("Need to free %d\n", needToFree)
	directories := []*FILE{}

	var findDirs func(file *FILE)
	findDirs = func(file *FILE) {
		if file.IsDir {
			directories = append(directories, file)
		}

		for i := range file.Children {
			findDirs(&file.Children[i])
		}
	}

	findDirs(tree)
	slices.SortFunc(directories, func(i, j *FILE) bool {
		return i.Size < j.Size
	})

	for _, dir := range directories {
		if dir.Size >= needToFree {
			fmt.Println(dir.Name, dir.Size)
			break
		}
	}
}
