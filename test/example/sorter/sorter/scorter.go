package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
)

var infile = flag.String("i","infile","File contains values for sorting.")
var outfile = flag.String("o","outfile","File to receive sorted values.")
var algorithm = flag.String("a","algorithm","Sort algorithm")

func main() {

	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ",*infile,"outfile = ",*outfile,"algorithm = ", *algorithm)
	}
	values , err := read(*infile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//todo call aligorthm func
	switch *algorithm {
		case "bubblesort":
		case "qsort":
		default:
			fmt.Println("")
			return
	}

	if err := write(values,*outfile); err != nil {
		fmt.Println("write fail : ", err.Error())
		return
	}

}

// read file
func read(infile string) (values []int,err error) {
	file,err := os.Open(infile)
	if err != nil {
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		line, isPrefix, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return values,err
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line , seems unexpectecd .")
			return values,err
		}
		str := string(line)
		value,err := strconv.Atoi(str)
		if err != nil {
			return values,err
		}
		values = append(values,value)
	}
	return values,nil
}

// write file
func write(values []int , outfile string) error {
	file,err := os.Create(outfile)
	if err != nil {
		return err
	}
	//延迟关闭文件句柄
	defer file.Close()
	for _,v := range values {
		str := strconv.Itoa(v)
		file.WriteString(str + "\n")
	}
}