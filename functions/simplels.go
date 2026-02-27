package functions
import "io"
import "os"
import "fmt"
import "sort"
import "path/filepath"

func SimpleLS( w io.Writer, args []string, useColor bool) {

	var files []string
	var dirs []string
	for _, path := range args {
		info, err := os.Lstat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot access '%s' : %v\n", path, err)
			continue
		}
		if info.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
	}
	sort.Strings(files)
	sort.Strings(dirs)

	for _, f := range files {
		fmt.Fprintln(w, filepath.Base(f))
	}
	multipleDirs := len(dirs) > 1 || (len(files) > 0 && len(dirs) >0)
	for i, dir := range dirs {
		if multipleDirs {
			if i > 0 || len(files) > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s:\n", dir)
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "gols: cannot open directory '%s': %v\n", dir, err)
			continue
		}
		var names []string
		for _, entry := range entries {
			name := entry.Name()
			if len(name) > 0 && name[0] == '.' {
				continue
			}
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fullPath := filepath.Join(dir,name)
			info, err := os.Lstat(fullPath)
			if err != nil {
				continue
			}
			mode := info.Mode()

			isDir := info.IsDir()
			isExec := mode.IsRegular() && (mode&0111) != 0

			printEntry(w, name, isDir, isExec, useColor)
		}
	}

}
func printEntry(w io.Writer, name string, isDir bool, isExec bool, useColor bool) {
    if !useColor {
        fmt.Fprintln(w, name)
        return
    }
    if isDir {
        fmt.Fprintf(w, "\033[34m%s\033[0m\n", name)
        return
    }
    if isExec {
        fmt.Fprintf(w, "\033[32m%s\033[0m\n", name)
        return
    }
    fmt.Fprintln(w, name)
}
