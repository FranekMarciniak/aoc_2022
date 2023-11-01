package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	Id            int
	Meals         []int
	CaloriesCount int
}

func NewElf(id int) *Elf {
	return &Elf{
		Id:            id,
		Meals:         []int{},
		CaloriesCount: 0,
	}
}

func (e *Elf) AddMeal(meal int) {
	e.Meals = append(e.Meals, meal)
	e.CaloriesCount += meal
}

type Population struct {
	Elfs     []Elf
	LeadElfs []Elf
}

func NewPopulation() *Population {
	return &Population{}
}

func (p *Population) GetLatestElf() *Elf {
	if len(p.Elfs) == 0 {
		elf := NewElf(1)
		p.AddElf(*elf)
		p.LeadElfs = append(p.LeadElfs, *elf)
		return elf
	}
	return &p.Elfs[len(p.Elfs)-1]
}

func (p *Population) GetLeadersCalories() int {
	total := 0
	for _, e := range p.LeadElfs {
		total += e.CaloriesCount
	}
	return total
}

func (p *Population) AddElf(e Elf) {
	p.Elfs = append(p.Elfs, e)
}

func (p *Population) AddNextElf() {
	latest := p.GetLatestElf()
	newElf := NewElf(latest.Id + 1)
	p.AddElf(*newElf)
	p.ReplaceLeader(*latest)
}

func (p *Population) FindSmallestLeader() Elf {
	leader := p.LeadElfs[0]
	for _, e := range p.LeadElfs {
		if e.CaloriesCount < leader.CaloriesCount {
			leader = e
		}
	}
	return leader
}

func (p *Population) ReplaceLeader(newElf Elf) {
	if len(p.LeadElfs) < 3 {
		p.LeadElfs = append(p.LeadElfs, newElf)
		return
	}

	smallestLeader := p.FindSmallestLeader()
	if newElf.CaloriesCount <= smallestLeader.CaloriesCount {
		return
	}

	for i, elf := range p.LeadElfs {
		if elf.Id == smallestLeader.Id {
			p.LeadElfs[i] = newElf
			break
		}
	}
}

func main() {
	population := NewPopulation()
	fileName := "input.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file %s: %v", fileName, err)
	}
	defer closeFile(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLine(scanner.Text(), population)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}

	fmt.Println(population.GetLeadersCalories())
}

func processLine(line string, population *Population) {
	if line == "" {
		population.AddNextElf()
		return
	}

	meal, err := strconv.Atoi(line)
	if err != nil {
		log.Fatalf("Failed to convert line to integer: %v", err)
	}

	elf := population.GetLatestElf()
	elf.AddMeal(meal)
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatalf("Failed to close file: %v", err)
	}
}
