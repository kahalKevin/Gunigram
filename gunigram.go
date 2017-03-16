package gunigram

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "strconv"
    "os"
    "strings"
    "sync"
)

var wg sync.WaitGroup

type Template map[string]int

type Templates []Template

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func isExist(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func copyMap(originalMap Template) Template{
	var newMap = make(Template)
	for k,v := range originalMap {
  		newMap[k] = v
	}
	return newMap
}

func Unigram(file string){	
	var header []string
	var template = make(Template)//make a map
	var allTemplate Templates
	finalNgram := ""

	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
	    line := scanner.Text()
	    words := strings.Fields(line)
	    for _,word := range words{
	    	template[word]=0
	    	if(!isExist(word,header)){
	    		header = append(header,word)
	    	}	    	
	    }
	}

	f, err = os.Open(file)
	check(err)
	defer f.Close()

	wg.Add(2)

	go func(){
		scanner = bufio.NewScanner(f)
		for scanner.Scan() {
		    line := scanner.Text()
		    words := strings.Fields(line)
		    var temp = copyMap(template)
		    for _,word := range words{
		    	temp[word]++
		    }
		    allTemplate = append(allTemplate,temp)
		}
		wg.Done()
	}()

	go func(){
		for i,feature := range header{
			if(i == len(header)-1){
				finalNgram += feature + "\n"
			}else{
				finalNgram += feature + ","
			}
		}
		wg.Done()
	}()

	wg.Wait()

	for _,gram := range allTemplate{
		for i,feature := range header{
			if(i == len(header)-1){				
				finalNgram += strconv.Itoa(gram[feature]) + "\n"
			}else{				
				finalNgram += strconv.Itoa(gram[feature]) + ","
			}
		}
	}

	fmt.Println(finalNgram)

	toFile := []byte(finalNgram)
	err = ioutil.WriteFile("ngram.csv", toFile, 0644)
	check(err)

}