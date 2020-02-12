package facade

import "testing"

func TestComputerFacade_Start(t *testing.T) {
	expected := 8
	computer := NewComputerFacade()
	res := computer.Start()
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}


var values = [16]byte{0, 1, 2, 3, 0, 1, 4, 9, 8, 9, 10, 11, 12, 13, 14, 15}



func TestComputerFacade_Swap(t *testing.T) {
	computer := NewComputerFacade()

	for i := 0; i < 16; i++ {
		computer.hd.write(i, byte(i))
	}

	for i := 0; i < 4; i++ {
		computer.ram.write(i, byte(i * i))
	}


	for i := 0; i < 4; i++ {
		computer.Swap(i, i + 4)
	}

	for i := 0; i < 16; i++ {
		res := computer.hd.read(i)
		if res != values[i] {
			t.Errorf("On pointer %d expected %d, got %d", i, values[i], res)
		}
	}

}