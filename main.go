package main

import (
	"errors"
	"fmt"
	"os"

	tkgpxgo "github.com/tkrajina/gpxgo/gpx"
)

func main() {
	gpxData, err := tkgpxgo.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	if gpxData.GetTrackPointsNo() == 0 {
		panic(errors.New("there is not track!"))
	}
	const metresInKM = 1000.00
	fmt.Printf("Length2D is %.2f km\n", gpxData.Length2D()/metresInKM)
	fmt.Printf("Length3D is %.2f km\n", gpxData.Length3D()/metresInKM)
}
