package main
import (
    "bufio"
 	"fmt"
    "os"
   	"strconv"
    "strings"
)

var negativeSign = 0
var foundI = false
var foundJ = false
var foundK = false

func main() {
	miScanner := bufio.NewReader(os.Stdin)
	fmt.Print("Number of Cases: ")
	tests, _ := miScanner.ReadString('\n')
	tests = tests[:len(tests)-1]
	n := ToInt(tests)
	for i := 0; i<n; i++{
		test(i+1)
	}

	//test(1)
}

func test(caseNumber int){
	miScanner2 := bufio.NewReader(os.Stdin)
	fmt.Print("L X: ")
	lx, _ := miScanner2.ReadString('\n')
    arraylx := strings.Split(lx," ")
    var l = ToInt(arraylx[0])
    var x = ToInt(arraylx[1][:len(arraylx[1])-1])


    fmt.Print("string: ")
    s, _ := miScanner2.ReadString('\n')
    s = s[:len(s)-1]
    /*l := 3
    x := 1
    s := "ijk"*/
  	//fmt.Println(l,x,s)

    finalString := ""
    _ = l

    for i := 0; i<x; i++{
    	finalString = finalString +s
    }

    //fmt.Println(finalString)


    finalString = find(finalString,"i")
    //fmt.Println(finalString)
    finalString = find(finalString,"j")
    //fmt.Println(finalString)
    finalString = find(finalString,"k")
    //fmt.Println(finalString)

    if foundI&&foundJ&&foundK{
    	fmt.Println("Case #"+strconv.Itoa(caseNumber)+" YES")
    } else {
    	fmt.Println("Case #"+strconv.Itoa(caseNumber)+" NO")
    }

    foundI=false
    foundJ=false
    foundK=false
    fmt.Println("")
}

func find(word,toFind string) string{
	x := strings.Split(word,"")

	if len(x)==0{
		return ""
	}

	a := x[0]
	if(check(a,toFind)){
		return word[1:len(word)]
	}
	for i:=1; i<len(x); i++{
		a = multiply(a,x[i])

		if(check(a,toFind)){
			return word[i+1:len(word)]
		}
	}
	return ""
}

func ToInt(s string) int{
	x,_ := strconv.Atoi(s)
	//fmt.Println("-"+strconv.Itoa(x)+"-")
	return x
}

func check(a,toFind string) bool{
	var positive = (negativeSign%2)==0
	if a==toFind&&positive{
		if toFind=="i"{
			foundI=true
		}
		if toFind=="j"{
			foundJ=true
		}
		if toFind=="k"{
			foundK=true
		}

		negativeSign = 0

		return true
	}

	return false
}


func multiply(a,b string) string{

	//var array2D [2][2]string
	//array2D[0][0] = "tr"

	if a=="1"{
		return b
	} else if a=="i"{
		if b=="1"{
			return "i"
		}
		if b=="i"{
			negativeSign++
			return "1"
		}
		if b=="j"{
			return "k"
		}
		if b=="k"{
			negativeSign++
			return "j"
		}
	} else if a=="j"{
		if b=="1"{
			return "j"
		}
		if b=="i"{
			negativeSign++
			return "k"
		}
		if b=="j"{
			negativeSign++
			return "l"
		}
		if b=="k"{
			return "i"
		}
	} else if a=="k"{
		if b=="1"{
			return "k"
		}
		if b=="i"{
			return "j"
		}
		if b=="j"{
			negativeSign++
			return "i"
		}
		if b=="k"{
			negativeSign++
			return "1"
		}
	}
	return ""
}