package beanstalkd

import (
	"fmt"
	"github.com/kr/beanstalk"
	"time"
)

type BeanstalkdTube struct {
	conn    *beanstalk.Conn
	host    string
	port    string
	tube    beanstalk.Tube
	tubeset beanstalk.TubeSet
}

func (bt *BeanstalkdTube) Put(body []byte, pri uint32, delay, ttr time.Duration) (id uint64, err error) {
	id, err = bt.tube.Put(body, pri, delay, ttr)
	return
}

func (bt *BeanstalkdTube) Reserve(timeout time.Duration) (id uint64, body []byte, err error) {
	id, body, err = bt.tubeset.Reserve(timeout)
	return
}

func beanstalkdDial(host, port string) (*beanstalk.Conn, error) {
	uri := fmt.Sprintf("%s:%s", host, port)
	c, err := beanstalk.Dial("tcp", uri)

	if err != err {
		return nil, err
	}

	return c, nil
}

func NewBeansTalkdTube(host, port, tube string) (*BeanstalkdTube, error) {
	conn, err := beanstalkdDial(host, port)

	if err != nil {
		return nil, err
	}

	conn.Tube = beanstalk.Tube{conn, tube}
	conn.TubeSet = *beanstalk.NewTubeSet(conn, tube)

	return &BeanstalkdTube{
		conn:    conn,
		host:    host,
		port:    port,
		tube:    conn.Tube,
		tubeset: conn.TubeSet,
	}, nil
}
