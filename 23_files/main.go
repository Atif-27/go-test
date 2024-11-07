package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	deleteFile()
}
func fileDesc(){
	f,err:=os.Open("./example.txt")
	defer f.Close()
	if err!=nil{
		 panic(err)
	}
	//! File info
	fi, err:= f.Stat()
	fmt.Println("File name: ", fi.Name())
	fmt.Println("Is it a Directory? : ", fi.IsDir())
	fmt.Println("File Size: ", fi.Size())
	fmt.Println("File Permission: ", fi.Mode())
	fmt.Println("File Modified At: ", fi.ModTime())
}
func readFile(){
	f,err:= os.ReadFile("./example.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(f))
} 
func readFile2(){
	file, err:= os.Open("./example.txt")
	defer file.Close()
	if err!=nil{
		panic(err)
	}
	stat, err:=file.Stat()
	if err!=nil{
		panic(err)
	}
	buffer:= make([]byte,stat.Size())
	len,err:=file.Read(buffer)
	if err!=nil{
		panic(err)
	}
	fmt.Println("No of Bytes Read", len)
	fmt.Print("File Content: ")
	fmt.Print(string(buffer))
}

func readFolder(){
	dir, err:= os.Open("../")
	if err!=nil{
		panic(err)
	}
	defer dir.Close()

	fileInfo, err:= dir.ReadDir(0)
	for _,content:= range fileInfo{
		fmt.Println(content)
	}
}

func createFile(){
	f,err:= os.Create("test.txt")
	if err!=nil{
		panic(err)
	}
	defer f.Close()
	//! Here we are writing to the file, all the writeString are appended together
	f.WriteString("GO is opensource language")
	f.WriteString("GO is also known as golang")
	f.WriteString("golang")
}
func createFile2(){
	f,err:= os.Create("test.txt")
	if err!=nil{
		panic(err)
	}
	defer f.Close()
	//! Here we are writing to the file using a byte slic
	buffer:= []byte("Hello Golang")
	n,err:=f.Write(buffer)
	fmt.Println("Bytes written",n)
	if err!=nil{
		panic(err)
	}
}

// example to example2
func readWriteStream(){
	sourceFile, err:= os.Open("./example.txt")
	defer sourceFile.Close()
	if err!=nil{
		panic(err)
	}
	destFile,err:= os.Create("./example2.txt")
	defer destFile.Close()
	if err!=nil{
		panic(err)
	}

	reader:= bufio.NewReader(sourceFile)
	writer:= bufio.NewWriter(destFile)
	
	for{
		b,err:= reader.ReadByte()
		if err!=nil{
			if(err.Error()!="EOF"){
				panic(err)
			}
			break
		}
		e:=writer.WriteByte(b)
		if e!=nil{
			panic(e)
		}
	}
	writer.Flush()
	fmt.Println("Written")
}

func deleteFile(){
	err:=os.Remove("example2.txt")
	if err!=nil{
		panic(err)
	}
	fmt.Println("File Deleted")
}


