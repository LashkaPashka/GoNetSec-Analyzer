package rules

import (
	"fmt"
	"strings"

	"github.com/lashkapashka/GoNetSec_Analyzer/internal/parser"
)

var _ parser.Parser = (*Analyzer)(nil)

const (
	portSecurityPsecureViolation = "PORT_SECURITY-2-PSECURE_VIOLATION"
	ospfAdjchg                   = "OSPF-5-ADJCHG"
	secLoginLoginFailed          = "SEC_LOGIN-4-LOGIN_FAILED"
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) LogAnalyzer(logChan <-chan string, alertChan chan<- string) {
	for logMsg := range logChan {
		logMsg = strings.TrimSpace(logMsg)

		if strings.Contains(logMsg, portSecurityPsecureViolation) {
			alertChan <- fmt.Sprintf("УГРОЗА L2 (MAC Spoofing):\n%s", logMsg)
		} else if strings.Contains(logMsg, ospfAdjchg) && strings.Contains(logMsg, "DOWN") {
			alertChan <- fmt.Sprintf("ПРОБЛЕМА СЕТИ (OSPF Down):\n%s", logMsg)
		} else if strings.Contains(logMsg, secLoginLoginFailed) {
			alertChan <- fmt.Sprintf("ПОПЫТКА ВЗЛОМА (Brute-force):\n%s", logMsg)
		} else {
			fmt.Printf("Инфо: %s\n", logMsg)
		}
	}
}
