package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

var records = map[string]string{
	"soham.service.": "Cool",
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
			
		case dns.TypeA:
			 s:= strings.Split(q.Name, ".")
			records[s[0]] = s[1]
			log.Printf("Query: %s", q.Name)
			ip := records[q.Name]
			log.Printf("IP: %s\n", ip)

			if ip != "" {
				o, err := makeResp("soham.service. 0 IN TXT " + ip)
				if err == nil {
					m.Answer = append(m.Answer, o[0])
				}
			}
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}
	w.WriteMsg(m)
}

func main() {

	
	port := 8000
	server := &dns.Server{Addr: ":" + strconv.Itoa(port), Net: "udp"}

	log.Printf("Starred at %d\n", port)

	dns.HandleFunc(".", handleDnsRequest)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}

func makeResp(ans string) ([]dns.RR, error) {
	out := make([]dns.RR, 0, len(ans))
	r, err := dns.NewRR(ans)
	if err != nil {
		return nil, err
	}
	out = append(out, r)
	return out, nil
}
