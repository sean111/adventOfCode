<?php

namespace AoC;

class Day1 {
    function __construct(private string $file) {
        if (!file_exists($this->file)) {
            throw new \Exception("file not found");
        }
    }
    
    public function sum(): int {
        $total = 0;
        $fp = fopen($this->file, 'r');
        while (!feof($fp)) {
            $line = fgets($fp);
            if (!$line) {
                throw new \Exception("error reading line from $this->file");
            }
            $total += $this->parseLine(str_split($line));
        }
        fclose($fp);
        return $total;
    }
    
    private function parseLine(array $line): int {
//        $tmp = $this->getFirstInt($line) . $this->getLastInt($line);
        if (!$first = $this->getFirstInt($line)) {
            throw new \Exception("error getting first int");
        }
        if (!$last= $this->getLastInt($line)) {
            throw new \Exception("error getting last int");
        }
        $tmp = $first.$last;
//        var_dump(['line'=> $line,'first' =>$first, 'last' =>$last, 'tmp' => $tmp, 'int' => (int)$tmp]);
        return (int)$tmp;
    }
    
    private function getFirstInt(array $line): string|false {
        for ($x = 0, $cnt = count($line); $x < $cnt; $x++) {
            if (is_numeric($line[$x])) {
                return $line[$x];
            }
        }
        return false;
    }
    
    private function getLastInt(array $line): string|false {
        for ($x = (count($line) - 1); $x >= 0; $x--) {
            if (is_numeric($line[$x])) {
                return $line[$x];
            }
        }
        return false;
    }
}