# sux
The terminal single plexer




usage:
 sux "invoke mongo" "ls -a" "xset dpms force off"


todo:
 1) parse input commands
 2) create three screens
 	size of screen is just whatever entire screen is supposed to be.
 	color of screen is whatever it already is
 	selectively delete screens. Just delete current screen.
 3) do each input command in each screen
 	can have a global list of command names. They will also represent the view name.
 	sux "ls -a" "ls -R" yields ["ls -a", "ls -R"]
 4) show the screen that is the output for the first screen
 5) allow for switching between screens
 	we can use ctr-up, ctr-down to do this. 
 	we can show current screen number to side of screen or somewhere...