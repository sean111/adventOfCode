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
    }

    public function findSum(): int {
        $this->totalRows = count($this->schematic);
        $this->totalColumns = count($this->schematic[0]);
        for ($r = 0; $r < $this->totalRows; $r++) {
            for ($c = 0; $c < $this->totalColumns; $c++) {
                $tmp = $this->schematic[$r][$c];
                if ($tmp == '.' || !is_numeric($tmp)) {
                    if ($this->touchesSymbol) {
                        $val = intval(implode('', $this->currentBuffer));
                        $this->sum += $val;
                    }
                    $this->currentBuffer = [];
                    $this->touchesSymbol = false;

                } else if (is_numeric($tmp)) {
                    if (!$this->touchesSymbol) {
                        if ($this->checkForSymbol($r, $c)) {
                            $this->touchesSymbol = true;
                        }
                    }
                    $this->currentBuffer[] = $tmp;
                }
                unset($tmp);
            }
            if ($this->touchesSymbol) {
                $val = intval(implode('', $this->currentBuffer));
                $this->sum += $val;
            }
            $this->currentBuffer = [];
            $this->touchesSymbol = false;
        }
        return $this->sum;
    }

    private function checkForSymbol(int $row, int $column): bool {

        // Check left
        if ($this->isSymbol($row, $column-1)) {
            return true;
        }

        // Check top left
        if ($this->isSymbol($row-1, $column-1)) {
            return true;
        }

        // Check top
        if ($this->isSymbol($row-1, $column)) {
            return true;
        }

        // Check top right
        if ($this->isSymbol($row-1, $column+1)) {
            return true;
        }

        // Check right
        if ($this->isSymbol($row, $column+1)) {
            return true;
        }

        // Check bottom right
        if($this->isSymbol($row+1, $column+1)) {
            return true;
        }

        // Check bottom
        if ($this->isSymbol($row+1, $column)) {
            return true;
        }

        // Check bottom left
        if ($this->isSymbol($row+1, $column-1)) {
            return true;
        }

        return false;
    }

    private function isSymbol($row, $column): bool {
        if ($row < 0 || $row >= $this->totalRows) {
            return false;
        }

        if ($column < 0 || $column >= $this->totalColumns) {
            return false;
        }

        if (!is_numeric($this->schematic[$row][$column]) && $this->schematic[$row][$column] != '.') {
            return true;
        }
        return false;
    }
}