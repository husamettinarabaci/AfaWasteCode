package main

import (
	"os"
	"strconv"
	"time"
)

var currentUser string
var turnLeftEnb string = os.Getenv("MOTOR_L_ENB")
var turnLeftPwm string = os.Getenv("MOTOR_L_PWM")
var turnRigthEnb string = os.Getenv("MOTOR_R_ENB")
var turnRigthPwm string = os.Getenv("MOTOR_R_PWM")

func main() {
	motorProc()
}

func motorProc() {

	rpio.Open()
	leftEnb, _ := strconv.Atoi(turnLeftEnb)
	rightEnb, _ := strconv.Atoi(turnRigthEnb)
	leftPwm, _ := strconv.Atoi(turnLeftPwm)
	rightPwm, _ := strconv.Atoi(turnRigthPwm)
	pinLeftEnb := rpio.Pin(leftEnb)
	pinRightEnb := rpio.Pin(rightEnb)
	pinLeftPwm := rpio.Pin(leftPwm)
	pinRightPwm := rpio.Pin(rightPwm)
	pinLeftEnb.Output()
	pinRightEnb.Output()
	pinLeftPwm.Output()
	pinRightPwm.Output()

	pinLeftEnb.High()
	pinRightEnb.High()
	pinLeftPwm.Low()
	pinRightPwm.Low()
	time.Sleep(1 * time.Second)
	pinLeftPwm.High()
	pinRightPwm.Low()
	time.Sleep(3 * time.Second)
	pinLeftPwm.Low()
	pinRightPwm.High()
	time.Sleep(3 * time.Second)
	pinLeftEnb.Low()
	pinRightEnb.Low()
	pinLeftPwm.Low()
	pinRightPwm.Low()
	rpio.Close()

}
