package nordlead3

import (
	"testing"
)

func TestDumpPerformanceSysex(t *testing.T) {
	memory := new(PatchMemory)
	inputSysex := validPerformanceSysex(t)
	inputSysexStruct, err := parseSysex(inputSysex)
	if err != nil {
		t.Errorf("Test sysex seems incorrect, need valid sysex to test dumping: %q", err)
	}
	performanceSysex := inputSysexStruct.rawBitstream()

	helperLoadFromSysex(t, memory, inputSysex)

	p, err := memory.get(validPerformanceRef)
	performance := p.(*Performance)
	if err != nil {
		t.Errorf("Error retrieving performance: %q", err)
	}

	outputSysex, err := performance.data.dumpSysex()
	if err != nil {
		t.Errorf("Error dumping performance: %q", err)
	}

	// Compare the decoded data for easier debugging
	decodedPS := unpackSysex(performanceSysex)
	decodedOS := unpackSysex(*outputSysex)
	binaryExpectEqual(t, &decodedPS, &decodedOS)
}
