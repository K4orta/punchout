import json, urllib2
import RPi.GPIO as GPIO
import sys, time, signal, threading

GPIO.setmode(GPIO.BCM)
led_pins = [4,17,22,23]
#lights
for i in led_pins:
    GPIO.setup(i, GPIO.OUT)

lights = {
    'clockIn': GPIO.PWM(4,50),
    'mealBreak': GPIO.PWM(17,50),
    'clockOut': GPIO.PWM(22,50),
    'error': GPIO.PWM(23,50)
}

# button 
GPIO.setup(25, GPIO.IN)

def start_lights():
    for light in lights:
        lights[light].start(0)

def stop_lights():
    for light in lights:
        lights[light].stop(0)

def punch_request():
    term = threading.Event()
    resp = urllib2.urlopen('http://0.0.0.0:8081/press')
    out = json.loads(resp.read())
    WaitingLight(lights[out['punch']], term).start()
    resp = urllib2.urlopen('http://0.0.0.0:8081/confirm')
    out = json.loads(resp.read())
    term.set()

def signal_handler(signal, frame):
    print 'You pressed Ctrl+C!'
    stop_lights()
    GPIO.cleanup()
    sys.exit(0)

class WaitingLight(threading.Thread):
    def __init__(self, light, term_event):
        threading.Thread.__init__(self)
        self.light = light
        self.term_event = term_event

    def run(self):
        while(not self.term_event.is_set()):
            self.light.ChangeDutyCycle(0)
            for i in range(20):
                self.light.ChangeDutyCycle(i)
                self.term_event.wait(0.02)
            for i in range(20):
                self.light.ChangeDutyCycle(20-i)
                self.term_event.wait(0.02)
        self.light.ChangeDutyCycle(100)
        for i in range(100):
            self.light.ChangeDutyCycle(100-i)
            time.sleep(0.01)

class index:
    def GET(self):
        return json.dumps({'done':True})

if __name__ == '__main__':
    signal.signal(signal.SIGINT, signal_handler)
    start_lights()
    
    while True: 
        if (GPIO.input(25) == False):
            punch_request()
            time.sleep(1)
