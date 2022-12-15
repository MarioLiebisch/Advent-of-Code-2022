package sensors

type Sensor struct {
	X, Y                     int
	bX, bY                   int
	R                        int
	top, bottom, left, right int
}

type Tunnels struct {
	sensors []Sensor
}
