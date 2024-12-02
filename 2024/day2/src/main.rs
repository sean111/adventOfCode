use std::fs::File;
use std::io::{stdout, BufReader, Result};
use std::io::prelude::*;

fn main() {
    let p1res = part1("data/data.txt").unwrap();
    writeln!(stdout(), "Part1: {}", p1res).unwrap();
    let p2res = part2("data/data.txt").unwrap();
    writeln!(stdout(), "Part2: {}", p2res).unwrap();
}

pub fn part1(file: &str) -> Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut safe_count: i32 = 0;
    reader.read_to_string(&mut contents)?;
    for line in contents.lines() {
        let parts = line.split(" ");
        let mut data: Vec<i32> = Vec::new();
        for i in parts {
            data.push(i.parse::<i32>().unwrap());
        }
        let check: bool = check_line(data);
        if check {
            writeln!(stdout(), "Safe Line: {}", line)?;
            safe_count += 1;
        }
    }
    Ok(safe_count)
}

pub fn part2(file: &str) -> Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut safe_count: i32 = 0;
    reader.read_to_string(&mut contents)?;
    for line in contents.lines() {
        let parts = line.split(" ");
        let mut data: Vec<i32> = Vec::new();
        for i in parts {
            data.push(i.parse::<i32>().unwrap());
        }
        let check: bool = dampner(data);
        if check {
            writeln!(stdout(), "Safe Line: {}", line)?;
            safe_count += 1;
        }
    }
    Ok(safe_count)
}

pub fn check_line(line: Vec<i32>) -> bool {
    let mut count: i32 = 0;
    let mut previous: i32 = 0;
    // 1 = up, -1 = down
    let mut direction: i8 = 0;
    for number in line {
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

pub fn dampner(line: Vec<i32>) -> bool {
    let count = line.len();
    for i in 0..count {
        let mut temp_line = line.to_vec();
        temp_line.remove(i);
        let check = check_line(temp_line);
        if check {
            return true;
        }
    }
    false
}

#[cfg(test)]

mod tests {
    use super::*;
    #[test]
    fn test_part1() {
        let result = part1("data/test.txt").unwrap();
        assert_eq!(result, 2);
    }
    #[test]
    fn test_part2() {
        let result = part2("data/test.txt").unwrap();
        assert_eq!(result, 4);
    }
}
