package main

import (
	"errors"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nxadm/tail"
)

// Parses nginx logs and streams out to a channel

type HTTP_METHOD int

type LogEntry struct {
	IP        net.IP
	time      time.Time
	method    string
	distro    string
	url       string
	version   string
	status    int
	bytesSent int
	agent     string
}

// It is critical that NGINX uses the following log format:
// "$remote_addr" "$time_local" "$request" "$status" "$body_bytes_sent" "$request_length" "$http_user_agent";

var reQuotes *regexp.Regexp

func ReadLogs(logFile string, channels []chan *LogEntry) (err error) {
	// Compile regular expressions
	reQuotes, err = regexp.Compile(`"(.*?)"`)
	if err != nil {
		return err
	}

	// Tail the log file `tail -F`
	tail, err := tail.TailFile(logFile, tail.Config{ReOpen: true})
	if err != nil {
		return err
	}

	for line := range tail.Lines {
		entry, err := ParseLine(line.Text)
		if err != nil {
			log.Printf("[WARN] failed to line %s %s", line.Text, err.Error())
		}

		// Send a pointer to the entry down each channel
		for _, ch := range channels {
			select {
			case ch <- entry:
			default:
				// TODO: Warn that a channel is starting to hang
			}
		}
	}

	return nil
}

func ParseLine(line string) (*LogEntry, error) {
	// "172.21.0.3" "03/Jan/2022:00:28:33 +0000" "GET /cosi_background.png HTTP/1.1" "200" "130107" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:95.0) Gecko/20100101 Firefox/95.0"
	quoteList := reQuotes.FindAllString(line, -1)

	if len(quoteList) != 6 {
		return nil, errors.New("invalid number of parameters in log")
	}

	var log *LogEntry

	// IPv4 or IPv6 address
	log.IP = net.ParseIP(quoteList[0])
	if log.IP == nil {
		return nil, errors.New("failed to parse ip")
	}

	// Time
	t := "02/Jan/2006:15:04:05 -0700"
	tm, err := time.Parse(t, quoteList[1])
	if err != nil {
		return nil, err
	}
	log.time = tm

	// Method url http version
	split := strings.Split(quoteList[2], " ")
	if len(split) != 3 {
		// this should never fail
		return nil, errors.New("invalid number of strings in request")
	}
	log.method = split[0]
	log.url = split[1]
	log.version = split[2]
	// TODO Extract the disto from the url

	// HTTP response status
	status, err := strconv.Atoi(quoteList[3])
	if err != nil {
		// this should never fail
		return nil, errors.New("could not parse http response status")
	}
	log.status = status

	// Bytes sent
	bytesSent, err := strconv.Atoi(quoteList[4])
	if err != nil {
		// this should never fail
		return nil, errors.New("could not parse bytes_sent")
	}
	log.bytesSent = bytesSent

	// User agent
	log.agent = quoteList[5]

	return log, nil
}
