package main

import (
	"crypto/tls"
	"log"

	"gopkg.in/ldap.v2"
)

func main(){
	ldapURL := "localhost:10389" //앞에 프로토콜 붙이면 에러 뜬다

	l, err := ldap.DialTLS("tcp",ldapURL,&tls.Config{InsecureSkipVerify: true}) //연결 역할

	if err != nil {
			log.Fatal(err)
	}

	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}
}