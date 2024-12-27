# recipes

This project uses recipe information fetched from [Dummy JSON](https://dummyjson.com/recipes?limit=0) and displays the information in a visually friendly manner on a webpage.

## Project Overview

![](/recipes/static/img/preview.png)

## Prerequisites

[Go](https://go.dev/doc/install) version 1.22 or higher

## Installation

- Clone repository to your local environment and navigate to recipes directory using the below terminal command:

  ```bash
  git clone https://github.com/johneliud/dummy-json-api-manipulation.git

  cd dummy-json-api-manipulation/recipes
  ```

## Usage

- The program can be run using either of the below commands:

  ```bash
  go run .
  ```

  OR

  ```bash
  go run . [SPECIFY-PORT]
  ```

**NOTE:** A message is logged indicating the port where the program's server is running. `8080` is the default port.

### Features

- Featured recipe highlight that randomly changes after every 15 seconds.

## Contributions

Contributions towards improving the project are welcome. Create a pull request to submit your proposed change. Feel free to reach [out](johneliud4@gmail.com)