package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestParseLocal(t *testing.T) {
	fmt.Println(ParseLocal(time.Now().UTC()))
}

func TestParseLocalString(t *testing.T) {
	fmt.Println(ParseLocalString("2019-09-11 23:20:00.000+0000").UTC().Add(time.Second * 28800).Format("2006-01-02 15:04:05.000+0000"))
}

func TestToScyllaTimestamp(t *testing.T) {
	fmt.Println(ToScyllaTimestamp(time.Now().String()))
}
