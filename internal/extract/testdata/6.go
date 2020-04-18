package testdata

import (
	"net"
	"time"
)

// 6 Checking several outgoing parameters.
// входные параметры: 0
// выходные параметры: режим template.QueryRow: 5 параметров, нативные, составные, массив, ошибка

// GoDao: generate
type GoDao6 struct {
	// language=PostgreSQL
	SelectBots func() (ip net.IPNet, connectTime []time.Time, isBot bool) `
        with "tmp"("ip", "connect_time", "is_bot") as (values (
            '2a02:6b8:b081:502::1:a'::inet,
            '{"2020-04-13T15:12:15+03:00","2020-04-13T14:12:15+03:00","2020-04-13T13:12:15+03:00"}'::timestamp with time zone[],
            true
        )) select "ip", "connect_time", "is_bot"
        from "tmp"
		where "is_bot";`
}
