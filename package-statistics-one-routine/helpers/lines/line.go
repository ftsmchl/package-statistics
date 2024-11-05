package lines

import (
	//"fmt"
	//"fmt"
	//"fmt"
	"strings"
	//"sync"
	//"regexp"
)

//func ProcessLine(line string, strToCount map[string]int, mu *sync.Mutex) error{
//func ProcessLine(line string, strToCount map[string]int, chanStr chan string) error{
func ProcessLine(line string, strToCount map[string]int) error{
//func ProcessLine(line string,  chanStr chan string) error{

	// remove all the whitespaces between the path of the packages and the actual packages
	theStringSplit := strings.Fields(line)
	//if len(theStringSplit) != 2 {
		//fmt.Println(theStringSplit)
		//return fmt.Errorf("the string is not as expected length = %v", len(theStringSplit))
	//}

	// split the packages based on the "," 
	//packages := strings.Split(theStringSplit[1], ",")
	packages := strings.Split(theStringSplit[len(theStringSplit) - 1], ",")

	//packagesLength := len(packages)
	//packagesLastRune := len(packagesLength)

	lastString := packages[len(packages) - 1]

	if lastString[len(lastString) - 1] == ']' {
		//lastString[lastString[:len(lastString)-1]]
		packages[len(packages) - 1] = packages[len(packages)-1][:len(lastString)-1]
	}



	// strToCount critical section lock is acquired 
	//mu.Lock()
	for  _, aPackage := range packages {
		//fmt.Print("Iam sending")

		//chanStr <- aPackage	
		//fmt.Println("I sent")



	
		
		_, ok := strToCount[aPackage]
		if !ok {
			strToCount[aPackage] = 1
		}else{
			strToCount[aPackage] += 1
		}
		
		
	} 
	//mu.Unlock()

	//fmt.Println("Inside the function ==> ", strToCount)


	return nil

} 
