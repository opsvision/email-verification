package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"strings"
)

var (
	sender string
	email  string
)

/**
 * Initialize the command line arguments
 */
func init() {
	flag.StringVar(&sender, "sender", "jdoe@acme.com", "the sender email address")
	flag.StringVar(&email, "email", "jdoe@acme.com", "the email address to verify")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

/**
 * Main entry point for our program
 */
func main() {
	// Make sure we have something to work with
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(-1)
	}

	// Extract the domain from the email addy
	domain := extractDomain(email)

	// Determine the MX server for the extracted domain
	records := getMXRecords(domain)
	if len(records) == 0 {
		log.Fatalf("No MX records found for %s\n", domain)
	}

	checkEmail(records[0], email)
}

/**
 *
 */
func extractDomain(email string) string {
	if index := strings.Index(email, "@"); index == 0 {
		log.Fatal("Malformed email address")
	}

	return email[strings.Index(email, "@")+1:]
}

/**
 *
 */
func getMXRecords(domain string) []string {
	var records []string

	mxs, _ := net.LookupMX(domain)
	for _, mx := range mxs {
		host := mx.Host

		// Remove trailing dot
		if last := len(host) - 1; last >= 0 && host[last] == '.' {
			host = host[:last]
		}

		records = append(records, host)
	}

	return records
}

/**
 *
 */
func checkEmail(server string, email string) {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "%s:%d", server, 25)

	// Connect
	//log.Printf("Connecting to %s\n", server)
	c, err := smtp.Dial(buffer.String())
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail(sender); err != nil {
		log.Fatal(err)
	}

	// Test email
	log.Printf("Checking email %s\n", email)
	status := "VALID"
	if err := c.Rcpt(email); err != nil {
		status = "INVALID"
	}
	fmt.Printf("%s|%s\n", email, status)

	// Quit gracefully
	//log.Printf("Closing gracefully\n")
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
