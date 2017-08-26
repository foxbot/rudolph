package generator

import (
	"time"
)

// 0000000000000000000000000000000000000000000000000000000000000000
// [                    A                   ][   B  ][   C        ]
// A: <0-42>  Timestamp (millis since epoch)
// B: <43-50> Seed
// C: <51-64> Worker ID

// Generates snowflakes with the given information.
type SnowflakeGenerator struct {
	epoch time.Time
	seed  uint64
	wid   uint64
}

// Create a new SnowflakeGenerator with the given Epoch and Worker ID
// Worker ID should fit within an uint14
func NewSnowflakeGenerator(epoch time.Time, workerId uint16) SnowflakeGenerator {
	return SnowflakeGenerator{
		epoch: epoch,
		seed:  0,
		wid:   uint64(workerId),
	}
}

const millisInNanos = 1*10 ^ 6

func (self *SnowflakeGenerator) Generate() uint64 {
	return self.GenerateAt(time.Now())
}

func (self *SnowflakeGenerator) GenerateAt(at time.Time) uint64 {
	millis := uint64(at.Sub(self.epoch).Nanoseconds() / millisInNanos)
	id := (millis << 22) | (self.seed << 14) | self.wid
	if self.seed > 254 {
		self.seed = 0
	} else {
		self.seed++
	}
	return id
}
