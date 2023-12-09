# Advent of Code in Golang

![Advent of Code](https://img.shields.io/badge/Advent%20of%20Code-2023-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/10wpressure/advent-of-code)](https://goreportcard.com/report/github.com/10wpressure/advent-of-code)

This repository contains my solutions for Advent of Code challenges implemented in Golang.

## Table of Contents

- [Advent of Code in Golang](#advent-of-code-in-golang)
  - [Table of Contents](#table-of-contents)
  - [Folder Structure](#folder-structure)
  - [Getting Started](#getting-started)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [License](#license)

## Folder Structure

The repository is organized by year, with each year containing the solutions for each day:

- **`aoc2023/`**: Solutions for Advent of Code 2023.
  - **`day1/`**: Contains the solution for Day 1.
  - **`day2/`**: Contains the solution for Day 2.
    ...

Each day's directory typically includes:

- **`input.txt`**: The input data for the specific day.
- **`main.go`**: The Golang code implementing the solution.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/10wpressure/advent-of-code.git

## Advent of Code Input Data Downloader

To simplify the process of downloading Advent of Code input data, you can use the [Advent of Code downloader](https://github.com/GreenLightning/advent-of-code-downloader) CLI utility.
Follow the instructions below to set it up:

1. Install the downloader using `go get`:

   ```bash
   go install github.com/GreenLightning/advent-of-code-downloader/aocdl@latest

2. Set your session cookie as a command line parameter:
    ```bash
    aocdl -session-cookie 0123456789...abcdef
    ```
Or create a configuration file named .aocdlconfig in your home directory or in the current directory and add the session-cookie key:
    ```bash
    {
    	"session-cookie": "0123456789...abcdef"
    }
    ```
3. Navigate to your current problem (i.e. cd aoc2023/day3) and use command:

    ```bash
    aocdl -year 2015 -day 1
