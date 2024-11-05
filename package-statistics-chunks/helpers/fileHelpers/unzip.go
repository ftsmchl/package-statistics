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
	"sync"

	//"time"
)

type pkgStruct struct{
	pkgString string	
	count int
}

func UnzipAndCreateArrPackages(filePath, arch string, chunkSize int) ([]byte, error) {

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

	scanner := bufio.NewScanner(gzReader) // Create a bufio.Scanner to read the Gzip reader line by line

	var wg sync.WaitGroup // Wait group to wait for the goroutines to finish
	stringToCountMap := make(map[string]int) // Mapping of a string to times count

	pkgStrChan := make(chan []string) // Common shared channel between senders and receivers
	done := make(chan struct{}) // The channel to signal the end of the pkgStrChan

	var counter = 0

	
	receiverCounter := 0
	
	go func(){
		for {
			select{
			case stringsRecv := <- pkgStrChan:
				receiverCounter += 1
				for _,string := range stringsRecv {
					_, ok := stringToCountMap[string]
					if !ok {
						stringToCountMap[string] = 1
					}else{
						stringToCountMap[string] += 1
					}
				}
			case <-done:
				fmt.Println("I am done")
				//return
				break
			}
		}
	}()

	var linesArr []string
	// Iterate through the lines
	for scanner.Scan() {
		counter += 1
		line := scanner.Text() // Get the current line


		linesArr = append(linesArr, line)


		if len(linesArr) == chunkSize{

			wg.Add(1)
			go func(linesArr []string){
				if err := lines.ProcessLines(linesArr,  pkgStrChan);err != nil {
					log.Fatalf("Error occured in ProcessLine: %v counter = %v", err, counter)
				}
				defer wg.Done()
			}(linesArr)

			linesArr = nil
		}

	}

	// Iterate through remaining lines left
	if len(linesArr) > 0 {
		wg.Add(1)
		go func(linesArr []string){
			if err := lines.ProcessLines(linesArr,  pkgStrChan);err != nil {
				log.Fatalf("Error occured in ProcessLine: %v counter = %v", err, counter)
			}
			defer wg.Done()
		}(linesArr)

		linesArr = nil
	}

	wg.Wait()

	fmt.Println("counter = ", counter)
	done <- struct{}{} // Close the receiver channel

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading Gzip file: %v\n", err)
	}

	var arrayPackages []pkgStruct

	// Port pkgs from map to array in order to reach compatibility from the sort.Slice function
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







