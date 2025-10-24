use std::fs;

const GRID_SIZE: usize = 1000;

#[derive(Clone, Copy)]
struct Light {
    status: bool,
}

struct Grid {
    items: [[Light; GRID_SIZE]; GRID_SIZE],
}

struct Coords {
    x: usize,
    y: usize,
}

impl Grid {
    fn print_lit(&self) {
        let mut lit: u32 = 0;
        for y in self.items {
            for x in y {
                if x.status {
                    lit += 1;
                }
            }
        }
        println!("lit: {}", lit);
    }

    fn turn_on(&mut self, start_x: usize, start_y: usize, stop_x: usize, stop_y: usize) {
        // +1 because we want <= not <
        for y in start_y..stop_y + 1 {
            for x in start_x..stop_x + 1 {
                self.items[x][y].status = true;
            }
        }
    }

    fn turn_off(&mut self, start_x: usize, start_y: usize, stop_x: usize, stop_y: usize) {
        for y in start_y..stop_y + 1 {
            for x in start_x..stop_x + 1 {
                self.items[x][y].status = false;
            }
        }
    }

    fn toggle(&mut self, start_x: usize, start_y: usize, stop_x: usize, stop_y: usize) {
        for y in start_y..stop_y + 1 {
            for x in start_x..stop_x + 1 {
                self.items[x][y].status = !self.items[x][y].status;
            }
        }
    }

    fn process_line(&mut self, line: String) {
        if line.starts_with("turn on") {
            let split = line.split_once("turn on ").expect("worked");
            let line = split.1;
            let (start_coords, stop_coords) = split_through_str(line.to_string());
            self.turn_on(start_coords.x, start_coords.y, stop_coords.x, stop_coords.y);
        } else if line.starts_with("turn off") {
            let split = line.split_once("turn off ").expect("worked");
            let line = split.1;
            let (start_coords, stop_coords) = split_through_str(line.to_string());
            self.turn_off(start_coords.x, start_coords.y, stop_coords.x, stop_coords.y);
        } else if line.starts_with("toggle") {
            let split = line.split_once("toggle ").expect("worked");
            let line = split.1;
            let (start_coords, stop_coords) = split_through_str(line.to_string());
            self.toggle(start_coords.x, start_coords.y, stop_coords.x, stop_coords.y);
        }
    }
}

fn split_through_str(input: String) -> (Coords, Coords) {
    let parts = input
        .split_once(" through ")
        .expect("split thorugh str not wokred");
    let start_coords = split_coords(parts.0.to_string());
    let stop_coords = split_coords(parts.1.to_string());
    (start_coords, stop_coords)
}

fn split_coords(input: String) -> Coords {
    let parts = input.split_once(",").expect("split coords not work");
    let x = parts.0.parse::<usize>().unwrap();
    let y = parts.1.parse::<usize>().unwrap();
    Coords { x: x, y: y }
}

fn main() {
    let lines: Vec<String> = fs::read_to_string("input.txt")
        .expect("cant read file")
        .split("\n")
        .map(|s| s.to_string())
        .collect();

    let mut grid = Grid {
        items: [[Light { status: false }; GRID_SIZE]; GRID_SIZE],
    };

    for line in lines {
        if line.trim() == "" {
            continue;
        }
        println!("line: {}", line);
        grid.process_line(line);
    }

    grid.print_lit();
}
