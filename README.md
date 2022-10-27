# Welcome!

This is example of client-server communication via sockets.

# Abstract:

Client starts new channel on port :1040, host: localhost.
It the main cycle it constantly reads user input commands and sends them to the server for processing.
Server accordingly constantly reads requests from client and returns result in respect to the input.

# State diagram:

![http://url/to/img.png](https://github.com/NikitaMasych/ACS-Lab2/blob/main/docs/StateDiagramLab2.png)

# Commands:

```
Commands:

-help: 
	See documentation.
-who:
	See credentials.
-time:
	See duration of the last sort.
-sort:
	Abstract:
		Sorts specified in quotes text.
	Arguments:
		-s (start position index, counting from 0)
		-l (length from start position; by default, string is sorted to the end)
		-t (type of sort: ascending or descending (1 and -1 accordingly); by default ascending)
	Examples:
		>> sort -s 2 -l 4 -t 1 "Obey SOLID!"
		2022/10/26 02:18:17 Server: "Ob SeyOLID!"
		>> sort -s 13 "Pov: you are javascript: "3" + 1 == "31", "3" - 1 == 2"
		2022/10/26 02:18:36 Server: "Pov: you are           """"""+,-1112333:====aacijprstv"
		>> sort "Read Clean Code by Robert Martin!"
		2022/10/26 02:18:48 Server: "     !CCMRRaaabbddeeeeilnnoorrtty"
		>> sort -t -1 "'Say the line senior dev' - 'it depends'"
		2022/10/26 02:19:03 Server:  "yvttssrponnnliiiheeeeeedddaS-''''       "
```

