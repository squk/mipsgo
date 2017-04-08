# mipsgo
MIPS Simulator written in Go with a web interface

### Custom Pseudo-Instructions
```assembly
pbin $rs  # print R[$rs] as binary
phex $rs  # print R[$rs] as hexadecimal
pdec $rs  # print R[$rs] as decimal
break     # insert breakpoint
```

### Building
Built with Go 1.8
* Download https://ace.c9.io/. Copy the "ace" directory into "server/public/js".  Ace serves as the in-browser editor for the MIPS assembly
* Download http://getbootstrap.com/. Copy the following files into their respective directory

| From        | To           |
| ------------- |:-------------:|
| bootstrap-3.3.7/dist/css/boostrap.min.css      | server/public/css/boostrap.min.css |
| bootstrap-3.3.7/dist/js/bootstrap.min.js      | server/public/js/boostrap.min.js      |
| bootstrap-3.3.7/dist/fonts/ | server/public/fonts/      | |
###### or use the bootstrap theme "Solar" from https://bootswatch.com/solar

By default, the server directory will serve at localhost:8080
