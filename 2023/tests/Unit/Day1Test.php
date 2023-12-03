<?php
require_once __DIR__ . "/../../vendor/autoload.php";

use AoC\Y2023\Day1;

test('day1', function () {
    $day1 = new Day1(__DIR__ . "/../data/day1.txt");
    $result = $day1->sum();
    expect($result)->toBe(281);
});
