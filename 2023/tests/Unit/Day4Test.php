<?php
use AoC\Y2023\Day4;
test('day 4', function () {
    $day4 = new Day4;
    $file = __DIR__ . "/../data/day4.txt";
    $fp = fopen($file, 'r');

    while(!feof($fp)) {
        if (!$line = trim(fgets($fp))) {
            throw new \Exception("error reading line");
        }
        $day4->newCard($line);
    }
    expect($day4->getPoints())->toBe(13);
});