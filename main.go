package main

import (
	"github.com/Batcherss/dns-leaker/logger"
	"github.com/Batcherss/dns-leaker/report"
	"github.com/Batcherss/dns-leaker/resolver"
	"github.com/Batcherss/dns-leaker/utils"
	"flag"
	"fmt"
	"sync"
)
// you can change our dns services.
var dnsServers = []string{
    "1.1.1.1:53", "1.0.0.1:53",  
    "8.8.8.8:53", "8.8.4.4:53", 
    "9.9.9.9:53", "149.112.112.112:53",  
    "208.67.222.222:53", "208.67.220.220:53",  
    "64.6.64.6:53", "64.6.65.6:53",  
    "185.228.168.168:53", "185.228.169.168:53", 
    "94.140.14.14:53", "94.140.15.15:53",  
    "76.76.19.19:53", "76.223.122.150:53", 
    "76.76.2.0:53", "76.76.10.0:53",  
    "2.56.220.2:53", "95.85.95.85:53",  
    "23.94.60.240:53", "23.94.60.241:53",  
    "77.88.8.8:53", "77.88.8.1:53",  
    "208.67.222.222:53", "185.51.200.2:53",
    "185.51.200.3:53", "198.101.242.72:53", 
    "9.9.9.10:53", "8.26.56.26:53",  
    "8.20.247.20:53", "185.228.168.9:53",  
    "185.228.169.9:53", "208.67.220.222:53",  
    "64.6.64.6:53", "64.6.65.6:53",  
    "156.154.70.1:53", "156.154.71.1:53",  
    "209.244.0.3:53", "209.244.0.4:53",  
    "194.226.34.132:53", "194.226.34.133:53",
    "77.88.8.2:53", "185.228.168.168:53",
    "185.228.168.10:53", "185.228.169.10:53",  
    "8.8.8.8:53", "8.8.4.4:53",  
    "9.9.9.11:53", "9.9.9.12:53", 
    "149.112.112.113:53", "208.67.220.223:53",  
    "8.20.247.21:53", "185.51.200.4:53", 
    "185.228.168.11:53", "209.244.0.5:53",  
    "209.244.0.6:53", "77.88.8.3:53",
}

func main() {
	var domain string
	var debug bool
	var logFile string

	flag.StringVar(&domain, "s", "", "Domain to check (required)")
	flag.BoolVar(&debug, "d", false, "Enable debug mode")
	flag.StringVar(&logFile, "f", "", "Log file (optional)")
	flag.Parse()

	if domain == "" {
		fmt.Println("Error: specify domain with -s")
		flag.Usage()
		return
	}

	logger.InitLogger(logFile)
	defer logger.CloseLogger()

	logger.Info("Starting DNS leak check for domain: " + domain)

	results := make(map[string][]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, server := range dnsServers {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()
			if debug {
				logger.Debug("Querying DNS server: " + server)
			}
			ips, err := resolver.Resolve(domain, server)
			mu.Lock()
			defer mu.Unlock()
			if err != nil {
				logger.Error(fmt.Sprintf("[Error] Server %s: %v", server, err))
			} else {
				logger.Success(fmt.Sprintf("Response from %s: %v", server, ips))
				results[server] = ips
			}
			logger.Info(fmt.Sprintf("Finished checking server %s, moving to next...", server))
		}(server)
	}
	wg.Wait()

	var reference []string
	for _, ips := range results {
		if len(ips) > 0 {
			reference = ips
			break
		}
	}

	if len(reference) == 0 {
		logger.Error("Failed to get a reference IP address. Check if the domain is reachable.")
		return
	}

	leakFound := false
	finalResults := []report.DNSResult{}

	for server, ips := range results {
		isSame, missing, extra := utils.Compare(reference, ips)
		entry := report.DNSResult{
			Server: server,
			IPs:    ips,
			Status: "OK",
		}

		if !isSame {
			leakFound = true
			entry.Status = "LEAK DETECTED"
			logger.Error(fmt.Sprintf("[LEAK DETECTED] Server: %s", server))
			if len(missing) > 0 {
				logger.Error(fmt.Sprintf("Expected but missing IPs: %s", missing))
			}
			if len(extra) > 0 {
				logger.Error(fmt.Sprintf("Unexpected IPs found: %s", extra))
			}
		}
		finalResults = append(finalResults, entry)
	}

	report.GenerateReport(finalResults)

	if leakFound {
		logger.Error("Result: DNS leaks detected! Check the logs and report.json")
	} else {
		logger.Success("Result: No DNS leaks detected. All servers responded consistently.")
	}
}
