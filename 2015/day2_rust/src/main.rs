use std::fs;

fn main() {
    let input = fs::read_to_string("input.txt").expect("Could not read file");
    let split: Vec<String> = input.split("\n").map(|s| s.to_string()).collect();

    let mut total_feet = 0;
    let mut ribbon_feet = 0;
    for sp in split {
        if sp == String::from("") {
            continue;
        }
        let present = process_line(&sp);
        total_feet += present.get_paper_square_feet();
        ribbon_feet += present.get_ribbon_sqft();
    }

    println!("Total feet: {}", total_feet);
    println!("Total ribbon: {}", ribbon_feet);

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
    fn get_ribbon_sqft(&self) -> i32 {
        let cubic = self.get_cubic_volume();
        let mut nums = [self.l, self.w, self.h];
        nums.sort();
        nums[0] + nums[0] + nums[1] + nums[1] + cubic
    }
    fn get_cubic_volume(&self) -> i32 {
        self.l * self.w * self.h
    }
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
    fn get_paper_square_feet(&self) -> i32 {
        2 * self.l * self.w + 2 * self.w * self.h + 2 * self.h * self.l + self.get_smallest_side()
    }
}

fn process_line(input: &String) -> Present {
    println!("Process line: {}", input);
    let parts: Vec<&str> = input.split("x").collect();

    let present = Present {
        l: parts[0].parse::<i32>().unwrap(),
        w: parts[1].parse::<i32>().unwrap(),
        h: parts[2].parse::<i32>().unwrap(),
    };
    present
}
