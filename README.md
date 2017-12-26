# Rasplight

## gRPC-Interface

## Setup
- connect Leds to you PI
- copy the content from ``raspi-dir`` to your raspberry
- enable service on pi with:
    - ``systemctl enable rpilight.service``
- Configure your LED-GPIO pins in ``/etc/rpilight/config.toml``
   - Defaults are GPIO02, GPIO03, GPIO04

## Telnet commands (Linux Terminal)
- Turn light on:
    - ```echo -n "on " | nc [ip] 6600```
- Turn light off:
    - ```echo -n "off " | nc [ip] 6600```
- Set Colors:
    - ```echo -n "colors 255 0 64 " | nc [ip] 6600```
- Set Brightness
    - ```echo -n "bright 40 " | nc [ip] 6600```
    
## Known issues
- [x] two on or off in a row kills programm
- [ ] telnet commands won't work without space at the end
- [ ] Awfull hack to get IP: Loop till ip hasbeen received instead of wait for system is online

## Roadmap
- [x] Switch LEDs per Telnet on off
- [x] Set Color and Dimm lights per telnet
- [x] gRPC support with StateChange stream
- [ ] add bonjour to find device in network
- [ ] add an alarm function to turn on/off on given time

- [ ] bring to Cloud-Services:
    - [ ] Google Home
    - [ ] Amazone Echo
    - [ ] IFTTT
    - [ ] Home-Assistant