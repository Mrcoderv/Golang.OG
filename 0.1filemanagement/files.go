package main

import (
	"os"
)

// program kata bata run gareko cham bhanen le ni farak parne raicha

func main() {

	// data, err := os.Open("hello.txt") // Relative paths use the process working directory, not the source file's directory.
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := data.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File name :", fileInfo.Name())
	// fmt.Println("File size :", fileInfo.Size())
	// fmt.Println("File mode :", fileInfo.Mode())
	// fmt.Println("folder ", fileInfo.IsDir())
	// fmt.Println("File modification time :", fileInfo.ModTime())

	// 	data, err := os.ReadFile("hello.txt")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(data))
	// 	// defer data.Close()
	// 	// for the large file
	// 	// read folder
	// 	dir, err := os.Open(".")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer dir.Close()
	// 	filesinfo, err := dir.Readdir(-1)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	count := 0
	// 	for _, fileinfo := range filesinfo {
	// 		fmt.Println("File name :", fileinfo.Name())
	// 		fmt.Println("File size :", fileinfo.Size())
	// 		fmt.Println("File mode :", fileinfo.Mode())
	// 		fmt.Println("folder ", fileinfo.IsDir())

	// 		fmt.Println("File modification time :", fileinfo.ModTime())
	// 		fmt.Println("---------------------------------------------------")
	// count++
	// 	}
	// 	defer fmt.Println("Total files in the folder: ", count)

	// writing on the file
	// f, err := os.Create("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// f.Write([]byte("Hello, World!"))

	// // replacinfg the content of the file
	// f, err = os.Create("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// f.WriteString("Hello, World!")
	// f.WriteString("\nhello ")
	// f.WriteString("world")
	// byetes := []byte("\n Hello, World!")
	// f.Write(byetes)
	// copying  and pesting the content of the file
	// sourceFile, err := os.Open("hello.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()
	// destFile, err := os.Create("hello_copy.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()
	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)
	// for {
	// 	line, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break
	// 		}
	// 		panic(err)
	// 	}
	// 	e := writer.WriteByte(line)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }
	// defer writer.Flush()

	// deleting the file
	sourceFile, err := os.Open("hello_copy.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	err = os.Remove("hello_copy.txt")
	if err != nil {
		panic(err)
	}

}

// here the file is closed after the function is executed. The defer statement is used to ensure that the file is closed even if an error occurs while reading the file.
