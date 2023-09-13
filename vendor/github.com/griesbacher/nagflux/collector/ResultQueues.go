package collector

import "github.com/mg-castel/nagflux/data"

type ResultQueues map[data.Target]chan Printable
