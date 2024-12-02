use std::fs::File;
use std::io::{stdout, BufReader};
use std::io::prelude::*;
use regex::Regex;
use std::collections::HashMap;

fn main() {
    let p1res = part1("data/data.txt").unwrap();
    let p2res = part2("data/data.txt").unwrap();
    writeln!(stdout(), "Part1: {}", p1res).unwrap();
    writeln!(stdout(), "Part2: {}", p2res).unwrap();
}

pub fn get_values(contents: String) -> Vec<(i32, i32)> {
    let regex = Regex::new(r"(\d+)\s\s\s(\d+)").unwrap();
    let values: Vec<(i32, i32)> = regex.captures_iter(contents.as_str()).map(|numbers| {
        let l1 = numbers.get(1).unwrap().as_str().parse::<i32>().unwrap();
        let l2 = numbers.get(2).unwrap().as_str().parse::<i32>().unwrap();
        (l1, l2)
    }).collect();
    values
}

pub fn part1(file: &str) -> std::io::Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut list1:Vec<i32> = Vec::new();
    let mut list2:Vec<i32> = Vec::new();
    let mut total: i32 = 0;
    reader.read_to_string(&mut contents)?;
    writeln!(stdout(), "{}", contents)?;
    let values = get_values(contents);
    writeln!(stdout(), "values {:?}", values)?;

    values.into_iter().for_each(|(l1, l2)| {
        list1.push(l1);
        list2.push(l2);
    });
    list1.sort();
    list2.sort();

    let count = list1.len();

    for x in 0..count {
        total += (list1[x] - list2[x]).abs();
    }

    writeln!(stdout(), "1:list1 {:?}", list1)?;
    writeln!(stdout(), "1:list2 {:?}", list2)?;

    Ok(total)
}

pub fn part2(file: &str) -> std::io::Result<i32> {
    let fp = File::open(file)?;
    let mut reader = BufReader::new(fp);
    let mut contents = String::new();
    let mut list1:Vec<i32> = Vec::new();
    let mut list2:HashMap<i32, i32> = HashMap::new();
    let mut total: i32 = 0;
    reader.read_to_string(&mut contents)?;
    let values = get_values(contents);
    values.into_iter().for_each(|(l1, l2)| {
        list1.push(l1);
        *list2.entry(l2).or_insert(0) += 1;
    });

    writeln!(stdout(), "2:list1 {:?}", list1)?;
    writeln!(stdout(), "2:list2 {:?}", list2)?;

    list1.into_iter().for_each(|i| {
        match list2.get(&i) {
            Some(res) => {
                total += i*res;
            }
            None => {}
        }
    });
    Ok(total)
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_part1() {
        let result = part1("data/test.txt").unwrap();
        assert_eq!(result, 11);
    }
    #[test]
    fn test_part2() {
        let result = part2("data/test.txt").unwrap();
        assert_eq!(result, 31);
    }
}
