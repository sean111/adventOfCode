<?php
use AoC\Y2023\Day2;

test('day2', function () {
    $file = __DIR__ . '/../data/day2.txt';

    $fp = fopen($file, 'r');

    $day2 = new Day2;

    while(!feof($fp)) {
        $input = fgets($fp);
        $day2->game($input);
    }

    $result = $day2->getSum();
    expect($result)->toBe(2286);
});

