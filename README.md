
1-Copy the Following Main Package in your  %GOPATH%src/github.com/user/firstprogram in firstprogram.go
2-use the command in command line// go get github.com/Arslan-MAS/quiz.git 
3-install using go install github.com/Arslan-MAS/firstprogram
4-go to %GOPATH%bin and type./firstprogram to run
"
package main

import (
"fmt"
"github.com/Arslan-MAS/quiz"
)

func main() {
    //reading an integer
    var option int
    fmt.Println("\n",1,"Print Covid cases in Pakistan\n",2,"Print Covid cases in SouthKorea\n",3,"Print Covid cases in France\n",4,"Print my personalized message to Coronavirus\n",0,"Exit\n")
    _, err:= fmt.Scan(&option) 
    if err == nil {
    fmt.Println("Input is ", option )
    fmt.Println(quiz.ReturnValue(option))
    }
}
"