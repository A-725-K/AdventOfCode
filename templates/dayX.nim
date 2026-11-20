import ../day_tasks
import std/strformat


proc part1*(lines: seq[string], day: int) =
  echo &"Day {day}, part 1"


proc part2*(lines: seq[string], day: int) =
  echo &"Day {day}, part 2"


tasks["day<DAY>_1"] = part1
tasks["day<DAY>_2"] = part2
