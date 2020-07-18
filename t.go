// buffer := ""

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		buffer = buffer + scanner.Text() + "\n"

// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "reading standard input: ", err)
// 	}


// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"strings"

// 	// "io/ioutil"
// 	"log"

// 	"github.com/jung-kurt/gofpdf"
// )

// func main() {

// 	// creates input reader and reads from standard input
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter the file name: ")
// 	// reading a string from standard input including trailing new line
// 	file, _ := reader.ReadString('\n')
// 	// replacing last new line character with an empty string
// 	file = strings.Replace(file, "\n", "", -1)
// 	// reading file taking the content and dumping it into pdf library
// 	content, err := ioutil.ReadFile(file)
// 	// checks if an error exist and outputs error
// 	if err != nil {
// 		log.Fatalf("%s file not found", file)

// 	}
// 	// assigns pdf to gofpdf library parameters (page orientation, unit of measurment, pagesize, directory to some fonts)
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	// creates page in a pdf flile
// 	pdf.AddPage()
// 	// set fonts (font, style, size)
// 	pdf.SetFont("Arial", "B", 16)

// 	// puts content in the pdf file (width, line height, content, border, alignment, style )
// 	pdf.MultiCell(190, 5, string(content), "0", "0", false)

// 	// output the file and close it
// 	_ = pdf.OutputFileAndClose("test.pdf")

// 	fmt.Println("PDF Created")

// }