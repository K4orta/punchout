import web, json
import RPi.GPIO as GPIO
import sys, signal

GPIO.setmode(GPIO.BCM)
led_pins = []
lights = {}
# button 
GPIO.setup(25, GPIO.IN)

while True: 
    if (GPIO.input(25) == False):
      print("Button Pressed")


