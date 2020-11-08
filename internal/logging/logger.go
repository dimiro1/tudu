package logging

import "log"

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

// PanicCouldNotInstantiateHandler log the error and call panic
func (l *Logger) PanicCouldNotInstantiateHandler(err error, handler string) {
	log.Panicf("error %s not able to instante %s", err, handler)
}

func (l *Logger) ErrorRendering(err error, handler string) {
	log.Printf("error %s rendering %s", err, handler)
}

func (l *Logger) ErrorInvalidRenderer(err error, handler string) {
	log.Printf("error %s invalid renderer %s", err, handler)
}

func (l *Logger) StartingApplication() {
	log.Printf("starting application")
}

func (l *Logger) RegisteringRoutes() {
	log.Printf("registering routes")
}

func (l *Logger) FinishedRegisteringRoutes() {
	log.Printf("finished Registering routes")
}

func (l *Logger) ListeningHTTP(address string) {
	log.Printf("listening %s", address)
}

func (l *Logger) FatalListeningHTTP(err error, address string) {
	log.Fatalf("error %s listening HTTP on address %s", err, address)
}
