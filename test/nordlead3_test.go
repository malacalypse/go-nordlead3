package nordlead3_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/malacalypse/nordlead3"
	"testing"
)

const (
	invalidPerformance       = "F0337F092900004F72636865737472612020202020484E000000000000000000000000000000000078001E007008000202014000000000760007436170297402050241605060080000000000DEADBEEF000000000000000C0621192D466205647032596C764B3C402024134A372349526E335C6000000000000000000002191A62305C6E3308000000000000000000004C32582C443B256861390804020100404827000C60287F6B5900190A40203F5E1140667F42402000003F403213000000000000131A20155A000000096807485858560F0022746132027404404000002B120F7F7C000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000E000000000000000000000000000000000000000000000000000700004045204000223F262822407F45383E4F002C252B0F0063006439640000000000000D454B0000000000100F0000000C1E00410830600160012200000217001F7F7800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001D200000014A000000000110000007486002080000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000E000033000000384800500001017E3336031B7E0A010660017E0148680000000000000001003A6800000000201E226262583C010B5B05480B50120200000228003F7F720000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000138002A00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400038000000000000000000000000000000000000000000000000001C00006618017C0314007C0002037C2A4C06377C1400750C017E67111E600000000000006A005D5000000000403C454545307802000545400340405800000C50007F7F68000000000000000000000000000000000000000000000000002B00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004200100007000000000000000000000000000000000000000000000000000380400F7"
	invalidProgram           = "F0337F0921000057656C636F6D650000000000000000000B0000000000000000000000000000000078001E00063000080009726F7A03701F681640116C44200A20104B00190B2000000000000006100A79000004040403642C2C2B0740102B2C18013A016028010045000F7F7E000000DEADBEEF000000000000001E40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000170000000000000000000000000000000000000000000000000000000000800070000000000000000000000000000000000000000000000000002601440F7"
	validPerformance         = "F0337F092901024F72636865737472612020202020484E000000000000000000000000000000000078001E00700800020201400000000076000743617029740205024160506008000000000000000000000000000000000C0621192D466205647032596C764B3C402024134A372349526E335C6000000000000000000002191A62305C6E3308000000000000000000004C32582C443B256861390804020100404827000C60287F6B5900190A40203F5E1140667F42402000003F403213000000000000131A20155A000000096807485858560F0022746132027404404000002B120F7F7C000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000E000000000000000000000000000000000000000000000000000700004045204000223F262822407F45383E4F002C252B0F0063006439640000000000000D454B0000000000100F0000000C1E00410830600160012200000217001F7F7800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001D200000014A000000000110000007486002080000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000E000033000000384800500001017E3336031B7E0A010660017E0148680000000000000001003A6800000000201E226262583C010B5B05480B50120200000228003F7F720000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000138002A00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000400038000000000000000000000000000000000000000000000000001C00006618017C0314007C0002037C2A4C06377C1400750C017E67111E600000000000006A005D5000000000403C454545307802000545400340405800000C50007F7F68000000000000000000000000000000000000000000000000002B00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004200100007000000000000000000000000000000000000000000000000000380400F7"
	validPerformanceBank     = 1
	validPerformanceLocation = 2
	validPerformanceName     = "Orchestra     HN"
	validPerformanceVersion  = 1.20
	validProgram             = "F0337F0921020257656C636F6D650000000000000000000B0000000000000000000000000000000078001E00063000080009726F7A03701F681640116C44200A20104B00190B2000000000000006100A79000004040403642C2C2B0740102B2C18013A016028010045000F7F7E00000000000000000000000000001E40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000170000000000000000000000000000000000000000000000000000000000800070000000000000000000000000000000000000000000000000002601440F7"
	validProgramBank         = 2
	validProgramLocation     = 2
	validProgramName         = "Welcome         "
	validProgramVersion      = 1.20
)

func TestLoadValidPerformanceFromSysex(t *testing.T) {
	memory := new(nordlead3.PatchMemory)
	sysex := stringToBytes(validPerformance)

	err := memory.LoadFromSysex(sysex)
	if err != nil {
		t.Errorf("Expected clean load from valid sysex. Got: %q", err)
	}

	expectedLocation := memory.Performances[validPerformanceBank][validPerformanceLocation]

	if expectedLocation == nil {
		t.Errorf("Did not load performance into expected location!")
	}

	loadedPerformance := expectedLocation.Performance

	if loadedPerformance == nil {
		t.Errorf("Did not load performance into expected location!")
	}

	if expectedLocation.PrintableName() != validPerformanceName {
		t.Errorf("Did not correctly compute the performance name.")
	}

	if expectedLocation.Version != validPerformanceVersion {
		t.Errorf("Did not correctly compute the performance version.")
	}
}

func TestLoadInvalidPerformanceFromSysex(t *testing.T) {
	memory := new(nordlead3.PatchMemory)
	sysex := stringToBytes(invalidPerformance)

	err := memory.LoadFromSysex(sysex)
	if err == nil {
		t.Errorf("Expected error from invalid sysex")
	}

	expectedLocation := memory.Performances[validPerformanceBank][validPerformanceLocation]

	if expectedLocation != nil {
		loadedPerformance := expectedLocation.Performance

		if loadedPerformance != nil {
			t.Errorf("Loaded invalid performance into memory!")
		}
	}
}

func TestLoadProgramFromSysex(t *testing.T) {
	memory := new(nordlead3.PatchMemory)
	sysex := stringToBytes(validProgram)

	err := memory.LoadFromSysex(sysex)
	if err != nil {
		t.Errorf("Expected clean load from valid sysex. Got: %q", err)
	}

	expectedLocation := memory.Programs[validProgramBank][validProgramLocation]
	loadedProgram := expectedLocation.Program

	if loadedProgram == nil {
		t.Errorf("Did not load program into expected location!")
	}

	if expectedLocation.PrintableName() != validProgramName {
		t.Errorf("Did not correctly compute the program name.")
	}

	if expectedLocation.Version != validProgramVersion {
		t.Errorf("Did not correctly compute the program version.")
	}
}

func TestDumpProgramToSysex(t *testing.T) {
	memory := new(nordlead3.PatchMemory)
	inputSysex := stringToBytes(validProgram)
	err := memory.LoadFromSysex(inputSysex)
	if err != nil {
		t.Errorf("Test sysex seems incorrect, need valid sysex to test dumping: %q", err)
	}

	outputSysex, err := memory.DumpProgram(validProgramBank, validProgramLocation)
	if err != nil {
		t.Errorf("Error dumping program: %q", err)
	}

	// Use string format for quick equal comparison
	if bytesToString(*outputSysex) != validProgram {
		location, explanation := locationOfDifference(inputSysex, *outputSysex)
		fmt.Printf("Input:  %x\n", inputSysex)
		fmt.Printf("Output: %x\n", *outputSysex)
		t.Errorf("Dumped sysex does not match input at offset %d: %q", location, explanation)
	}
}

func stringToBytes(s string) []byte {
	var result []byte
	fmt.Sscanf(s, "%X", &result)
	return result
}

func bytesToString(b []byte) string {
	var result string
	fmt.Sscanf(string(b), "%X", &result)
	return result
}

func locationOfDifference(b1, b2 []byte) (int, error) {
	r1 := bytes.NewReader(b1)
	r2 := bytes.NewReader(b2)
	i := 0

	for {
		c1, err1 := r1.ReadByte()
		c2, err2 := r2.ReadByte()
		if c1 == c2 && err1 == err2 {
			// skip
		} else {
			minIndex := max(0, i-5)
			maxIndex := min(min(i+5, len(b1)), len(b2))
			fmt.Printf("S1 (%d) / S2 (%d) / i (%d) : (%d, %d)\n", len(b1), len(b2), i, minIndex, maxIndex)
			explanation := fmt.Sprintf("Bytes 1: %x^%x | Bytes 2: %x^%x", b1[minIndex:i], b1[i:maxIndex], b2[minIndex:i], b2[i:maxIndex])
			return i, errors.New(explanation)
		}
		i++
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}