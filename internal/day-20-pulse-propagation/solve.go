package day20pulsepropagation

import (
	"fmt"
	"strings"

	"github.com/ewoutquax/advent-of-code-2023/pkg/register"
	"github.com/ewoutquax/advent-of-code-2023/pkg/utils"
)

const Day string = "20"

var counter int

func init() {
	register.Day(Day+"a", solvePart1)
	register.Day(Day+"b", solvePart2)
}

func solvePart1(inputFile string) {
	lines := utils.ReadFileAsLines(inputFile)
	u := ParseInput(lines)

	for idx := 0; idx < 1000; idx++ {
		u.PressButton()
	}

	fmt.Printf("Result of day-%s / part-1: %d\n", Day, u.NrPulses())
}

func solvePart2(inputFile string) {
	// lines := utils.ReadFileAsLines(inputFile)
	// u := ParseInput(lines)

	nrClicks := commonLcm([]int{3947, 4019, 3943, 4007})

	fmt.Printf("Result of day-%s / part-2: %d\n", Day, nrClicks)
}

type TypePulse uint

const (
	TypeLow TypePulse = iota + 1
	TypeHigh
)

type ModuleName string

type Module interface {
	getName() ModuleName
	AddDestination(Module)
	RegisterInput(Module)
	ReceivePulse(Module, TypePulse, *Universe)
}

type ModuleBase struct {
	Name         ModuleName
	Destinations []Module
}

type ModuleBroadcaster struct {
	ModuleBase
}

type ModuleOutput struct {
	ModuleBase
}

type ModuleFlipFlop struct {
	ModuleBase
	IsOn bool
}

type ModuleConjunction struct {
	ModuleBase
	Inputs map[ModuleName]TypePulse
}

func (mb *ModuleBroadcaster) getName() ModuleName { return "broadcaster" }
func (mo *ModuleOutput) getName() ModuleName      { return "output" }
func (mf *ModuleFlipFlop) getName() ModuleName    { return mf.Name }
func (mc *ModuleConjunction) getName() ModuleName { return mc.Name }

func (mb *ModuleBroadcaster) AddDestination(m Module) { mb.Destinations = append(mb.Destinations, m) }
func (mo *ModuleOutput) AddDestination(m Module)      { mo.Destinations = append(mo.Destinations, m) }
func (mf *ModuleFlipFlop) AddDestination(m Module)    { mf.Destinations = append(mf.Destinations, m) }
func (mc *ModuleConjunction) AddDestination(m Module) { mc.Destinations = append(mc.Destinations, m) }

func (mb *ModuleBroadcaster) RegisterInput(m Module) { return }
func (mo *ModuleOutput) RegisterInput(m Module)      { return }
func (mf *ModuleFlipFlop) RegisterInput(m Module)    { return }
func (mc *ModuleConjunction) RegisterInput(m Module) { mc.Inputs[m.getName()] = TypeLow }

func (mb *ModuleBroadcaster) ReceivePulse(sender Module, typePulse TypePulse, u *Universe) {
	for _, receiver := range mb.Destinations {
		u.AddToQueue(mb, receiver, typePulse)
	}
}

func (mf *ModuleFlipFlop) ReceivePulse(sender Module, typePulse TypePulse, u *Universe) {
	if typePulse == TypeLow {
		mf.IsOn = !mf.IsOn

		// if mf.IsOn {
		// 	fmt.Printf("FlipFlop %v is now ON\n", mf.Name)
		// } else {
		// 	fmt.Printf("FlipFlop %v is now OFF\n", mf.Name)
		// }

		var sendPulseType TypePulse
		if mf.IsOn {
			sendPulseType = TypeHigh
		} else {
			sendPulseType = TypeLow
		}
		for _, destinationModule := range mf.Destinations {
			u.AddToQueue(mf, destinationModule, sendPulseType)
		}
	}
}

func (mc *ModuleConjunction) ReceivePulse(sender Module, typePulse TypePulse, u *Universe) {
	mc.Inputs[sender.getName()] = typePulse

	// fmt.Printf("ModuleConjunction: ReceivePulse: '%s' remembered inputs mc.Inputs: %v\n", mc.Name, mc.Inputs)

	var sendTypePulse TypePulse = TypeLow
	for subname, rememberedPulse := range mc.Inputs {
		if rememberedPulse != TypeHigh {
			sendTypePulse = TypeHigh
		} else {
			if mc.getName() == "cl" {
				fmt.Printf("cl has high input for subname: %v\n", subname)
				switch subname {
				case "dt":
					u.highAfterNrPressesDt = append(u.highAfterNrPressesDt, counter)
				case "js":
					u.highAfterNrPressesJs = append(u.highAfterNrPressesJs, counter)
				case "qs":
					u.highAfterNrPressesQs = append(u.highAfterNrPressesQs, counter)
				case "ts":
					u.highAfterNrPressesTs = append(u.highAfterNrPressesTs, counter)
				default:
					panic("No valid case found")
				}
			}
		}
	}

	for _, destinationModule := range mc.Destinations {
		u.AddToQueue(mc, destinationModule, sendTypePulse)
	}
}

func (mo *ModuleOutput) ReceivePulse(_ Module, _ TypePulse, _ *Universe) {}

type PulseSenderReceiver struct {
	Sender   Module
	Receiver Module
	TypePulse
}

type Universe struct {
	Modules    map[ModuleName]Module
	PulseQueue []PulseSenderReceiver

	highAfterNrPressesDt []int
	highAfterNrPressesJs []int
	highAfterNrPressesQs []int
	highAfterNrPressesTs []int

	nrLowPulses  int
	nrHighPulses int
}

func (u *Universe) AddToQueue(sender Module, receiver Module, typePulse TypePulse) {
	u.PulseQueue = append(u.PulseQueue, PulseSenderReceiver{
		Sender:    sender,
		Receiver:  receiver,
		TypePulse: typePulse,
	})
}

func (u *Universe) PressButton() {
	// fmt.Println("")
	// fmt.Println("Button is pressed!")
	// fmt.Println("------------------")

	u.AddToQueue(&ModuleOutput{}, u.Modules["broadcaster"], TypeLow)
	// u.Modules["broadcaster"].ReceivePulse(&ModuleOutput{}, TypeLow, u)

	for len(u.PulseQueue) > 0 {
		pulse := u.PulseQueue[0]
		u.PulseQueue = u.PulseQueue[1:]

		// fmt.Printf("solve: %s -%s-> %s\n", pulse.Sender.getName(), convPulse(pulse.TypePulse), pulse.Receiver.getName())

		if pulse.TypePulse == TypeHigh {
			u.nrHighPulses++
		} else {
			u.nrLowPulses++
		}

		(pulse.Receiver).ReceivePulse(pulse.Sender, pulse.TypePulse, u)
	}
}

func (u *Universe) NrPulses() int {
	fmt.Printf("u.nrLowPulses: %v\n", u.nrLowPulses)
	fmt.Printf("u.nrHighPulses: %v\n", u.nrHighPulses)

	return u.nrLowPulses * u.nrHighPulses
}

func PressButtonUntilRxOn(u *Universe) int {
	var doContinue bool = true

	fmt.Printf("u.Modules[\"cl\"]: %v\n", u.Modules["cl"])

	for doContinue {
		fmt.Printf("counter: %v\n", counter)
		counter++

		u.PressButton()

		if len(u.highAfterNrPressesDt) > 0 &&
			len(u.highAfterNrPressesJs) > 0 &&
			len(u.highAfterNrPressesQs) > 0 &&
			len(u.highAfterNrPressesTs) > 0 {
			fmt.Printf("u.highAfterNrPressesDt: %v\n", u.highAfterNrPressesDt)
			fmt.Printf("u.highAfterNrPressesJs: %v\n", u.highAfterNrPressesJs)
			fmt.Printf("u.highAfterNrPressesQs: %v\n", u.highAfterNrPressesQs)
			fmt.Printf("u.highAfterNrPressesTs: %v\n", u.highAfterNrPressesTs)

			panic("Done")
		}
	}

	return counter
}

func ParseInput(lines []string) Universe {
	u := Universe{
		Modules: make(map[ModuleName]Module, len(lines)),
	}

	u.Modules["output"] = &ModuleOutput{
		ModuleBase: ModuleBase{
			Name:         "output",
			Destinations: make([]Module, 0),
		},
	}

	var module Module

	// Build each module
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		switch {
		case parts[0] == "broadcaster":
			module = buildBroadcaster(parseName(line))
		case string(parts[0][0]) == "%":
			module = buildFlipFlop(parseName(line))
		case string(parts[0][0]) == "&":
			module = buildConjunction(parseName(line))
		default:
			panic("No valid case found")
		}
		u.Modules[module.getName()] = module
	}

	// Register each destination module as both input and destinations
	for _, line := range lines {
		moduleName := parseName(line)
		module = u.Modules[moduleName]

		parts := strings.Split(line, " -> ")
		for _, subname := range strings.Split(parts[1], ", ") {
			subModule := u.Modules[ModuleName(subname)]
			if subModule == nil {
				subModule = &ModuleFlipFlop{
					ModuleBase: ModuleBase{
						Name:         ModuleName(subname),
						Destinations: make([]Module, 0),
					},
					IsOn: false,
				}
			}

			module.AddDestination(subModule)
			subModule.RegisterInput(module)

			u.Modules[subModule.getName()] = subModule
			u.Modules[module.getName()] = module

			// fmt.Println("Registered submodule")
			// fmt.Println("------------------")
			// fmt.Printf("current: %v\n", moduleName)
			// fmt.Printf("childname: %v\n", subname)
			// fmt.Printf("u.Modules: %v\n\n", u.Modules)
		}
	}

	return u
}

func buildBroadcaster(name ModuleName) *ModuleBroadcaster {
	return &ModuleBroadcaster{
		ModuleBase: ModuleBase{
			Name:         name,
			Destinations: make([]Module, 0),
		},
	}
}

func buildFlipFlop(name ModuleName) *ModuleFlipFlop {
	return &ModuleFlipFlop{
		ModuleBase: ModuleBase{
			Name:         name,
			Destinations: make([]Module, 0),
		},
		IsOn: false,
	}
}

func buildConjunction(name ModuleName) *ModuleConjunction {
	return &ModuleConjunction{
		ModuleBase: ModuleBase{
			Name:         name,
			Destinations: make([]Module, 0),
		},
		Inputs: make(map[ModuleName]TypePulse),
	}
}

func parseName(line string) (name ModuleName) {
	parts := strings.Split(line, " ")
	if parts[0] == "broadcaster" {
		name = "broadcaster"
	} else {
		name = ModuleName(parts[0][1:])
	}

	return
}

func convPulse(pulse TypePulse) string {
	if pulse == TypeLow {
		return "low"
	}

	return "high"
}

func commonLcm(nrs []int) int {
	var int1 int = 1
	for _, nr := range nrs {
		int1 = lcm(int1, nr)
	}

	return int1
}

func lcm(n, m int) int {
	return (n * m) / gcd(n, m)
}

func gcd(n, m int) int {
	if m == 0 {
		return n
	}
	return gcd(m, n%m)
}
