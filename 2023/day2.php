<?php
require_once __DIR__ . '/vendor/autoload.php';

use AoC\Y2023\Day2;

$day2 = new Day2(12, 13, 14);

$file = __DIR__ . '/data/day2.txt';

if (!file_exists($file)) {
    die('file does not exist');
}

$fp = fopen($file, 'r');

while(!feof($fp)) {
    if (!$line = fgets($fp)) {
        die("error parsing game");
    }
    $day2->game($line);
}
fclose($fp);

print "Sum: " .$day2->getSum() . PHP_EOL;