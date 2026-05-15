package parser


type Parser interface {
	LogAnalyzer(logChan <-chan string, alertChan chan<- string)
}
