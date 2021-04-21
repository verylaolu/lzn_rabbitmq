package main

import (
	 "github.com/verylaolu/lzn_rabbitmq/Interface"
)
func main()  {
	//push()
	pull()
}
func push(){
	Push := Interface.Push_OBJ{}
	//Push.Dead()
	//Push.Normal()
	Push.Delay()
}
func pull(){
	Pull:=Interface.Pull_OBJ{}
	//Pull.Dead()
	//Pull.Normal()
	Pull.Delay()
}
