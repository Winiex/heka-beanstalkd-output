package beanstalkd

import (
	"errors"
	p "github.com/mozilla-services/heka/pipeline"
)

type BeanstalkdOutputConfig struct {
	Host string `toml:"beanstalkd_host"`
	Port string `toml:"beanstalkd_port"`
	Tube string `toml:"beanstalkd_tube"`
}

type BeanstalkdOutput struct {
	config    *BeanstalkdOutputConfig
	beansTube *BeanstalkdTube
}

func (bo *BeanstalkdOutput) ConfigStruct() interface{} {
	return &BeanstalkdOutputConfig{
		Host: "127.0.0.1",
		Port: "11000",
		Tube: "default",
	}
}

func (bo *BeanstalkdOutput) Init(config interface{}) (err error) {
	conf := config.(*BeanstalkdOutputConfig)
	bo.config = conf

	return
}

func (bo *BeanstalkdOutput) Run(or p.OutputRunner, h p.PluginHelper) (err error) {
	if or.Encoder() == nil {
		return errors.New("Encoder must be specified.")
	}

	var (
		e        error
		outBytes []byte
	)

	inChan := or.InChan()

	for pack := range inChan {
		outBytes, e = or.Encode(pack)
		if e != nil {
			or.LogError(e)
			continue
		}

		if outBytes == nil {
			continue
		}

		newTube, err := NewBeansTalkdTube(conf.Host, conf.Port, conf.Tube)

		if err != nil {
			continue
		}

		bo.beansTube.Put(outBytes, 0, 0, 3)

		newTube.Close()

		pack.Recycle()
	}

	return
}

func init() {
	p.RegisterPlugin("BeanstalkdOutput", func() interface{} {
		return new(BeanstalkdOutput)
	})
}
