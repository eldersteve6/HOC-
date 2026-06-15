Before students build the ASCII Art project, they need to be comfortable with:
Reading data from a file
Using maps
Error handling
String manipulation
Loops and slices
GET Ready 


QUEST 1
Quest 1: Character Counter
NO AI chatbot should be used- Just Team work with table members
Problem
Create a Go program that:
Accepts a word as a command-line argument.
Counts how many times each character appears.
Displays the result using a map.
Handles errors properly.
Example
go run . hello
Output:
h -> 1
e -> 1
l -> 2
o -> 1
Example 2
go run . banana
Output:
b -> 1
a -> 3
n -> 2
Requirements
Use a map.
Use error handling.
If no argument is provided:
go run .
Output:
Error: please provide a word

Starter Structure
package main

import (
   "fmt"
   "os"
)

func main() {
   // Check argument count

   // Get word

   // Create map

   // Count letters

   // Print result
}

If you want to do it like professional modularize the quest

Please Relate your understanding to your ASCII-ART Project while we wait for QUEST 2 individual understanding test.
Note: Please ensure you can do the above quest yourself without any help not cramming but proper understanding

QUEST 2
Extension 1 (10 minutes)
Ignore uppercase and lowercase differences.
Input:
go run . Hello
Output:
h -> 1
e -> 1
l -> 2
o -> 1
Hint:
strings.ToLower()

Extension 2 (15 minutes)
Read the word from a file.
File:
banana
Run:
go run . data.txt
Output:
b -> 1
a -> 3
n -> 2
Hint:
os.ReadFile()
This introduces the file handling they'll need for the banner files in the ASCII Art project.

TEST CASE
Your program must be able to handle:
go run . HELLOhello

OUTPUT:
h > 2
e > 2
l > 4
o > 2


Also at the same handle:

go run . data.txt 

For example if the data.txt file contains:
hiHI 

OUTPUT:
h > 2
I > 2 

If you run go run . byte.txt 

b > 1
y > 1
t > 2
e > 1
. > 1
x > 1


If you are up to it make the data.txt to return an error that says “An Empty Data File” when it is empty instead of outputting nothing in the terminal

QUEST 3
Follow-Up Exercise (Closer to ASCII Art)
Character Dictionary
Create a map where each letter has a value.
letters := map[string]string{
   "A": "Apple",
   "B": "Ball",
   "C": "Cat",
}
Input:
go run . ABC
Output:
Apple
Ball
Cat
If a character doesn't exist:
go run . AD
Output:
Apple
Error: D not found
This exercise is almost the same mental model as the ASCII Art project:
asciiMap["A"] = []string{
   "  A  ",
   " A A ",
   "AAAAA",
}
Later, instead of:
A -> Apple
they'll do:
A -> ASCII representation
So the progression becomes:
Maps → Error Handling → File Reading → Character Lookup → ASCII Art, which mirrors exactly how the actual project should be built.


SITE : https://docs.google.com/document/d/1rdhLastCjZMhsmioHjCrT8eMBxKaN21BGa_oCwaPgng/edit?tab=t.cavww11r9q1x
