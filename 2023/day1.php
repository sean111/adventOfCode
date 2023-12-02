<?php
require_once __DIR__ . '/vendor/autoload.php';

use AoC\Y2023\Day1;

try {
    $day1 = new Day1(__DIR__ . '/data/day1.txt');
    print "Sum: " . $day1->sum() . PHP_EOL;
} catch (\Exception $e) {
    die($e->getMessage());
}