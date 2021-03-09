package main

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type UniqueRand struct {
	generated   map[int]bool    //keeps track of
	rng         *rand.Rand      //underlying random number generator
	scope       int             //scope of number to be generated
}

func NewUniqueRand(N int) *UniqueRand{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return &UniqueRand{
		generated: map[int]bool{},
		rng:        r1,
		scope:      N,
	}
}

func (u *UniqueRand) Int() int {
	if u.scope > 0 && len(u.generated) >= u.scope {
		return -1
	}
	for {
		var i int
		if u.scope > 0 {
			i = u.rng.Int() % u.scope
		}else{
			i = u.rng.Int()
		}
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func MakeRandomMap() map[string]interface{} {
	result := map[string]interface{}{}
	rng := NewUniqueRand(2*10)
	rng2 := NewUniqueRand(2*10000000000)
	for i:=0; i<10; i++{
		result[fmt.Sprintf("junkdata%v", rng.Int())] = rng2.Int() ^ i
	}
	return result
}
func jsonShuffleObject1(data []byte) string {
	str := string(data)
	if bytes.ContainsRune(data,'{') ||bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		remove1 := strings.Trim(str, "{}")
		splits := strings.Split(remove1, ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}
func jsonShuffleObject2(data []byte) string {
	str := string(data)
	if bytes.ContainsRune(data,'{') ||bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		remove1 := strings.Replace(str,"{","",0)
		remove2 := strings.Replace(str,"{","",len(remove1))
		splits := strings.Split(remove2, ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}
func jsonShuffleObject3(data []byte) string {
	str := string(data)
	if bytes.ContainsRune(data,'{') ||bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		remove1 := strings.ReplaceAll(str,"{","")
		remove2 := strings.ReplaceAll(remove1,"{","")
		splits := strings.Split(remove2, ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}
func jsonShuffleObject4(data []byte) string {
	str := string(data)
	if bytes.ContainsRune(data,'{') ||bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		re := regexp.MustCompile(`(?m)^{(.*?)}$`)
		remove1 := re.ReplaceAllString(str,"$1")
		splits := strings.Split(remove1, ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}
func jsonShuffleObject5(data []byte) string {
	if bytes.ContainsRune(data, '{') || bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		rune1 := bytes.Runes(data)
		rune1 = rune1[1:len(rune1)-1]
		splits := strings.Split(string(rune1), ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}
func jsonShuffleObject6(data []byte) string {
	if bytes.ContainsRune(data,'{') || bytes.ContainsRune(data,'}') {
		sb := strings.Builder{}
		sb.WriteString("{")
		rune1 := string(data)
		rune1 = rune1[1:len(rune1)-1]
		splits := strings.Split(string(rune1), ",")
		rand.Shuffle(len(splits), func(i, j int) {
			splits[i], splits[j] = splits[j], splits[i]
		})
		sb.WriteString(strings.Join(splits,","))
		sb.WriteString("}")
		return sb.String()
	} else {
		return ""
	}
}

func main() {
			junkdata := MakeRandomMap()
			junkdata["time"] = time.Now()
			junkdata["data"] = "testdata"
			jsondata,_ := json.Marshal(junkdata)
			d1 := jsonShuffleObject1(jsondata)
			d2 := jsonShuffleObject2(jsondata)
			d3 := jsonShuffleObject3(jsondata)
			d4 := jsonShuffleObject4(jsondata)
			d5 := jsonShuffleObject5(jsondata)
			d6 := jsonShuffleObject6(jsondata)
			log.Println(d1)
			log.Println(d2)
			log.Println(d3)
			log.Println(d4)
			log.Println(d5)
			log.Println(d6)

			loopcount := 100000
			t := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject1(jsondata)
			}
			log.Println("1번방식 ",time.Since(t))
			t1 := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject2(jsondata)
			}
			log.Println("2번방식 ",time.Since(t1))
			t3 := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject3(jsondata)
			}
			log.Println("3번방식 ",time.Since(t3))
			t4 := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject4(jsondata)
			}
			log.Println("4번방식 ",time.Since(t4))

			t5 := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject5(jsondata)
			}
			log.Println("5번방식 ",time.Since(t5))
			t6 := time.Now()
			for i:= 0; i< loopcount; i++ {
				_ = jsonShuffleObject6(jsondata)
	}
	log.Println("6번방식 ",time.Since(t6))
}
