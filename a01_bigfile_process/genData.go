package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func genMd5Data(limit int, path string) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputWriter := bufio.NewWriter(outputFile)

	for i := 0; i < limit; i++ {
		if i != 0 {
			_, err = outputWriter.WriteString("\r\n")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		rd := time.Now().Nanosecond() + rand.Int()
		s := strconv.Itoa(rd)
		r := md5.Sum([]byte(s))
		line := hex.EncodeToString(r[:])
		_, err = outputWriter.WriteString(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	outputWriter.Flush()
	outputFile.Close()
}

func genRandNumData(limit int, path string) {
	outputFile, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	outputWriter := bufio.NewWriter(outputFile)

	for i := 0; i < limit; i++ {
		if i != 0 {
			_, err = outputWriter.WriteString("\r\n")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		min := 10000
		max := 100000
		rd := rand.Intn(max-min) + min
		line := strconv.Itoa(rd)
		_, err = outputWriter.WriteString(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	outputWriter.Flush()
	outputFile.Close()
}
