package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Car struct {
	Maker string
	Model string
	Type  string
}

var (
	fileName = flag.String("data", "./Data.csv", "Dosya Okundu")
	query    = flag.String("q", "suv", "Query Bilgi Mesajı")
)

func findMaker(key string, a []Car) {

	for i := 0; i < len(a); i++ {
		key = strings.ToLower(key)
		if key == strings.ToLower(a[i].Maker) || key == strings.ToLower(a[i].Model) || key == strings.ToLower(a[i].Type) {
			fmt.Println(a[i])
		}
	}
}

// ReadData Dosya Okuma Fonksiyonu
func ReadData(fileName string) ([]Car, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Hataa!")
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	var dizi []Car

	for scanner.Scan() {
		line := scanner.Text()
		col := strings.Split(line, ",")
		car := Car{Type: col[2], Maker: col[0], Model: col[1]}
		dizi = append(dizi, car)
	}
	return dizi, err
}
func main() {
	flag.Parse()

	data, err := ReadData(*fileName)
	if nil != err {
		log.Fatal("Data okunamadı", err)
	}

	findMaker(strings.ToLower(*query), data)

}
