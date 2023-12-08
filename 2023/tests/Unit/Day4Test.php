<?php
use AoC\Y2023\Day4;
test('day 4', function () {
    $day4 = new Day4;
    $sum = $day4->getPoints();
    expect($sum)->toBe(13);
});