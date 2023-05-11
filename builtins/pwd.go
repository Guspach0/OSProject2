package pwd // package declaration

import (
	"fmt"
	"os"
  "path/filepath"
)

func Pwd(options []string) { // prints the  pathname of the current working directory
  
  var dir string
  var dir error
  
  if containsOption(options, "-P") { // the pathname printed will not contain symbolic links for -p
		dir, err = os.Getwd()
	} else if containsOption(options, "-L") { // pathname printed may contain symbolic links. 
		dir, err = filepath.EvalSymlinks(".")
	} else {
		dir, err = filepath.Abs(".")
	}
  
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(dir)
}

func containsOption(options []string, option string) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}
	return false
}
