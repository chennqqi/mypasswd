package main

import (
	"fmt"
	"log"

	"github.com/chennqqi/mypasswd"
)

func init() {
	log.SetFlags(0)
}

func readStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", err
	}

	str := string(input)
	str = strings.TrimRight(str, "\n")

	return str, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	m := &mypasswd.Mypasswd{}

	var err error
	flag.StringVar(&m.Passwd, "p", "", "Password string")
	flag.BoolVar(&m.Old, "old", false, "Returns the value of the pre-4.1 implementation of PASSWORD()")
	flag.Parse()

	if m.Passwd == "" {
		m.Passwd, err = readStdin()
	}

	if err != nil {
		log.Fatal(err)
	}

	hash, err := m.Password()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}
