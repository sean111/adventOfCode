<?php

namespace AoC\Y2023;

class Day4
{
    private int $totalPoints = 0;
    public function getPoints(): int
    {
        return $this->totalPoints;
    }

    public function newCard(string $input): void
    {
        $points = 0;
        $tmp = strpos($input, ":");
        $tmp = mb_substr($input, $tmp+2);
        $tmp = explode("|", $tmp);
        $winningNumbers = array_filter(explode(" ", $tmp[0]));
        $userNumbers = array_filter(explode(" ", $tmp[1]));
        unset($tmp);
        foreach ($winningNumbers as $number) {
            if (in_array($number, $userNumbers)) {
                if ($points == 0) {
                    $points = 1;
                } else {
                    $points *= 2;
                }
            }
        }
        $this->totalPoints += $points;
    }
}