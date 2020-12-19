package vw

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"gitlab.com/opennota/morph"
)

var ports = []string{"58000"}

type VwDaemon struct {
	port  string
	i     uint64
	conns []net.Conn
	m     sync.Mutex
}

type VwStorage struct {
	demons []*VwDaemon
	i      uint64
	m      sync.Mutex
}

func NewVwDemon(port string) (*VwDaemon, error) {
	c := make([]net.Conn, 0, len(ports))
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	for i := 0; i < len(ports); i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			return nil, err
		}
		c = append(c, conn)
	}

	return &VwDaemon{
		port:  port,
		i:     0,
		conns: c,
		m:     sync.Mutex{},
	}, nil
}

func NewVwStorage() (*VwStorage, error) {
	d := make([]*VwDaemon, 0, len(ports))
	for _, port := range ports {
		vad, err := NewVwDemon(port)
		if err != nil {
			return nil, err
		}
		d = append(d, vad)
	}

	return &VwStorage{
		demons: d,
		i:      0,
		m:      sync.Mutex{},
	}, nil
}

func (s *VwStorage) getConn() net.Conn {
	s.m.Lock()
	defer s.m.Unlock()
	s.i++
	if s.i == uint64(len(ports)) {
		s.i = 0
	}
	return s.demons[s.i].getConn()
}

func (d *VwDaemon) getConn() net.Conn {
	d.m.Lock()
	defer d.m.Unlock()
	d.i++
	if d.i == uint64(len(ports)) {
		d.i = 0
	}

	return d.conns[d.i]
}

func FormatData(data string) (formatData []string, err error) {
	reg, err := regexp.Compile("[^a-zA-Zа-яА-Я0-9]+")
	if err != nil {
		return formatData, err
	}

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		line += string(time.Now().Unix())
		processedString := reg.ReplaceAllString(line, " ")
		words := strings.Split(processedString, " ")
		normWords := make([]string, 0, len(words))
		for _, word := range words {
			_, norms, tags := morph.XParse(word)
			if len(norms) > 0 && tags[0] != "PREP" {
				normWords = append(normWords, norms[0])
			}
		}
		formatData = append(formatData, fmt.Sprintf("| %s\n", strings.Join(normWords, " ")))
	}
	return formatData, err
}

func (s *VwStorage) Predict(formatData []string) (predict []int, err error) {
	conn := s.getConn()
	r := bufio.NewReader(conn)

	for _, line := range formatData {
		result, err := sendToPredict(conn, r, line)
		if err != nil {
			return predict, err
		}
		intResult, err := strconv.Atoi(result)
		if err != nil {
			return predict, err
		}
		predict = append(predict, intResult)
	}

	return predict, err
}

func sendToPredict(conn net.Conn, r *bufio.Reader, line string) (string, error) {
	_, err := conn.Write([]byte(line))
	if err != nil {
		return "", err
	}

	bin, err := r.ReadBytes(byte(10))
	if err != nil {
		return "", err
	}
	if len(bin) > 0 {
		bin = bin[:len(bin)-1]
	}
	return string(bin), nil
}
