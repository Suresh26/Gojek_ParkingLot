package slot

import (
	"fmt"
	"os"
	car "parking_lot/functional_spec/car"
	"sort"
	"strconv"
	"text/tabwriter"
)

type Slot struct {
	SlotNumber  int
	Car         car.Car
	IsAvailable bool
}

type Slots struct {
	TotalSlot         int
	AvailableCapacity int
	SlotInfo          map[int]Slot
}

type SlotCommand interface {
	CreateSlots(int) string
	ParkCarinNearestAvailSlot(car.Car) string
	LeaveFromParking(int) string
	SlotStatus()
	GetCarDetails(string, string) string
}

//CreateSlots create slots based on the slot capacity
func (slots *Slots) CreateSlots(cp int) (message string) {
	if cp <= 0 {
		message = "Enter Valid Parking Capacity Count"
		return
	}
	initialCap := slots.TotalSlot
	slots.TotalSlot += cp
	slots.AvailableCapacity += cp
	slots.createEmptySlot(initialCap)
	message = fmt.Sprintf("%s %d %s", "Created a parking lot with", initialCap+cp, "slots")
	return message
}

//Create Empty slots based on the slot capacity
func (slots *Slots) createEmptySlot(initialCap int) {
	if slots.SlotInfo == nil {
		slots.SlotInfo = make(map[int]Slot)
	}
	slotCapacity := &slots.TotalSlot
	for initialCap < *slotCapacity {
		slot := Slot{}
		slotNo := initialCap + 1
		slot.SlotNumber = slotNo
		slot.IsAvailable = true
		slots.SlotInfo[slotNo] = slot
		initialCap++
	}
}

// ParkCarinNearestAvailSlot - Park the car to the nearest parking slot
func (slots *Slots) ParkCarinNearestAvailSlot(carData car.Car) (message string) {
	availSlot := checkAvailableSlot(slots.SlotInfo, len(slots.SlotInfo))
	if availSlot == 0 && slots.AvailableCapacity == 0 {
		message = "Sorry, parking lot is full"
		return
	}
	slot := Slot{}
	slot.IsAvailable = false
	slot.SlotNumber = availSlot
	slot.Car = carData
	slots.SlotInfo[availSlot] = slot
	slots.AvailableCapacity--
	message = fmt.Sprintf("%s: %d", "Allocated slot number", availSlot)
	return message
}

func checkAvailableSlot(slot map[int]Slot, lenght int) int {
	sortedData := sortByKey(slot, len(slot))
	for _, k := range sortedData {
		if slot[k].IsAvailable {
			return slot[k].SlotNumber
		}
	}
	return 0
}

func sortByKey(slot map[int]Slot, lenght int) []int {
	keys := make([]int, 0, lenght)
	for k := range slot {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

//LeaveFromParking method is for clear the car from the specific slot
func (slots *Slots) LeaveFromParking(slotNo int) (message string) {
	if slotNo == 0 {
		message = "Leave Status: There is no slot zero"
		return
	}
	if _, ok := slots.SlotInfo[slotNo]; !ok {
		message = "This Slot No is not found"
		return
	}

	slots.AvailableCapacity++
	slotData := slots.SlotInfo[slotNo]
	slotData.IsAvailable = true
	slotData.Car = car.Car{}
	slots.SlotInfo[slotNo] = slotData

	message = fmt.Sprintf("%s %d %s", "Slot number", slotNo, "is free")
	return message
}

// SlotStatus method helps to know the status of the parking slot
func (slots *Slots) SlotStatus() {

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 8, '\t', 0)
	defer w.Flush()

	slotData := slots
	slotInfo := slotData.SlotInfo
	fmt.Fprintf(w, "%s\t%s\t%s\t", "Slot No.", "Registration No", "Colour")
	sortedData := sortByKey(slotInfo, len(slotInfo))
	for _, k := range sortedData {
		if slotInfo[k].IsAvailable == true {
			continue
		}
		fmt.Fprintf(w, "\n%d\t%s\t%s\t", slotInfo[k].SlotNumber, slotInfo[k].Car.Number, slotInfo[k].Car.Color)
	}
}

// GetCarDetails Method helps to find the car Reg No and Slot No based on the color and Reg No
func (slots *Slots) GetCarDetails(cmd string, verify string) (message string) {
	isCarFound := false
	message = ""
	for _, val := range slots.SlotInfo {
		switch cmd {
		case "color":
			if val.Car.Color == verify {
				if isCarFound {
					message = fmt.Sprintf("%s,%s", message, val.Car.Number)
				} else {
					message = val.Car.Number
				}
				isCarFound = true
			}
		case "slotNo":
			if val.Car.Color == verify {
				if isCarFound {
					message = fmt.Sprintf("%s,%d", message, val.SlotNumber)
				} else {
					message = strconv.Itoa(val.SlotNumber)
				}
				isCarFound = true
			}
		case "regNo":
			if val.Car.Number == verify {
				if isCarFound {
					message = fmt.Sprintf("%s,%d", message, val.SlotNumber)
				} else {
					message = strconv.Itoa(val.SlotNumber)
				}
				isCarFound = true
			}
		}
	}
	if !isCarFound {
		message = "Not found"
	}
	return message
}
