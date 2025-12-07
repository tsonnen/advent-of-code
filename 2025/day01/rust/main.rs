extern crate regex;

use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

use regex::Regex;

fn read_input_file() -> io::Result<Vec<String>> {
    BufReader::new(File::open("input.txt")?).lines().collect()
}

fn modulo(x: i32, y: i32) -> i32 {
   let r = x % y;
   if r < 0 {
      r + y
   } else {
      r
   }
}

fn main(){
   let lines = read_input_file().expect("Failed to read input file");
   let times_at_zero = count_times_at_zero(&lines);
   let times_past_zero = count_times_past_zero(&lines);

   println!("Times at zero: {}", times_at_zero);
   println!("Times past zero: {}", times_past_zero);
}

fn count_times_at_zero(lines: &Vec<String>) -> i32 {
    let mut current_value = 50;
    let mut times_at_zero = 0;
    let re = Regex::new(r"(L|R)(\d+)").unwrap();

    for line in lines {
        let captures = re.captures(line).unwrap();

        match &captures[1] {
            "L" => current_value += &captures[2].parse::<i32>().unwrap(),
            "R" => current_value -= &captures[2].parse::<i32>().unwrap(),
            _ => panic!("Unexpected direction"),
        }

        current_value = current_value % 100;
        if current_value == 0 {
            times_at_zero += 1;
        }
    }
    times_at_zero
}

fn count_times_past_zero(lines: &Vec<String>) -> i32 {
    let mut current_value = 50;
    let mut times_past_zero = 0;
    let re = Regex::new(r"(L|R)(\d+)").unwrap();

    for line in lines {
        let captures = re.captures(line).unwrap();
        let previous_value = current_value;

        let spaces_turned = &captures[2].parse::<i32>().unwrap();

        match &captures[1] {
            "L" => {
                current_value += spaces_turned;
            },
            "R" => {
                current_value -= spaces_turned;
            },
            _ => panic!("Unexpected direction"),
        }

        times_past_zero += current_value.abs() / 100 + if previous_value != 0 && current_value <= 0 { 1 } else { 0 };

        current_value = modulo(current_value, 100);

    }
    times_past_zero
}