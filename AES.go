package main

import (
	"crypto/aes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	file := "_00000001_20240412-160714_1130h85.mp4"
	key := "admin1"

	cmd := exec.Command("ffmpeg",
		"-i",
		file,
		"-decryption_key",
		string(putKeyInHash(key)),
		strings.Split(file, ".")[0]+"_decrypted.mp4")

	fmt.Println(cmd)

	Decrypt, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(Decrypt)
}

func putKeyInHash(k string) []byte {
	key := []byte(k)
	hash := make([]byte, 32)
	for i, v := range key {
		hash[i] = v
	}
	return hash
}

func Decrypt(k string, file string) {
	t1 := time.Now()

	//key := "e00cf25ad42683b3df678c61f42c6bda" //keyMD5
	hash := putKeyInHash(k)

	fmt.Printf("%s", hash)

	video, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer video.Close()
	info, _ := video.Stat()
	buffer := make([]byte, info.Size())
	n, err := video.Read(buffer)
	if n != len(buffer) {
		panic("different n")
	} else if err != nil {
		panic(err)
	}

	newFile := DecryptAES(hash[:], buffer)
	nFile, _ := os.Create(strings.Split(file, ".")[0] + "_decrypted.mp4")
	defer nFile.Close()
	nFile.Write(newFile)

	fmt.Println("\ntime Elapsed: ", time.Since(t1))
}

func DecryptAES(key []byte, ct []byte) []byte {
	//ciphertext, _ := hex.DecodeString(ct)
	ciphertext := ct
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	return pt
}
