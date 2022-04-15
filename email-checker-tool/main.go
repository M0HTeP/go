package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)   //   создаем сканер
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")   //   выводим сообщение о том, какие данные будут получены припроверке

	for scanner.Scan() {
		checkDomain(scanner.Text())   //   функция скапнирует введенный текст
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n, err")    //   в случае ошибки - выводим ее
	}
}
   //   функция проверки домена
func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool   //   переключатели
	var spfRecord, dmarcRecord string   //   текст

	mxRecords, err := net.LookupMX(domain)   //   получение текста о домене, либо ошибки

	if err != nil {
		log.Printf("Error: %v\n", err)   //   вывод ошибки
	}
	if len(mxRecords) > 0 {
		hasMX = true   //   если присутствуют записи, переключаем в позицию True
	}

	txtRecords, err := net.LookupTXT(domain)   //   получаем записи о домене
	if err != nil {
		log.Printf("Error: %v\n", err)    //   вывод ошибки
	}

	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){   //   если есть записи об spf
			hasSPF = true   //   переключаем
			spfRecord = record   //   получаем записи в переменную
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)   //   ищем текст _dmarc в домене
	if err != nil {
		log.Printf("Error: %v\n", err)   //   выводим ошибку
	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){   //   если найден префикс DMARC1
			hasDMARC = true   //   переключаем
			dmarcRecord = record   //   получаем запись
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)   //   выводим всю охапку полученной инфы
}
