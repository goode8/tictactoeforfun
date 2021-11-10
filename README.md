# tictactoeforfun
this was in a go class it and was super fun

Looking now, there is a lot of repeated code. Given more time, I would fix that. And make sure the score is working. And add a bit more pizzazz. 

Tic-Tac-Toe Elements:
Board
Current Game Play
Current Game display (output) at each round of play
Current Game Play (input with some error detection (5 tries))
Name, Level, Choice of square for X, play again.
There are four moves on each side and perhaps a fifth move by the human.
Find feature - looks for valid choice of square, is it number or X, O.
Update Board replaces square number with X or O
Choose Level
..omitted for time: Auto, machine to machine. (I tested a few strategies and just let them tie.)
Random (uses random number generator seeded by time)
Easy (only defense-looking for two Xs)
Medium (only offense-looking for two Ohs)
Difficult (offense and defense), but you can beat it.
Error Handle: if no level chosen, defaults to Difficult

End of Game checks for three in a row and a winner after each Update Board.
If all squares are filled without a winner – Boards displays Cat’s Game.
Play Again – y gets you in, anything else, ends game 

and goes to scoring...does it go to scoring? 
Scoring rolled over with Play Again
Read/Write High Score from/to file
Output High Score current game and historical games included a prompt of your current game
score or message if you didn’t score.
