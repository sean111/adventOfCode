use std::fs::File;
use std::io::{stdout, BufReader, Result};
use std::io::prelude::*;

fn main() {
    let p1res = part1("data/data.txt").unwrap();
    writeln!(stdout(), "Part1: {}", p1res).unwrap();
}

pub fn part1(file: &str) -> Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut safe_count: i32 = 0;
    reader.read_to_string(&mut contents)?;
    for line in contents.lines() {
        let check: bool = check_line(line);
        if check {
            writeln!(stdout(), "Safe Line: {}", line)?;
            safe_count += 1;
        }
    }
    Ok(safe_count)
}

pub fn check_line(line: &str) -> bool {
    let parts = line.split(" ");
    // let mut data: Vec<i32> = Vec::new();
    let mut count: i32 = 0;
    let mut previous: i32 = 0;
    // 1 = up, -1 = down
    let mut direction: i8 = 0;
    for i in parts {
        let number = i.trim().parse::<i32>().unwrap();
        let mut temp_direction: i8 = 0;
        if count > 0 {
            let temp = number - previous;
            if temp > 0 {
                temp_direction = 1
            } else if temp < 0 {
                temp_direction = -1
            }

            if temp_direction == 0 {
                return false;
            }

            if direction != 0 {
                if temp_direction != direction || temp_direction == 0 {
                    return false;
                }
            } else {
                direction = temp_direction;
            }

            if temp.abs() > 3 {
                return false;
            }
        }
        count += 1;
        previous = number;
    }
    true
}

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_part1() {
        let result = part1("data/test.txt").unwrap();
        assert_eq!(result, 2);
    }
}