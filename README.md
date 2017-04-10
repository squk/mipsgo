# mipsgo
MIPS Simulator written in Go with a web interface

### Custom Pseudo-Instructions
```assembly
pbin $rs  # print R[$rs] as binary
phex $rs  # print R[$rs] as hexadecimal
pdec $rs  # print R[$rs] as decimal
break     # insert breakpoint
```

By default, the server directory will serve at localhost:8080
View it live at http://mips.nieves.io/
