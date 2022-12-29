package main

import (
  "os"
  "fmt"
  "sort"
  "bufio"
  "strings"
  "strconv"
)

const (
  THRESHOLD = 100000
  TOTAL_SPACE = 70000000
  UPDATE_SPACE = 30000000
)

type File struct {
  filename string
  size int
}

func NewFile(fn string, sz int) File {
  return File{fn, sz}
}

type Directory struct {
  dirname string
  files []File
  dirs []*Directory
  parent *Directory
}

func NewDirectory(n string, p *Directory) *Directory {
  return &Directory{n, []File{}, []*Directory{}, p}
}

func (d *Directory) FullName() string {
  if d.dirname == "/" {
    return "/"
  }

  parentName := d.parent.FullName()
  if parentName == "/" {
    parentName = ""
  }
  return parentName + "/" + d.dirname
}

func (fs *Directory) GetDirectory(dirname string) *Directory {
  for _, d := range fs.dirs {
    if d.dirname == dirname {
      return d
    }
  }

  if fs.dirname == dirname {
    return fs
  }
  
  return nil
}

func parseCd(currDir *Directory, dir string) *Directory {
  if dir == ".." {
    return currDir.parent
  }
  return currDir.GetDirectory(dir)
}

func parseLs(dirContent []string, currDir *Directory) {
  for _, dc := range dirContent {
    stat := strings.Split(dc, " ")
    if stat[0] == "dir" {
      currDir.dirs = append(currDir.dirs, NewDirectory(stat[1], currDir))
    } else {
      sz, err := strconv.Atoi(stat[0])
      if err != nil {
        panic("Cannot convert file size")
      }
      currDir.files = append(currDir.files, NewFile(stat[1], sz))
    }
  }
}

func (fs *Directory) PrintFs(depth int) {
  if fs == nil {
    return
  }
  for i := 0; i < depth; i++ {
    fmt.Printf("-")
  }
  fmt.Printf(" %s (dir, %d files, %d dirs)\n", fs.FullName(), len(fs.files), len(fs.dirs))
  for _, d := range fs.dirs {
    d.PrintFs(depth+1)
  }
  for _, f := range fs.files {
    for i := 0; i < depth+1; i++ {
      fmt.Printf("-")
    }
    fmt.Printf(" %s (file, size=%d)\n", fs.FullName() + f.filename, f.size)
  }
}

func nextCmdIdx(lines []string, currIdx int) int {
  n := len(lines)
  for i := currIdx+1; i < n; i++ {
    if string(lines[i][0]) == "$" {
      return i
    }
  }
  return n
}

func parseInput(f *os.File) *Directory {
  scanner := bufio.NewScanner(f)
  lines := []string{}
  fs := NewDirectory("/", nil)
  currDir := fs.GetDirectory("/")
  
  for (scanner.Scan()) {
    line := scanner.Text()
    lines = append(lines, line)
  }

  for i, line := range lines {
    fields := strings.Split(line, " ")
    if fields[0] == "$" {
      switch fields[1] {
      case "ls":
        parseLs(lines[i+1:nextCmdIdx(lines, i)], currDir)
      case "cd":
        if len(fields) != 3 {
          panic("cd command not valid")
        }
        currDir = parseCd(currDir, fields[2])
      default:
        panic("Command not known")
      }
    }
  }

  return fs
}

func (fs *Directory) GetDirsSize(sizes *[]int) int {
  if fs == nil {
    return 0
  }
  currSize := 0
  for _, f := range fs.files {
    currSize += f.size
  }
  for _, d := range fs.dirs {
    currSize += d.GetDirsSize(sizes)
  }
  (*sizes) = append(*sizes, currSize)
  return currSize
}

func part1(fs *Directory) {
  var dirsSizes []int
  fs.GetDirsSize(&dirsSizes)
  
  var spaceToFree int = 0
  for _, sz := range dirsSizes {
    if sz <= THRESHOLD {
      spaceToFree += sz
    }
  }
  fmt.Println("The space available to free is:", spaceToFree, "bytes")
}

func part2(fs *Directory) {
  var dirsSizes []int
  diskUsed := fs.GetDirsSize(&dirsSizes)
  diskNeeded := UPDATE_SPACE - (TOTAL_SPACE - diskUsed)
  sort.Ints(dirsSizes)

  for _, sz := range dirsSizes {
    if sz >= diskNeeded {
      fmt.Println("I can free", sz, "bytes")
      break
    }
  }
}

func main() {
  f, err := os.Open("input")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  fs := parseInput(f)
  // fs.PrintFs(1)

  part1(fs)
  part2(fs)
}

