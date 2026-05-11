package counter 


type Counter struct{
	Value int 
}

// increment 

func(c *Counter) increment(){
	c.Value++

}

// value

func (c *Counter) value()int{
	return c.Value
}

