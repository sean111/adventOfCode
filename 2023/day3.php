<?php

require_once __DIR__ . '/vendor/autoload.php';

use AoC\Y2023\Day3;

$day3 = new Day3;

$file = __DIR__ . '/data/day3.txt';

$fp = fopen($file, 'r');
while(!feof($fp)) {
    if (!$row = trim(fgets($fp))) {
        throw new Exception('error reading file');
    }
    $day3->addRow($row);
}

print "Sum: " . $day3->findSum() . PHP_EOL;