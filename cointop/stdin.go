package cointop

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ReadAPIKeyFromStdin reads the user inputed API from the stdin prompt
func (ct *Cointop) ReadAPIKeyFromStdin(name string) string {
	ct.debuglog("ReadAPIKeyFromStdin()")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s API Key: ", name)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(text)
}
