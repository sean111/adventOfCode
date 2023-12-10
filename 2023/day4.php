<?php

require_once __DIR__ . '/vendor/autoload.php';

use AoC\Y2023\Day4;

$file = __DIR__ . '/data/day4.txt';

$day4 = new Day4;
$fp = fopen($file, 'r');

while(!feof($fp)) {
    if (!$line = trim(fgets($fp))) {
        throw new \Exception("error reading line");
    }
    $day4->newCard($line);
}

print "Total Points: " . $day4->getPoints() . PHP_EOL;