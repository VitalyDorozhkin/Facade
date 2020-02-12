package facade

import "testing"

var values = [16]byte{0, 1, 2, 3, 0, 1, 4, 9, 8, 9, 10, 11, 12, 13, 14, 15}

func TestComputerFacade_Start(t *testing.T) {
	expected := 8
	computer := NewComputerFacade()
	res := computer.Start()
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestComputerFacade_Swap(t *testing.T) {
	c := NewComputerFacade()

	for i := 0; i < 16; i++ {
		c.hd.write(i, byte(i))
	}

	for i := 0; i < 4; i++ {
		c.ram.write(i, byte(i*i))
	}

	for i := 0; i < 4; i++ {
		c.Swap(i, i+4)
	}

	for i := 0; i < 16; i++ {
		res := c.hd.read(i)
		if res != values[i] {
			t.Errorf("On pointer %d expected %d, got %d", i, values[i], res)
		}
	}
}

func TestComputerFacade(t *testing.T) {

	c := NewComputerFacade()
	testsWrite := []struct {
		data         byte
		wantPosition int
	}{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{1, 6},
		{1, 7},
		{1, 8},
		{1, 9},
		{1, 10},
	}
	testsFree := []struct {
		data         byte
		wantPosition int
	}{
		{1, 5},
		{1, 11},
		{1, 12},
		{1, 13},
	}
	for _, tt := range testsWrite {
		if gotPosition := c.Write(tt.data); gotPosition != tt.wantPosition {
			t.Errorf("Write() = %v, want %v", gotPosition, tt.wantPosition)
		}
	}
	c.hd.cleanAll()
	for _, tt := range testsWrite {
		if gotPosition := c.Write(tt.data); gotPosition != tt.wantPosition {
			t.Errorf("Write() = %v, want %v", gotPosition, tt.wantPosition)
		}
	}
	c.hd.free(5)
	for _, tt := range testsFree {
		if gotPosition := c.Write(tt.data); gotPosition != tt.wantPosition {
			t.Errorf("Write() = %v, want %v", gotPosition, tt.wantPosition)
		}
	}
}
