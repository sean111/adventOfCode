use std::fs::File;
use std::io::{stdout, BufReader, Result};
use std::io::prelude::*;
use regex::Regex;

fn main() {
    writeln!(stdout(), "Part 1: {}", part1("data/data.txt").unwrap()).unwrap();
    writeln!(stdout(), "Part 2: {}", part2("data/data.txt").unwrap()).unwrap();
}

pub fn part1(file: &str) -> Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    reader.read_to_string(&mut contents)?;
    let regex = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").unwrap();
    let mut total:i32 = 0;
    regex.captures_iter(contents.as_str()).for_each(|i| {
        total += i.get(1).unwrap().as_str().parse::<i32>().unwrap() * i.get(2).unwrap().as_str().parse::<i32>().unwrap()
    });
    Ok(total)
}

pub fn part2(file: &str) -> Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    reader.read_to_string(&mut contents)?;
    let regex = Regex::new(r"(?m)mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)").unwrap();
    let mut total:i32 = 0;
    let mut calculate: bool = true;
    regex.captures_iter(contents.as_str()).for_each(|i| {
        if i.get(0).unwrap().as_str().contains("do()") {
            calculate = true;
        } else if i.get(0).unwrap().as_str().contains("don't()") {
            calculate = false;
        }

        if calculate && i.get(0).unwrap().as_str().contains("mul") {
            total += i.get(1).unwrap().as_str().parse::<i32>().unwrap() * i.get(2).unwrap().as_str().parse::<i32>().unwrap()
        }
    });
    Ok(total)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_part1() {
        assert_eq!(part1("data/test.txt").unwrap(), 161);
    }
    #[test]
    fn test_part2() {
        assert_eq!(part2("data/test_2.txt").unwrap(), 48);
    }
}