<?php

require_once __DIR__ . "/../../vendor/autoload.php";

test('sum', function () {
    $day1 = new AoC\Day1(__DIR__ . "/../../test.txt");
    $result = $day1->sum();
    expect($result)->toBe(142);
});
