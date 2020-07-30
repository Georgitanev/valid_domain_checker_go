package main
// Open domain names from "short.txt" and check them if they are valid or not
// print only valid domain names
// save valid domains in "save_doc.txt"
// Enjoy :) 


import (
     "fmt"
     "os"
     "log"
	 "bufio"
	 "strings"
	 "regexp"
)

//list_tlds_1dotgr = ('.gr')
//list_tlds_2dotsgr = ('.com.gr','.edu.gr','.gov.gr','.net.gr','.org.gr', '.mil.gr', '.mod.gr', '.sch.gr')


func main() {
    file, err := os.Open("short.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var re = regexp.MustCompile(".bg|.gr|.es")
    tlds1dot := ".a.bg|.b.bg|.c.bg|.d.bg|.e.bg|.f.bg|.g.bg|.h.bg|.i.bg|.j.bg|.k.bg|.l.bg|.m.bg|.n.bg|" +
            ".o.bg|.p.bg|.q.bg|.v.bg|.r.bg|.s.bg|.t.bg|.u.bg|.x.bg|.y.bg|.z.bg|.1.bg|.2.bg|.3.bg|" +
            ".4.bg|.5.bg|.1.bg|.6.bg|.7.bg|.8.bg|.9.bg|.0.bg" +
            ".com.gr|.edu.gr|.gov.gr|.net.gr|.org.gr|.mil.gr|.mod.gr|.sch.gr"
    var re2 = regexp.MustCompile(tlds1dot)
    file1, err := os.OpenFile("save_doc.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatalf("failed creating file: %s", err)
	 }
 		
    for scanner.Scan() {
        if (strings.Count(scanner.Text(), ".") == 1 && re.MatchString(scanner.Text()) == true ) || (strings.Count(scanner.Text(), ".") == 2 && re2.MatchString(scanner.Text()) == true ) {
			var word1 string
			word1 = scanner.Text()
			fmt.Println(word1)
        	datawriter := bufio.NewWriter(file1)
         datawriter.WriteString(word1 + "\n")
    	   datawriter.Flush()
		}
    }
file1.Close()}

