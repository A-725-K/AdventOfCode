import std/cmdline
import std/macros
import std/os
import std/sequtils
import std/strformat
import std/sugar

from std/strutils import parseInt, splitLines, isEmptyOrWhitespace, endsWith
from std/syncio import readFile

import day_tasks


macro importAllDays(dir: static string) =
  result = newStmtList()
  for f in staticExec("ls -1 days").splitLines():
    if f.endsWith(".nim"):
      result.add parseStmt("import days/" & f.splitFile.name)


proc checkCliParams(argv: seq[string]) =
  if
    argv.len() == 1 and (
      argv[0] == "-h" or
      argv[0] == "--help" or
      argv[0] == "-?"
    ):
    echo "Expecting 2 or 3 arguments: <DAY: 1|..|25> <PART: 1|2> [<FILENAME>]"
    quit(QuitSuccess)

  if argv.len == 0 or argv.len > 3:
    echo "Usage: nim c -r aoc_main.nim <DAY: 1|..|25 <PART: 1|2> [<SIZE: small|big>]"
    quit(QuitFailure)


proc main() =
  let argv = commandLineParams()
  checkCliParams(argv)

  var
    day, part: int
    inputFilename, srcFilename, taskKey: string

  try:
    day = parseInt(argv[0])
    if day < 1 or day > 25:
      echo "Select the correct day between 1 and 25: ", day
      quit(QuitFailure)
  except:
    echo "Expecting a number, got: ", argv[0]
    quit(QuitFailure)

  try:
    part = parseInt(argv[1])
    if part < 1 or part > 2:
      echo "Decide if you want to run part 1 or 2: ", part
      quit(QuitFailure)
  except:
    echo "Expecting a number, got: ", argv[1]
    quit(QuitFailure)

  inputFilename =
    if argv.len == 3:
      case argv[2]:
      of "small": &"inputs/day{day}/mini_input"
      of "big": &"inputs/day{day}/input"
      else: argv[2]
    else:
      &"inputs/day{day}/mini_input"

  srcFilename = &"days/day{day}.nim"
  taskKey = &"day{day}_{part}"

  var lines = readFile(inputFilename)
    .splitLines(keepEol=false)
    .filter(s => not isEmptyOrWhitespace(s))

  if taskKey notin tasks:
    echo &"This tasks has not been registered yet: {taskKey}"
    quit(QuitFailure)

  # Execute the code for the daily task
  tasks[taskKey](lines, day)


importAllDays("days")
main()
