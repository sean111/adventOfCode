<?php
namespace AoC\Y2023;

class Day2 {
    private array $maxCubes = ['red' => 0, 'blue' => 0, 'green' => 0];
    private int $currentGameId = 0;
    private int $sumOfIds = 0;

    public function __construct(int $red, int $green, int $blue) {
        $this->maxCubes = [
            'red' => $red,
            'blue' => $blue,
            'green' => $green
        ];
    }

    public function game(string $gameInput): void {
        $this->currentGameId++;
        $tmp = explode(':', $gameInput);
        $sets = explode(';', $tmp[1]);
        unset($tmp);
        for ($x = 0, $cnt = count($sets); $x < $cnt; $x++) {
            $data = explode(',', $sets[$x]);

            // Check inputs
            for ($i = 0, $iCount = count($data); $i < $iCount; $i++) {
                $input = trim($data[$i]);
                $tmp = intval($input);

                switch(mb_substr($input, -1)) {
                    case 'd':
                        if ($tmp > $this->maxCubes['red']) {
                            return;
                        }
                        break;
                    case 'e':
                        if ($tmp > $this->maxCubes['blue']) {
                            return;
                        }
                        break;
                    case 'n':
                        if ($tmp > $this->maxCubes['green']) {
                            return;
                        }
                        break;
                }
            }
        }
        $oldSum = $this->sumOfIds;
        $this->sumOfIds += $this->currentGameId;
    }

    public function getSum(): int {
        return $this->sumOfIds;
    }
}