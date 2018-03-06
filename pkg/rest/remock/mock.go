package remock

import (
	"math/rand"
	"net/url"
	"time"

	"git.containerum.net/ch/kube-client/pkg/rest"
	"github.com/sirupsen/logrus"
)

var (
	_ rest.REST = &Mock{}
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
func (mock *Mock) Get(req rest.Rq) error {
	mock.log.Infof("GET %q", req.URL.Build())
	validator := RqValidator{req}
	if err := validator.ValidateURL(); err != nil {
		return err
	}
	return nil
}

func (mock *Mock) Put(req rest.Rq) error {
	mock.log.Infof("PUT %q", req.URL.Build())
	validator := RqValidator{req}
	if err := validator.ValidateURL(); err != nil {
		return err
	}
	return nil
}

func (mock *Mock) Post(req rest.Rq) error {
	mock.log.Infof("POST %q", req.URL.Build())
	validator := RqValidator{req}
	if err := validator.ValidateURL(); err != nil {
		return err
	}
	return nil
}

func (mock *Mock) Delete(req rest.Rq) error {
	mock.log.Infof("DELETE %q", req.URL.Build())
	validator := RqValidator{req}
	if err := validator.ValidateURL(); err != nil {
		return err
	}
	return nil
}

func (mock *Mock) SetToken(token string) {

}

type RqValidator struct {
	rest.Rq
}

func (rqv *RqValidator) ValidateURL() error {
	_, err := url.ParseRequestURI(rqv.URL.Build())
	return err
}
