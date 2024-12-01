<?php
use AoC\Y2023\Day5;
test('day 5', function() {
    $file = __DIR__ . '/../data/day5.txt';
    $day5 = new Day5($file);
    expect($day5->getSoilFromSeed(79))->toBe(81)
        ->and($day5->getFertilizerFromSoil(81))->toBe(81)
        ->and($day5->getLowestLocation())->toBe(35);
});