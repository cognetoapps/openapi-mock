package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type delayHandler struct {
	nextHandler     http.Handler
	ExpRate         float64
	MinDelaySeconds float64
	minDelay        time.Duration
	dist            distuv.Exponential
	logger          logrus.FieldLogger
}

func (h delayHandler) randDelay() time.Duration {
	return secondsToDuration(h.dist.Rand()) + h.minDelay
}

// DelayHandler returns delayHandler address
// It also trigger initialization of delayHandler so it can precompute and
// initialize all necessary info to speed-up the work later
func DelayHandler(logger logrus.FieldLogger, nextHandler http.Handler, min, max float64) http.Handler {
	h := newDelayHandler(logger, nextHandler, min, max)
	return &h
}

// ServeHTTP generates random delay and sleeps for the delay amount
// It also logs the delay amount
func (h *delayHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	delay := h.randDelay()
	h.logger.Info(fmt.Sprintf("Injecting %dms delay", int(delay/time.Millisecond)))
	time.Sleep(delay)
	h.nextHandler.ServeHTTP(writer, request)
}

func newDelayHandler(logger logrus.FieldLogger, next http.Handler, expRate, min float64) delayHandler {
	h := delayHandler{
		logger:          logger,
		nextHandler:     next,
		ExpRate:         min,
		MinDelaySeconds: min,
	}
	h.dist = distuv.Exponential{
		Rate: expRate,
		Src:  rand.NewSource(uint64(time.Now().UnixNano())),
	}
	h.minDelay = secondsToDuration(min)
	return h
}

func secondsToDuration(seconds float64) time.Duration {
	return time.Duration(int(seconds*1000.0)) * time.Millisecond
}
