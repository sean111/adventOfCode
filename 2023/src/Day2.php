<?php
namespace AoC\Y2023;

class Day2 {
    private array $maxCubes;
    private int $currentGameId = 0;
    private int $sum = 0;

    public function game(string $gameInput): void {
        $this->maxCubes = ['red' => 0, 'blue' => 0, 'green' => 0];
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
                            $this->maxCubes['red'] = $tmp;
                        }
                        break;
                    case 'e':
                        if ($tmp > $this->maxCubes['blue']) {
                            $this->maxCubes['blue'] = $tmp;
                        }
                        break;
                    case 'n':
                        if ($tmp > $this->maxCubes['green']) {
                            $this->maxCubes['green'] = $tmp;
                        }
                        break;
                }
            }
        }
        $this->sum += $this->getPower();
    }

    private function getPower():int {
        return array_product($this->maxCubes);
    }
    public function getSum(): int {
        return $this->sum;
    }
}