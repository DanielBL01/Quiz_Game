## Project Help Settings

```sh
Usage of quiz_game.exe:
  -Limit int
        The time limit for the quiz in seconds (default 30)
  -problems string
        Answer simple math questions from csv file (default "problems.csv")
```

## Playing the Game
What's great about Go is it's very easy to build and download binaries. Create your csv problem set and challenge yourself with a timer limit!

```sh
github.com\DanielBL01\Go_Projects\quiz_game>quiz_game.exe -problems=problems.csv -Limit=5
Problem #1: 5+5 = 10
Correct!
Problem #2: 1+1 = 2
Correct!
Problem #3: 8+3 = 11
Correct!
Problem #4: 1+2 = 3
Correct!
Problem #5: 8+6 = 14
Correct!
Problem #6: 3+1 =
You got 5 out of 8 questions correct!
```
