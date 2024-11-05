package lines

import (
	"strings"
)

func ProcessLines(lines []string,  chanStr chan <- []string) error{

	var packagesToSend []string

	// remove all the whitespaces between the path of the packages and the actual packages
	for _, line := range lines {
		theStringSplit := strings.Fields(line)

		// split the packages based on the "," 
		packages := strings.Split(theStringSplit[len(theStringSplit) - 1], ",")
		lastString := packages[len(packages) - 1]

		if lastString[len(lastString) - 1] == ']' {
			packages[len(packages) - 1] = packages[len(packages)-1][:len(lastString)-1]
		}

		for _, aPackage := range packages {
			packagesToSend = append(packagesToSend, aPackage)
		}

	}


	chanStr <- packagesToSend 


	return nil

} 
