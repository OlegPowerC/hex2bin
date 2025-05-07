package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	hfname := flag.String("hex", "", "filename with hex data")
	bname := flag.String("bin", "", "filename with binary data")
	flag.Parse()

	if len(*hfname) < 3 {
		fmt.Println("please provide hex filename")
		return
	}
	if len(*bname) < 3 {
		fmt.Println("please provide binary filename")
		return
	}
	hexstring := ""
	filepatchshtype := "/"
	if runtime.GOOS == "windows" {
		filepatchshtype = "\\"
	}

	exfullpath, exerr := os.Executable()
	if exerr != nil {
		panic(any(exerr))
	}

	exPath := filepath.Dir(exfullpath)
	_, err := os.Stat(exPath + filepatchshtype + *hfname)
	if os.IsNotExist(err) {
		fmt.Println("hex file not found")
		return
	}
	f, fopenerr := os.Open(exPath + filepatchshtype + *hfname)
	if fopenerr != nil {
		panic(any(fopenerr))
	}
	defer f.Close()
	lb, _ := io.ReadAll(f)
	hexstring = string(lb)
	hexstring = strings.Replace(hexstring, "\r", "", -1)
	hexstring = strings.Replace(hexstring, "\n", "", -1)
	bdata, berr := hex.DecodeString(hexstring)
	if berr != nil {
		panic(any(berr))
	}
	fmt.Println("Write file:", exPath+filepatchshtype+*bname)
	werr := os.WriteFile(*bname, bdata, 0644)
	if werr != nil {
		panic(any(werr))
	}
}
