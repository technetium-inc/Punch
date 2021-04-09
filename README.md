# Punch
 An esolang.
 
## Credits
The syntax was heavily inspired by [This language right here](https://github.com/FreakC-Foundation/FreakC) by [nguyenphuminh](https://github.com/nguyenphuminh). Check it out!!

## Syntax
 Like [Titanium](https://github.com/JavaCode7/Titanium), Punch runs on commands. What makes it different is the lack of flow control in the latter and the lack of variable definition in the former. The 7 commands that Punch has are detailed here:

 - `TrashTalk[str] <string>;` equivalent to `fmt.Print(<string>)` in Go.
 - `TrashTalkLn[str] <string>;` equivalent to `fmt.Println(<string>)` in Go.
 - `TrashTalk[int] <anything>;` you can use this command to print anything (no newline) but you should really use it to print just integers.
 - `TrashTalkLn[int] <anything>;` same as above but with `\n`.
 - `FaceSlap[var,int] <A-Y> <int>;` takes in any letter from A to Y and assigns the second parameter (which needs to be an int)
 - `TrashTalk[var] <A-Y>;` prints out the value of the variable. (no `\n`)
 - `TrashTalkLn[var] <A-Y>;` same as above but with `\n`.
 
 All statements must end with a semicolon and there ***must*** be a newline at the end of you file.

 In addition to the basic syntax, you also need to know about the lives system. Every program in Punch starts out with 100 lives and every statement that you type takes away from that number. anything to do with stdout (TrashTalk & TrashTalkLn) is worth 5 and variable assignments (FaceSlap) are worth 4.

## Downloads
### [1.0.0](https://github.com/technetium-inc/Punch/releases/tag/v1.0.0)
[Windows](https://github.com/technetium-inc/Punch/releases/download/v1.0.0/punch.exe)
