package filehelpers

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sorting-example/helpers/lines"
	//"sync"
	//"time"
)

type pkgStruct struct{
	pkgString string	
	count int
}

func UnzipAndCreateArrPackages(filePath, arch string) ([]byte, error) {

	filePath = filepath.Join(filePath, "Contents-"+arch+".gz")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Wrong opening the file: %v\n", err)
	}
	defer file.Close()

	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("Error creating Gzip reader: %v\n", err)
	}
	defer gzReader.Close() // Ensure the Gzip reader is closed after we're done

	// Create a bufio.Scanner to read the Gzip reader line by line
	scanner := bufio.NewScanner(gzReader)

	stringToCountMap := make(map[string]int)

	var counter = 0

	//var linesArr []string
	// Iterate through the lines
	for scanner.Scan() {
		counter += 1
		line := scanner.Text() // Get the current line
		if err := lines.ProcessLine(line,  stringToCountMap);err != nil {
			log.Fatalf("Error occured in ProcessLine: %v counter = %v", err, counter)
		}


	}

	/*
	if len(linesArr) > 0 {
		wg.Add(1)
		go func(linesArr []string){
			for _, lineProcess := range linesArr {
				if err := lines.ProcessLine(lineProcess,  pkgStrChan);err != nil {
					log.Fatalf("Error occured in ProcessLine: %v counter = %v", err, counter)
				}
			}
			defer wg.Done()
		}(linesArr)

		linesArr = nil
	}
	*/




	fmt.Print("I am closing done")
	fmt.Print("counter = ", counter)
//	done <- struct{}{}


	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading Gzip file: %v\n", err)
	}

	//time.Sleep(5 * time.Second)
	fmt.Println("edw mono ena")

	// the array of packages
	var arrayPackages []pkgStruct

	// port pkgs from map to array in order to use the sort.Slice function
	for pkg, counts := range stringToCountMap {
		arrayPackages = append(arrayPackages, pkgStruct{
			pkgString: pkg,
			count: counts,
		})
	} 

	sort.Slice(arrayPackages, func(i,j int) bool{
		return arrayPackages[i].count > arrayPackages[j].count
	})


	if len(arrayPackages) >= 10 {
		fmt.Println("statistics of Packages in descending order ", arrayPackages[:10])
	}else{
		fmt.Println("statistics of Packages in descending order ", arrayPackages)
	}

	return nil, nil
}







