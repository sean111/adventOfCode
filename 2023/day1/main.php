<?php
require_once __DIR__ . '/vendor/autoload.php';

try {
    $day1 = new AoC\Day1(__DIR__ . '/data.txt');
    print "Sum: " . $day1->sum() . PHP_EOL;
} catch (\Exception $e) {
    die($e->getMessage());
}