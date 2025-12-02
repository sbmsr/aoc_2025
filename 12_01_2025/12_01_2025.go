package main

import (
  "os" 
  "bufio"
  "strconv"
)

func part1() int {
  file, err := os.Open("input.txt")
  if (err != nil) { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  acc := 0
  current_position := 50

  for scanner.Scan() {
    line := scanner.Text()
    num, err := strconv.Atoi(line[1:])
    if (err != nil) { panic(err) }

    if (line[0] == 'L') {
      current_position -= num
    } else {
      current_position += num
    }

    current_position %= 100

    if (current_position < 0) {
      current_position += 100
    } 
    
    if (current_position == 0) { acc += 1 }
  }

  return acc
}

func part2() int {
  file, err := os.Open("input.txt")
  if (err != nil) { panic(err) }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  acc := 0
  current_position := 50

  for scanner.Scan() {
    line := scanner.Text()
    num, err := strconv.Atoi(line[1:])
    if (err != nil) { panic(err) }

    click := 1
    if (line[0] == 'L') {
      click = -1
    }

    for (num > 0) {
      current_position += click
      num -= 1
      current_position %= 100
      if (current_position == 0) { acc += 1 } 
    }
  }

  return acc
}

func main() {
  println("part1:", part1())
  println("part2:", part2())
}
