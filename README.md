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
