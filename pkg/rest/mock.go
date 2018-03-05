package rest

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	_ REST = &Mock{}
)

type Mock struct {
	log  *logrus.Logger
	rand *rand.Rand
}

func NewMock() *Mock {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}
	log.SetLevel(logrus.DebugLevel)
	log.Infoln("using mock API")
	randSrc := rand.NewSource(time.Now().UnixNano())
	return &Mock{
		log:  log,
		rand: rand.New(randSrc),
	}
}
func (mock *Mock) Get(req Rq) error {
	mock.log.Infof("GET %q", req.Path.Build())
	return nil
}

func (mock *Mock) Put(req Rq) error {
	mock.log.Infof("PUT %q", req.Path.Build())
	return nil
}

func (mock *Mock) Post(req Rq) error {
	mock.log.Infof("POST %q", req.Path.Build())
	return nil
}

func (mock *Mock) Delete(req Rq) error {
	mock.log.Infof("DELETE %q", req.Path.Build())
	return nil
}

func (mock *Mock) SetToken(token string) {

}
