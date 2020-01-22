package parking

import (
	slot "parking_lot/functional_spec/slot"
	"testing"

	constant "parking_lot/functional_spec/constant"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFileCmd(t *testing.T) {
	slots := &slot.Slots{}
	//Test Case 1 - Creating the parking lot with 4 slots
	expectedResult := "Created a parking lot with 4 slots"
	cmd := constant.CreateParking
	data := "4"
	actualResult := ""
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 2 - Creating the parking lot with 0 slots
	expectedResult = "Enter Valid Parking Capacity Count"
	cmd = constant.CreateParking
	data = "0"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 3 - Creating the parking lot with (-ve) slots
	expectedResult = "Enter Valid Parking Capacity Count"
	cmd = constant.CreateParking
	data = "-1"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 4 -  Park the car to nearest slot park KA-01-HH-1234 White -
	expectedResult = "Allocated slot number: 1"
	cmd = constant.Park
	data = "KA-01-HH-1234 White"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Park 3 more cars to nearest slot
	expectedResult = "Allocated slot number: 2"
	data = "KA-01-HH-9999 Red"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	expectedResult = "Allocated slot number: 3"
	data = "KA-01-BB-0001 Black"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	expectedResult = "Allocated slot number: 4"
	data = "KA-01-HH-7777 Red"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 5 - All the slots are occupied, if another car comes to the parking lot
	expectedResult = "Sorry, parking lot is full"
	data = "KA-01-HH-7787 Green"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 6 -  Leave the car from 4th slot
	expectedResult = "Slot number 4 is free"
	cmd = constant.Leave
	data = "4"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 7 - Find the registered car No by color
	expectedResult = "KA-01-HH-1234"
	cmd = constant.RegNoForCarWcolor
	data = "White"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 8 - Find the slot No by color
	expectedResult = "1"
	cmd = constant.SlotNoCarColor
	data = "White"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 9 - Find the slot no by reg no
	expectedResult = "3"
	cmd = constant.SlotNoCarRegNo
	data = "KA-01-BB-0001"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

	//Test Case 10 - Find the non-parking car no
	expectedResult = "Not found"
	cmd = constant.SlotNoCarRegNo
	data = "KA-01-BB-0008"
	slots, actualResult = ReadFileCmd(cmd, data, slots)
	assert.Equal(t, expectedResult, actualResult)
	require.Equal(t, expectedResult, actualResult)

}
