package main
//
// import (
//   "os"
//   "fmt"
//   "bufio"
//   "strconv"
//   queue "AdventOfCode/ds/queue"
// )
//
// func getQueue(str string) *queue.Queue[string] {
//   q := queue.NewQueue[string]()
//
//   strLen := len(str)
//   var nStr string
//   for i := 0; i < strLen; i++ {
//     if str[i] == '[' {
//       q.Enqueue("[")
//     } else if str[i] == ']' {
//       if nStr != "" {
//         q.Enqueue(nStr)
//         nStr = ""
//       }
//       q.Enqueue("]")
//     } else if str[i] == ',' {
//       if nStr != "" {
//         q.Enqueue(nStr)
//         nStr = ""
//       }
//     } else {
//       nStr += string(str[i])
//     }
//   }
//
//   return q
// }
//
// func printQueues(left, right *queue.Queue[string]) {
//   fmt.Println("LEFT:")
//   for _, el := range left.GetQueue() {
//     fmt.Print(el, ", ")
//   }
//   fmt.Println()
//   fmt.Println("RIGHT:")
//   for _, el := range right.GetQueue() {
//     fmt.Print(el, ", ")
//   }
//   fmt.Println("\n")
// }
//
// func cmpInt(l, r string) int {
//   n1, err := strconv.Atoi(l)
//   if err != nil {
//     fmt.Println("1--------->", l)
//     panic("Cannot convert to number")
//   }
//   n2, err := strconv.Atoi(r)
//   if err != nil {
//     fmt.Println("2--------->", r)
//     panic("Cannot covert to number")
//   }
//
//   if n1 > n2 {
//     return -1
//   }
//   if n1 < n2 {
//     return 1
//   }
//   return 0
// }
//
// func comparePackets(l, r string, i int) bool {
//   left, right := getQueue(l), getQueue(r)
//   printQueues(left, right)
//
//   for !left.IsEmpty() && !right.IsEmpty() {
//     if i == 9 {
//       printQueues(left, right)
//     }
//     l, r := left.Dequeue(), right.Dequeue()
//
//     // both numbers, compare them
//     if l != "[" && l != "]" && r != "[" && r != "]" {
//       if cmpInt(l, r) == -1 {
//         return false
//       }
//       continue
//     }
//
//     // left number, right list
//     if l != "[" && l != "]" && r == "[" {
//       for right.Top() != "]" {
//         rr := right.Dequeue()
//         if rr == "[" {
//           return true
//         }
//         if cmpInt(l, rr) == -1 {
//           return false
//         }
//       }
//       right.Dequeue()
//       continue
//     }
//
//     // left list, right number
//     if l == "[" && r != "[" && r != "]" {
//       for left.Top() != "]" {
//         ll := left.Dequeue()
//         if ll == "[" {
//           return false
//         }
//         if cmpInt(ll, r) == -1 {
//           return false
//         }
//       }
//       left.Dequeue()
//     }
//   }
//
//   if !left.IsEmpty() && right.IsEmpty() {
//     return false
//   }
//
//   return true
// }
//
// func part1(scanner *bufio.Scanner) {
//   i := 0
//   index := 1
//   var idxs []int
//
//   var l, r string
//   for scanner.Scan() {
//     if i%3 == 0 {
//       l = scanner.Text()
//     } else if i%3 == 1 {
//       r = scanner.Text()
//     } else {
//       if comparePackets(l, r, index) {
//         idxs = append(idxs, index)
//       }
//       index++
//     }
//     i++
//   }
//
//   sum := 0
//   for _, idx := range idxs {
//     sum += idx
//   }
//
//   fmt.Println("The sum of the indexes of the in-order packets is:", sum)
// }
//
// func main() {
//   f, err := os.Open("input")
//   if err != nil {
//     panic(err)
//   }
//   defer f.Close()
//
//   scanner := bufio.NewScanner(f)
//   part1(scanner)
// }

import (
  "fmt"
  "os/exec"
)

func main() {
  cmd, err := exec.Command("/usr/bin/python3", "day13.py").Output()
  if err != nil {
    panic(err)
  }
  fmt.Println(string(cmd))
}

