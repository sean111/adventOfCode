<?php
ini_set('memory_limit', '-1');
require_once __DIR__ . '/vendor/autoload.php';

use AoC\Y2023\Day5;

$file = __DIR__ . '/data/day5.txt';

$day5 = new Day5($file);

$lowest = $day5->getLowestLocation();

print "Lowest: " . $lowest . PHP_EOL;
