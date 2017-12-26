# Rasplight

## gRPC-Interface

## Setup
- Raspberry:
    - Red-LED: GPIO02
    - Green-LED: GPIO03
    - Blue-LED: GPIO04
- Telnet-Port: 6600

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