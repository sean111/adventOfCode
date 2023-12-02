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
            throw new \Exception("error getting first int (" . implode('',$line) . ")");
        }

        if (!$last= $this->getLastInt($line)) {
            throw new \Exception("error getting last int (" . implode('',$line) . ")");
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

            // Check for strings
            if ($x+3 < $cnt) {
                $tmp = implode('', array_slice($line, $x, 3));
                switch($tmp) {
                    case "one":
                        return 1;
                    case "two":
                        return 2;
                    case "thr":
                        if (implode('', array_slice($line, $x, 5)) == "three") {
                            return 3;
                        }
                        break;
                    case "fou":
                        if (implode('', array_slice($line, $x, 4)) == "four") {
                            return 4;
                        }
                        break;
                    case "fiv":
                        if (implode('', array_slice($line, $x, 4)) == "five") {
                            return 5;
                        }
                        break;
                    case "six":
                        return 6;
                    case "sev":
                        if (implode('', array_slice($line, $x, 5)) == "seven") {
                            return 7;
                        }
                        break;
                    case "eig":
                        if (implode('', array_slice($line, $x, 5)) == "eight") {
                            return 8;
                        }
                        break;
                    case "nin":
                        if (implode('', array_slice($line, $x, 4)) == "nine") {
                            return 9;
                        }
                        break;
                }
            }
        }
        return false;
    }
    
    private function getLastInt(array $line): string|false {
        for ($x = (count($line) - 1); $x >= 0; $x--) {
            if (is_numeric($line[$x])) {
                return $line[$x];
            }

            //Check for strings
            $tmp = implode('', array_slice($line, $x-3, 3));
            switch($tmp) {
                case "one":
                    return 1;
                case "two":
                    return 2;
                case "thr":
                    if (implode('', array_slice($line, $x-3, 5)) == "three") {
                        return 3;
                    }
                    break;
                case "fou":
                    if (implode('', array_slice($line, $x-3, 4)) == "four") {
                        return 4;
                    }
                    break;
                case "fiv":
                    if (implode('', array_slice($line, $x-3, 4)) == "five") {
                        return 5;
                    }
                    break;
                case "six":
                    return 6;
                case "sev":
                    if (implode('', array_slice($line, $x-3, 5)) == "seven") {
                        return 7;
                    }
                    break;
                case "eig":
                    if (implode('', array_slice($line, $x-3, 5)) == "eight") {
                        return 8;
                    }
                    break;
                case "nin":
                    if (implode('', array_slice($line, $x-3, 4)) == "nine") {
                        return 9;
                    }
                    break;
            }
        }
        return false;
    }
}