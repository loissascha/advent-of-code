use std::fs;

fn main() {
    println!("Hello, world!");

    let testInput = String::from("^v^v^v^v^v");
    let d = houses_delivered(testInput);

    println!("The t is {}", d);

    let input = fs::read_to_string("input.txt").expect("Can't read file.");
    println!("Input: {}", input);
    let dd = houses_delivered(input);

    println!("The dd is {}", dd);
}

struct House {
    pos_x: i32,
    pos_y: i32,
    presents: u32,
}

fn houses_delivered(input: String) -> usize {
    let mut houses: Vec<House> = Vec::new();

    let mut current_pos_x: i32 = 0;
    let mut current_pos_y: i32 = 0;
    houses.push(House {
        pos_x: 0,
        pos_y: 0,
        presents: 1,
    });

    let chars: Vec<char> = input.chars().collect();
    for char in chars {
        if char == '>' {
            current_pos_x += 1;
        }
        if char == '<' {
            current_pos_x -= 1;
        }
        if char == '^' {
            current_pos_y += 1;
        }
        if char == 'v' {
            current_pos_y -= 1;
        }
        let mut found = false;
        for h in houses.iter_mut() {
            if h.pos_x == current_pos_x && h.pos_y == current_pos_y {
                found = true;
                h.presents += 1;
            }
        }
        if !found {
            let house = House {
                pos_x: current_pos_x,
                pos_y: current_pos_y,
                presents: 1,
            };
            houses.push(house);
        }
    }

    houses.len()
}
