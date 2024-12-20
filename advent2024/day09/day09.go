package advent202409

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/owendavies93/advent-go/util"
)

type Day struct {
	disk  map[int]int
	files map[int]File
	end   int
	next  int
}

type File struct {
	loc int
	len int
}

func (d *Day) Part1() any {
	loc := 0
	for loc < d.end {
		id, ok := d.disk[d.end]
		if ok {
			delete(d.disk, d.end)

			for {
				if _, ok := d.disk[loc]; !ok {
					break
				}
				loc++
			}
			d.disk[loc] = id
		}
		d.end--
	}

	total := 0
	for k, val := range d.disk {
		total += val * k
	}
	return total
}

func (d *Day) Part2() any {
	id := d.next - 1

	for id >= 0 {
		i := 0
		for i < d.files[id].loc {
			exists := false
			for j := 0; j < d.files[id].len; j++ {
				if _, ok := d.disk[i+j]; ok {
					exists = true
					break
				}
			}
			if !exists {
				for j := 0; j < d.files[id].len; j++ {
					delete(d.disk, d.files[id].loc+j)
					d.disk[i+j] = id
				}
				break
			}
			i++
		}
		id--
	}

	total := 0
	for k, val := range d.disk {
		total += val * k
	}
	return total
}

func (d *Day) ParseInput() {
	lines, err := util.ReadStrings("inputs/2024/09")
	if err != nil {
		fmt.Println("Error parsing input:", err)
	}

	parts := strings.Split(lines[0], "")

	d.disk = make(map[int]int)
	d.files = make(map[int]File)
	fill := true
	d.next = 0
	curr := 0

	for _, part := range parts {
		len, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error parsing input:", err)
		}

		if fill {
			d.files[d.next] = File{loc: curr, len: len}
			for j := 0; j < len; j++ {
				d.disk[curr+j] = d.next
			}
			d.next++
			fill = false
		} else {
			fill = true
		}
		curr += len
	}

	for k := range d.disk {
		if k > d.end {
			d.end = k
		}
	}
}
