package redis

import (
	"testing"
	"time"
)

var p *Pool

func init() {
	host := "127.0.0.1:6379"
	password := ""
	p = &Pool{
		MaxIdle: 30,
		MaxActive: 30,
		IdleTimeout: time.Duration(time.Minute),
		Dial: func() (Conn, error) {
			c, err := Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},TestOnBorrow: func(c Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func TestRedis(t *testing.T) {
	testSet(t, p)
	testGet(t,p)
	if err := p.Close(); err != nil {
		t.Errorf("redis: close error(%v)", err)
	}
}

func testSet(t *testing.T, p *Pool) {
	key   := "test"
	value := "test"
	conn  := p.Get()
	defer conn.Close()
	if reply, err := conn.Do("set", key, value); err != nil {
		t.Errorf("redis: conn.Do(SET, %s, %s) error(%v)", key, value, err)
	} else {
		t.Logf("redis: set status: %s", reply)
	}
}

func testGet(t *testing.T,p *Pool) {
	key := "test"
	conn := p.Get()
	defer conn.Close()
	if reply,err := conn.Do("GET",key);err != nil {
		t.Errorf("redis: conn.Do(SET, %s) error(%v)", key, err)
	} else {
		t.Logf("redis: get value by key : %s ",reply)
	}
}

