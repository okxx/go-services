package redis

import "time"

var Server *Pool

type redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}
var Driver = &redis{}

func Setup() {
	Server = &Pool{
		MaxIdle: 30,
		MaxActive: 30,
		IdleTimeout: time.Duration(time.Minute),
		Dial: func() (Conn, error) {
			c, err := Dial("tcp", Driver.Host)
			if err != nil {
				return nil, err
			}
			if Driver.Password != "" {
				if _, err := c.Do("AUTH", Driver.Password); err != nil {
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