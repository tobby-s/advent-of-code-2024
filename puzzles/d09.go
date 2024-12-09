package puzzles

import (
	"strconv"

	"github.com/tobby-s/advent-of-code-2024/utils"
)

type diskblock struct {
	id     int
	length int
}

func D9P1() (result int) {
	data := utils.LoadData(`https://adventofcode.com/2024/day/9/input`)
	filesystem := data[0]
	disk := make([]int, 0)
	fileid := 0
	for i, b := range filesystem {
		l, _ := strconv.Atoi(string(b))
		if i%2 == 0 {
			for j := 0; j < l; j++ {
				disk = append(disk, fileid)
			}
			fileid += 1
		} else {
			for j := 0; j < l; j++ {
				disk = append(disk, -1)
			}
		}
	}
	lastFileIndex := 0
	for i := len(disk) - 1; i > 0; i-- {
		if disk[i] != -1 {
			lastFileIndex = i
			break
		}
	}
	firstSpaceIndex := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			firstSpaceIndex = i
			break
		}
	}
	for lastFileIndex > firstSpaceIndex {
		disk[lastFileIndex], disk[firstSpaceIndex] = disk[firstSpaceIndex], disk[lastFileIndex]
		for disk[lastFileIndex] == -1 {
			lastFileIndex -= 1
		}
		for disk[firstSpaceIndex] != -1 {
			firstSpaceIndex += 1
		}
	}
	for i, id := range disk {
		if id > -1 {
			result += i * id
		}
	}
	return
}

type (
	file struct {
		start, end, id int
	}
)

func D9P2() (result int) {
	data := utils.LoadData(`https://adventofcode.com/2024/day/9/input`)
	filesystem := data[0]
	disk := make([]file, 0)
	gaps := make([]file, 0)
	fileid := 0
	currentIndex := 0
	for i, b := range filesystem {
		l, _ := strconv.Atoi(string(b))
		if i%2 == 0 {
			disk = append(disk, file{
				start: currentIndex,
				end:   l + currentIndex,
				id:    fileid,
			})
			fileid += 1
		} else {
			gaps = append(gaps, file{
				start: currentIndex,
				end:   l + currentIndex,
			})
		}
		currentIndex = currentIndex + l
	}
	for i := len(disk) - 1; i >= 0; i-- {
		width := disk[i].end - disk[i].start
		for j := 0; gaps[j].start < disk[i].start; j++ {
			if gaps[j].end-gaps[j].start >= width {
				disk[i].start, disk[i].end, gaps[j].start = gaps[j].start, gaps[j].start+width, gaps[j].start+width
				break
			}
		}
	}
	for i := 0; i < len(disk); i++ {
		result += (disk[i].start + disk[i].end - 1) * (disk[i].end - disk[i].start) / 2 * disk[i].id
	}
	return
}
