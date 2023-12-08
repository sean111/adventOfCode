<?php

namespace AoC\Y2023;

use PHPUnit\Event\Runtime\PHP;

class Day3
{
    private array $schematic = [];
    private int $totalRows = 0;
    private int $totalColumns = 0;
    private array $currentBuffer = [];
    private bool $touchesSymbol = false;
    private int $sum = 0;

    public function addRow(string $row): void {
        $columns = str_split($row);
        $this->schematic[] = $columns;
        if ($this->totalColumns == 0) {
            $this->totalColumns = count($columns) -1;
        }
        $this->totalRows++;
    }

    public function findSum(): int {
        for ($r = 0; $r < $this->totalRows; $r++) {
            for ($c = 0; $c < $this->totalColumns; $c++) {
                $tmp = $this->schematic[$r][$c];
                if ($tmp == '.' || !is_numeric($tmp)) {
                    if ($this->touchesSymbol) {
                        $val = intval(implode('', $this->currentBuffer));
                        dump("Adding $val to $this->sum");
                        $this->sum += $val;
                    }
                    $this->currentBuffer = [];
                    $this->touchesSymbol = false;
                } else if (is_numeric($tmp)) {
                    if (!$this->touchesSymbol) {
//                        print "Checking $tmp ($r, $c)" . PHP_EOL;
                        if ($this->checkForSymbol($r, $c)) {
//                            print "\t$tmp touches a symbol" . PHP_EOL;
                            $this->touchesSymbol = true;
                        }
                    }
                    $this->currentBuffer[] = $tmp;
                }
                unset($tmp);
            }
            $this->currentBuffer = [];
            $this->touchesSymbol = false;
        }
        return $this->sum;
    }

    private function checkForSymbol(int $row, int $column): bool {

        // Check left
//        print "\tLeft: " . PHP_EOL;
        if ($this->isSymbol($row, $column-1)) {
            return true;
        }

        // Check top left
//        print "\tTop Left: " . PHP_EOL;
        if ($this->isSymbol($row-1, $column-1)) {
            return true;
        }

        // Check top
//        print "\tTop: " . PHP_EOL;
        if ($this->isSymbol($row-1, $column)) {
            return true;
        }

        // Check top right
//        print "\tTop Right: " . PHP_EOL;
        if ($this->isSymbol($row-1, $column+1)) {
            return true;
        }

        // Check right
//        print "\tRight: " . PHP_EOL;
        if ($this->isSymbol($row, $column+1)) {
            return true;
        }

        // Check bottom right
//        print "\tBottom Right: " . PHP_EOL;
        if($this->isSymbol($row+1, $column+1)) {
            return true;
        }

        // Check bottom
//        print "\tBottom: " . PHP_EOL;
        if ($this->isSymbol($row+1, $column)) {
            return true;
        }

        // Check bottom left
//        print "\tBottom Left: " . PHP_EOL;
        if ($this->isSymbol($row+1, $column-1)) {
            return true;
        }

        return false;
    }

    private function isSymbol($row, $column): bool {
        if ($row < 0 || $row >= $this->totalRows) {
            return false;
        }

        if ($column < 0 || $column > $this->totalColumns) {
            return false;
        }

        if (!is_numeric($this->schematic[$row][$column]) && $this->schematic[$row][$column] != '.') {
            print "\t($row, $column) => " . $this->schematic[$row][$column] . PHP_EOL;
            return true;
        }
        return false;
    }
}