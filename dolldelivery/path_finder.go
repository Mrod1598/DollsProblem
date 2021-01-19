package dolldelivery

import (
	"fmt"
	"math"
)

const infinity = math.MaxUint32

// PathFinder must be implemented
type PathFinder func(startLocation, targetLocation string, streets []Street) (distance int, path []string)

// Street defines a connection between two houses
type Street struct {
	From     string
	To       string
	Distance int
}

//GoFindShortestPath finds the shortest path and distance using Dijkstra
var GoFindShortestPath = PathFinder(func(startLocation, targetLocation string, streets []Street) (distance int, path []string) {
	unqCount, unqHomes := goFindNumUnqStreets(streets)

	if unqCount <= 0 {
		fmt.Println("there are no homes in this neighborhood")
		return distance, path
	} else if found := find(unqHomes, startLocation); !found {
		fmt.Println("there are no homes in this neighborhood with the name: ", startLocation)
		return distance, path
	} else if found := find(unqHomes, targetLocation); !found {
		fmt.Println("there are no homes in this neighborhood with the name: ", targetLocation)
		return distance, path
	}

	prev, dist := Dijkstra(streets, startLocation, unqHomes)

	if dist[targetLocation] == infinity {
		fmt.Println("target Unreachable!")
		return dist[targetLocation], path
	}

	var tempPath []string

	//get the path from the target location back to the start
	for at := targetLocation; len(at) != 0; at = prev[at] {
		tempPath = append(tempPath, at)
	}
	//reverse the string to get it in the right order (startlocation to end location)
	for i := 0; i < len(tempPath)/2; i++ {
		j := len(tempPath) - i - 1
		tempPath[i], tempPath[j] = tempPath[j], tempPath[i]
	}

	return dist[targetLocation], tempPath
})

//Dijkstra this implements a shortest path algorithm that returns the values of how to get there with the first map and the distance of
//the cost to get there with the second map
func Dijkstra(neighborhood []Street, startLocation string, unqHomes []string) (map[string]string, map[string]int) {
	visited := make(map[string]bool)
	prev := make(map[string]string)
	dist := make(map[string]int)

	//initialize all variables with applicable values
	for _, home := range unqHomes {
		visited[home] = false
		dist[home] = infinity
	}

	//init priority queue
	pq := make(PriorityQueue, 0)

	//start the queue off with the start location.
	firstItem := &Item{
		Value:    startLocation,
		Priority: 0,
	}
	pq.Push(firstItem)
	dist[startLocation] = 0
	//begin Djikstras (with some modifications for SPEEDDD, mostly so you don't have to check all values, most of the time)
	for pq.Len() != 0 {
		//get next item in Queue
		currentItem := pq.Pop().(*Item)
		visited[currentItem.Value] = true

		//get all the neighbors connected to current home (node)
		for _, street := range neighborhood {
			//skip item if we've visited the home before or if it's not connected to the our current home.
			if street.From != currentItem.Value || visited[street.To] {
				continue
			}

			//calculate the distance from current home to the next home and see if it's lower than the current distance we already have for it
			newDist := dist[currentItem.Value] + street.Distance
			if newDist < dist[street.To] {
				prev[street.To] = currentItem.Value
				dist[street.To] = newDist

				newItem := &Item{
					Value:    street.To,
					Priority: 0,
				}
				pq.Push(newItem)
			}
		}
	}
	return prev, dist
}

//GoFindNumUnqStreets finds the amount of unique streets in the slice
func goFindNumUnqStreets(streets []Street) (numOfStreets int, unqHomes []string) {
	numOfStreets = 0
	keys := make(map[string]bool)
	for _, currentStreet := range streets {
		if _, value := keys[currentStreet.From]; !value {
			keys[currentStreet.From] = true
			numOfStreets++
			unqHomes = append(unqHomes, currentStreet.From)
		} else if _, value := keys[currentStreet.To]; !value {
			keys[currentStreet.To] = true
			numOfStreets++
			unqHomes = append(unqHomes, currentStreet.To)
		}
	}
	return numOfStreets, unqHomes
}

//Find checks whether or not a value exists in a slice
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
