# Delivery Slot Allocation System (Golang)

## Overview
The Delivery Slot Allocation System is a Golang-based backend application that allocates a limited number of delivery slots to multiple concurrent requests.
The project demonstrates safe handling of shared resources using goroutines, mutexes, and synchronization, ensuring correct allocation without race conditions under high concurrency.

This project focuses on concurrency fundamentals commonly required in backend systems.

## Tech Stack
- Language: Go (Golang)
- Concurrency: Goroutines
- Synchronization: sync.Mutex, sync.WaitGroup
- Frameworks: None (Core Go)

## Features
- Handles multiple concurrent slot requests
- Ensures thread-safe slot allocation
- Prevents race conditions
- Handles slot exhaustion correctly
- Simple and readable concurrency-focused design

## Project Structure
delivery-slot-allocation/
├── main.go
└── README.md

## How to Run
1. Clone the repository
2. Run the application using:
   go run main.go

## Author
Anshu Kumar Roy
