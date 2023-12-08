<?php
use AoC\Y2023\Day3;

test('day 3', function () {
    $day3 = new Day3;

    $file = __DIR__ . '/../data/day3.txt';

    $fp = fopen($file, 'r');
    while(!feof($fp)) {
        if (!$row = trim(fgets($fp))) {
            throw new Exception('error reading file');
        }
        $day3->addRow($row);
    }
    $sum = $day3->findSum();
    expect($sum)->toBe(4361);
});