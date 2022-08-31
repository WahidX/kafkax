package xkafka

import "fmt"

// Need to add validations
type ConsumerOptions struct {
	Topic     string
	GroupID   string
	Partition int
	Offset    int64
	IsJSON    bool
}

func (c *ConsumerOptions) Validate() bool {
	if c.Partition != 0 && len(c.GroupID) != 0 {
		fmt.Println("Error: Either Partition or GroupID may be specified at once")
		return false
	}

	if c.Partition <= 1 {
		c.Partition = 0
	}

	if c.Offset <= 1 {
		c.Offset = 0
	}

	return true
}
