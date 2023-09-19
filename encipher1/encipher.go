package main

import (
	"bufio"
	"crypto/des"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	key := []byte{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("vim-go hoooi")
	//readFiles()
	//writeFiles()
	plaintext, err := os.ReadFile("hello.go.txt")
	check(err)

	//fmt.Print(string(dat)[len(dat)-100:])
	//fmt.Print("--------------")

	fmt.Println(len(plaintext))
	var num int
	num = len(plaintext) / 8
	fmt.Println("jumlah blok : ", num)

	block, err := des.NewCipher(key)
	check(err)

	ciphertext := make([]byte, len(plaintext))

	fmt.Println(plaintext[0:8])
	block.Encrypt(ciphertext, plaintext)
	fmt.Println(ciphertext[0:8])

}

func readFiles() {
	dat, err := os.ReadFile(`E:\Program_Files\progjar\enchiper\tmp\dat`)
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open(`E:\Program_Files\progjar\enchiper\tmp\dat`)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

func writeFiles() {

	d1 := []byte("hello\ngo\n")
	d1[0] = 'H'
	fmt.Println(d1)
	fmt.Println(string(d1))
	err := os.WriteFile(`E:\Program_Files\progjar\enchiper\tmp\dat`, d1, 0644)
	check(err)

	f, err := os.Create(`E:\Program_Files\progjar\enchiper\tmp\dat`)
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
