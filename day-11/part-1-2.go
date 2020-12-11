package main

import (
	"io/ioutil"
	"strings"
)

const (
	floor = iota
	empty
	occupied
)

type Matrix struct {
	arr  []int
	rows int
	cols int
}

func (m *Matrix) iter() []int {
	return m.arr
}

func (m *Matrix) copy(other *Matrix) {
	if m.arr == nil {
		m.arr = make([]int, len(other.arr))
	}
	copy(m.arr, other.arr)
	m.rows = other.rows
	m.cols = other.cols
}

func (m *Matrix) get(row int, col int) (int, bool) {
	if row >= m.rows || col >= m.cols || row < 0 || col < 0 {
		return -1, false
	}

	return m.arr[row*m.cols+col], true
}

func (m *Matrix) set(row int, col int, val int) {
	if row >= m.rows || col >= m.cols {
		panic("Index out of bounds")
	}

	m.arr[row*m.cols+col] = val
}

func newMatrix(lines []string) *Matrix {
	rows := len(lines)
	cols := len(lines[0])
	m := &Matrix{make([]int, rows*cols), rows, cols}
	for i, line := range lines {
		for j, ru := range line {
			var seat int
			if ru == '.' {
				seat = floor
			} else if ru == 'L' {
				seat = empty
			} else {
				seat = occupied
			}
			m.set(i, j, seat)
		}
	}
	return m
}

func occupiedAdjacent(m *Matrix, row int, col int) int {
	adjacentSum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				if seat, ok := m.get(row+i, col+j); ok && seat == occupied {
					adjacentSum++
				}
			}
		}
	}
	return adjacentSum
}

func occupiedInSight(m *Matrix, row int, col int) int {
	adjacentSum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				for distance := 1; distance < m.rows; distance++ {
					seat, ok := m.get(row+i*distance, col+j*distance)
					if ok && seat == floor {
						continue
					}
					if ok && seat == occupied {
						adjacentSum++
					}
					break
				}
			}
		}
	}
	return adjacentSum
}

func applySeatRules(curSeats *Matrix, nextSeats *Matrix, tolerance int, occupiedNearMethod func(*Matrix, int, int) int) bool {
	seatsStateChanged := false
	rows := curSeats.rows
	cols := curSeats.cols
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			seat, ok := curSeats.get(i, j)
			if ok && seat == empty && occupiedNearMethod(curSeats, i, j) == 0 {
				nextSeats.set(i, j, occupied)
				seatsStateChanged = true
			} else if ok && seat == occupied && occupiedNearMethod(curSeats, i, j) >= tolerance {
				nextSeats.set(i, j, empty)
				seatsStateChanged = true
			} else {
				nextSeats.set(i, j, seat)
			}
		}
	}
	return seatsStateChanged
}

func solve(lines []string, tolerance int, occupiedNearMethod func(*Matrix, int, int) int) int {
	seats := newMatrix(lines)

	// occupy all seats
	for i := 0; i < seats.rows; i++ {
		for j := 0; j < seats.cols; j++ {
			if seat, ok := seats.get(i, j); ok && seat == empty {
				seats.set(i, j, occupied)
			}
		}
	}

	seatsSwap := &Matrix{nil, 0, 0}
	seatsSwap.copy(seats)

	for applySeatRules(seats, seatsSwap, tolerance, occupiedNearMethod) {
		seats, seatsSwap = seatsSwap, seats
	}

	seatsOccupied := 0
	for _, seat := range seats.iter() {
		if seat == occupied {
			seatsOccupied++
		}
	}
	return seatsOccupied
}

func part1(lines []string) {
	println(solve(lines, 4, occupiedAdjacent))
}

func part2(lines []string) {
	println(solve(lines, 5, occupiedInSight))
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	lines = lines[:len(lines)-1]
	// 	test := strings.Split(`L.LL.LL.LL
	// LLLLLLL.LL
	// L.L.L..L..
	// LLLL.LL.LL
	// L.LL.LL.LL
	// L.LLLLL.LL
	// ..L.L.....
	// LLLLLLLLLL
	// L.LLLLLL.L
	// L.LLLLL.LL`, "\n")
	// fmt.Printf("%v\n", test)

	part1(lines)
	part2(lines)
}
