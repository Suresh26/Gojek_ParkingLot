package parking

import (
	"log"
	"os"
	car "parking_lot/functional_spec/car"
	constant "parking_lot/functional_spec/constant"
	slot "parking_lot/functional_spec/slot"
	"strconv"
	"strings"
)

//ReadFileCmd - Run the functionality based on the commands
func ReadFileCmd(command string, data string, slotData *slot.Slots) (*slot.Slots, string) {
	var slotcmd slot.SlotCommand
	message := ""
	slotcmd = slotData
	switch strings.TrimSpace(command) {
	case constant.CreateParking:
		parkingCapacity, err := strconv.Atoi(data)
		if err != nil {
			log.Fatal(err)
		}
		message = slotcmd.CreateSlots(parkingCapacity)
	case constant.Park:
		carDetails := strings.SplitN(data, " ", 2)
		carInfo := car.Car{}
		carInfo.Number = carDetails[0]
		carInfo.Color = carDetails[1]
		message = slotcmd.ParkCarinNearestAvailSlot(carInfo)
	case constant.Leave:
		slotNo, err := strconv.Atoi(data)
		if err != nil {
			log.Fatal(err)
		}
		message = slotcmd.LeaveFromParking(slotNo)
	case constant.RegNoForCarWcolor:
		message = slotcmd.GetCarDetails("color", data)
	case constant.SlotNoCarColor:
		message = slotcmd.GetCarDetails("slotNo", data)
	case constant.SlotNoCarRegNo:
		message = slotcmd.GetCarDetails("regNo", data)
	case constant.Status:
		slotcmd.SlotStatus()
	case constant.Exit:
		os.Exit(0)
	default:
		message = "Command is not found"
	}
	return slotData, message
}
