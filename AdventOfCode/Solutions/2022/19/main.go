package main

import (
	"fmt"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"

	"adventOfCode/helpers"
)

type Blueprint struct {
	Id                     int
	OreRobotCost           int
	ClayRobotCost          int
	ObsidianRobotOreCost   int
	ObsidianRobotClayCost  int
	GeodeRobotOreCost      int
	GeodeRobotObsidianCost int

	MaxOreCost      int
	MaxClayCost     int
	MaxObsidianCost int
}

// --- Day 18: Boiling Boulders ---
func run(input string) (interface{}, interface{}) {
	var blueprints []Blueprint

	for _, blueprintText := range strings.Split(input, "\n") {
		blueprint := Blueprint{}

		_, err := fmt.Sscanf(blueprintText, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian",
			&blueprint.Id, &blueprint.OreRobotCost, &blueprint.ClayRobotCost, &blueprint.ObsidianRobotOreCost, &blueprint.ObsidianRobotClayCost, &blueprint.GeodeRobotOreCost, &blueprint.GeodeRobotObsidianCost)
		if err != nil {
			panic("")
		}
		blueprint.MaxOreCost = helpers.Max(helpers.Max(helpers.Max(blueprint.OreRobotCost, blueprint.ClayRobotCost), blueprint.ObsidianRobotOreCost), blueprint.GeodeRobotOreCost)
		blueprint.MaxClayCost = blueprint.ObsidianRobotClayCost
		blueprint.MaxObsidianCost = blueprint.GeodeRobotObsidianCost
		blueprints = append(blueprints, blueprint)
	}

	maxGeodes := helpers.Map(blueprints, func(blueprint Blueprint) int {
		cache := make(map[Factory]int)
		return RunSim(Factory{
			Robots:    Robots{OreRobot: 1},
			Resources: Resources{},
			TimeLeft:  24,
		}, blueprint, cache) * blueprint.Id
	})

	if len(blueprints) == 2 {
		maxGeodesB := helpers.Map(blueprints[:2], func(blueprint Blueprint) int {
			cache := make(map[Factory]int)
			return RunSim(Factory{
				Robots:    Robots{OreRobot: 1},
				Resources: Resources{},
				TimeLeft:  32,
			}, blueprint, cache) * blueprint.Id
		})

		return helpers.Sum(maxGeodes), maxGeodesB[0] * maxGeodesB[1]
	}

	maxGeodesB := helpers.Map(blueprints[:3], func(blueprint Blueprint) int {
		cache := make(map[Factory]int)
		return RunSim(Factory{
			Robots:    Robots{OreRobot: 1},
			Resources: Resources{},
			TimeLeft:  32,
		}, blueprint, cache) * blueprint.Id
	})

	return helpers.Sum(maxGeodes), maxGeodesB[0] * maxGeodesB[1] * maxGeodesB[2]
}

type Factory struct {
	Robots    Robots
	Resources Resources
	TimeLeft  int
}

type Robots struct {
	OreRobot      int
	ClayRobot     int
	ObsidianRobot int
	GeodeRobot    int
}

type Resources struct {
	Ore      int
	Clay     int
	Obsidian int
	Geodes   int
}

var BestEver = 0

func rangeSum(first, last int) int {
	return last*(last+1)/2 - ((first - 1) * first / 2)
}

func RunSim(factory Factory, bluePrint Blueprint, cache map[Factory]int) int {
	if factory.TimeLeft == 1 {
		AddResources(&factory.Resources, factory.Robots)
		return factory.Resources.Geodes
	}
	// If we can create a Geode robot each minute we should
	if factory.Robots.OreRobot >= bluePrint.GeodeRobotOreCost && factory.Robots.ObsidianRobot >= bluePrint.GeodeRobotObsidianCost {
		return factory.Resources.Geodes + rangeSum(factory.Robots.GeodeRobot, factory.Robots.GeodeRobot+factory.TimeLeft-1)
	}
	factory.TimeLeft--

	bestFactoryScore := 0

	if factory.Robots.OreRobot < bluePrint.MaxOreCost && factory.Resources.Ore >= bluePrint.OreRobotCost {
		newFactory := factory
		newFactory.Resources.Ore -= bluePrint.OreRobotCost
		newFactory.Robots.OreRobot++

		AddResources(&newFactory.Resources, factory.Robots)
		bestFactoryScore = helpers.Max(bestFactoryScore, TestNewFactory(newFactory, bluePrint, cache))
	}
	if factory.Robots.ClayRobot < bluePrint.MaxClayCost && factory.Resources.Ore >= bluePrint.ClayRobotCost {
		newFactory := factory
		newFactory.Resources.Ore -= bluePrint.ClayRobotCost
		newFactory.Robots.ClayRobot++

		AddResources(&newFactory.Resources, factory.Robots)
		bestFactoryScore = helpers.Max(bestFactoryScore, TestNewFactory(newFactory, bluePrint, cache))
	}
	if factory.Robots.ObsidianRobot < bluePrint.MaxObsidianCost && factory.Resources.Ore >= bluePrint.ObsidianRobotOreCost && factory.Resources.Clay >= bluePrint.ObsidianRobotClayCost {
		newFactory := factory
		newFactory.Resources.Ore -= bluePrint.ObsidianRobotOreCost
		newFactory.Resources.Clay -= bluePrint.ObsidianRobotClayCost
		newFactory.Robots.ObsidianRobot++

		AddResources(&newFactory.Resources, factory.Robots)
		bestFactoryScore = helpers.Max(bestFactoryScore, TestNewFactory(newFactory, bluePrint, cache))
	}
	if factory.Resources.Ore >= bluePrint.GeodeRobotOreCost && factory.Resources.Obsidian >= bluePrint.GeodeRobotObsidianCost {
		newFactory := factory
		newFactory.Resources.Ore -= bluePrint.GeodeRobotOreCost
		newFactory.Resources.Obsidian -= bluePrint.GeodeRobotObsidianCost
		newFactory.Robots.GeodeRobot++

		AddResources(&newFactory.Resources, factory.Robots)
		bestFactoryScore = helpers.Max(bestFactoryScore, TestNewFactory(newFactory, bluePrint, cache))
	}

	newFactory := factory
	AddResources(&newFactory.Resources, factory.Robots)
	bestFactoryScore = helpers.Max(bestFactoryScore, TestNewFactory(newFactory, bluePrint, cache))

	if factory.TimeLeft == 23 {
		println(fmt.Sprintf("%+v", factory))
	}
	return bestFactoryScore
}

func TestNewFactory(factory Factory, bluePrint Blueprint, cache map[Factory]int) int {
	factoryScore := 0
	if val, ok := cache[factory]; ok {
		factoryScore = val
	} else {
		factoryScore = RunSim(factory, bluePrint, cache)
		cache[factory] = factoryScore
	}
	return factoryScore
}

func AddResources(resources *Resources, robots Robots) {
	resources.Ore += robots.OreRobot
	resources.Clay += robots.ClayRobot
	resources.Obsidian += robots.ObsidianRobot
	resources.Geodes += robots.GeodeRobot
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
