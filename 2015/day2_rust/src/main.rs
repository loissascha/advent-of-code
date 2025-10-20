use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Could not read file");
    let split: Vec<String> = input.split("\n").map(|s| s.to_string()).collect();

    let mut total_feet = 0;
    for sp in split {
        if sp == String::from("") {
            continue;
        }
        total_feet += process_line(&sp);
    }

    println!("Total feet: {}", total_feet);

    // let test_line = String::from("2x3x4");
    // let sqf = process_line(&test_line);
    // println!("Tets Sqf: {}", sqf);
}

struct Present {
    l: i32,
    w: i32,
    h: i32,
}

impl Present {
    fn get_smallest_side(&self) -> i32 {
        let s1 = self.l * self.w;
        let s2 = self.w * self.h;
        let s3 = self.h * self.l;
        let mut smallest = s1;
        if s2 < smallest {
            smallest = s2;
        }
        if s3 < smallest {
            smallest = s3;
        }
        smallest
    }
    fn get_square_feet(&self) -> i32 {
        2 * self.l * self.w + 2 * self.w * self.h + 2 * self.h * self.l + self.get_smallest_side()
    }
}

fn process_line(input: &String) -> i32 {
    println!("Process line: {}", input);
    let parts: Vec<&str> = input.split("x").collect();

    let present = Present {
        l: parts[0].parse::<i32>().unwrap(),
        w: parts[1].parse::<i32>().unwrap(),
        h: parts[2].parse::<i32>().unwrap(),
    };
    present.get_square_feet()
}
