package client

const (
	whoInfo = `
---------------------------------------------------------
                     Welcome!
This is laboratory work 2 for Computer Networks subject.
                    Variant 15.
         Designed and developed by Nikita Masych.
           Distributed under MIT License 2022.
---------------------------------------------------------
`
	helpInfo = `
---------------------------------------------------------
Commands:
-help: 
	See documentation.
-who:
	See credentials.
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
		2022/10/26 02:18:48 Server:
		>> sort -t 0 "'Say the line senior dev' - 'it depends'"
		2022/10/26 02:19:03 Server:  
---------------------------------------------------------
`
	unknownInfo           = "Undefined command"
	noSortMeasurementInfo = "No sort has been conducted"
)
