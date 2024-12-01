<?php

namespace AoC\Y2023;

class Day5
{
    private $seeds = [];
    private $seedToSoil = [];
    private $soilToFertilizer = [];
    private $fertilizerToWater = [];
    private $waterToLight = [];
    private $lightToTemperature = [];
    private $tempteratureToHumidity = [];
    private $humidityToLocation = [];

    public function __construct(string $file) {
        $fp = fopen($file, "r");
        print "Parsing file...." . PHP_EOL;
        $parser = null;
        while (($data = fgets($fp)) !== false) {
            $marker = substr($data, 0, 5);
            switch ($marker) {
                case "seeds":
                    $this->parseSeeds($data);
                    break;
                case "seed-":
                    $parser = "parseSeedsToSoil";
                    break;
                case "soil-":
                    $parser = "parseSoilToFertilizer";
                    break;
                case "ferti":
                    $parser = "parseFertilizerToWater";
                    break;
                case "water":
                    $parser = "parseWaterToLight";
                    break;
                case "light":
                    $parser = "parseLightToTemperature";
                    break;
                case "tempu":
                    $parser = "parseTemperatureToHumidity";
                    break;
                case "humid":
                    $parser = "parseHumidityToLocation";
                    break;
                default:
                    if ($parser !== null && strlen($data) > 1) {
                        if (method_exists($this, $parser)) {
                            $this->{$parser}($data);
                        }
                    }
                    break;
            }
        }
        fclose($fp);
    }

    private function parseSeeds($line) {
        print "Parsing seeds..." . PHP_EOL;
        $temp = explode(" ", $line);
        for ($x = 1, $len = count($temp); $x < $len; $x++) {
            $this->seeds[] = (int)$temp[$x];
        }
    }

    private function parseSeedsToSoil($line) {
        print "Parsing seeds to soil..." . PHP_EOL;
        $temp = explode(" ", $line);
        $soil = (int)$temp[0];
        $seed = (int)$temp[1];
        $range = (int)$temp[2];
        $this->seedToSoil[] = ['seed' => $seed, 'soil' => $soil, 'range' => $range];
    }

    private function parseSoilToFertilizer($line) {
        print "Parsing soil to fertilizer..." . PHP_EOL;
        $temp = explode(" ", $line);
        $fertilizer = (int)$temp[0];
        $soil = (int)$temp[1];
        $range = (int)$temp[2];
        $this->soilToFertilizer[] = ['soil' => $soil, 'fertilizer' => $fertilizer, 'range' => $range];
    }

    private function parseFertilizerToWater($line) {
        print "Parsing fertilizer to water..." . PHP_EOL;
        $temp = explode(" ", $line);
        $water = (int)$temp[0];
        $fertilizer = (int)$temp[1];
        $range = (int)$temp[2];
        $this->fertilizerToWater[] = ['fertilizer' => $fertilizer, 'water' => $water, 'range' => $range];
    }

    private function parseWaterToLight($line) {
        print "Parsing water to light..." . PHP_EOL;
        $temp = explode(" ", $line);
        $light = (int)$temp[0];
        $water = (int)$temp[1];
        $range = (int)$temp[2];
        $this->waterToLight[] = ['water' => $water, 'light' => $light, 'range' => $range];
    }

    private function parseLightToTemperature($line) {
        print "Parsing light to temperature..." . PHP_EOL;
        $temp = explode(" ", $line);
        $temperature = (int)$temp[0];
        $light = (int)$temp[1];
        $range = (int)$temp[2];
        $this->lightToTemperature[] = ['light' => $light, 'temperature' => $temperature, 'range' => $range];
    }

    private function parseTemperatureToHumidity($line) {
        print "Parsing temperature to humidity..." . PHP_EOL;
        $temp = explode(" ", $line);
        $humidity = (int)$temp[0];
        $temperature = (int)$temp[1];
        $range = (int)$temp[2];
        $this->tempteratureToHumidity[] = ['temperature' => $temperature, 'humur' => $humidity, 'range' => $range];
    }

    private function parseHumidityToLocation($line) {
        print "Parsing humidity to location..." . PHP_EOL;
        $temp = explode(" ", $line);
        $location = (int)$temp[0];
        $humidity = (int)$temp[1];
        $range = (int)$temp[2];
        $this->humidityToLocation[] = ['humidity' => $humidity, 'location' => $location, 'range' => $range];
    }

    public function getSoilFromSeed($seed) {
        foreach ($this->seedToSoil as $item) {
            if ($seed >= $item['seed'] && $seed <= $item['seed']+$item['range']) {
                $diff = $seed - $item['seed'];
                return $item['soil']+$diff;
            }
        }
        return $seed;
    }

    public function getFertilizerFromSoil($soil) {
        foreach($this->soilToFertilizer as $item) {
            if ($soil >= $item['soil'] && $soil <= $item['soil']+$item['range']) {
                $diff = $soil - $item['soil'];
                return $item['fertilizer']+$diff;
            }
        }
        return $soil;
    }

    public function getWaterFromFertilizer($fertilizer) {
        foreach($this->fertilizerToWater as $item) {
            if ($fertilizer >= $item['fertilizer'] && $fertilizer <= $item['fertilizer']+$item['range']) {
                $diff = $fertilizer - $item['fertilizer'];
                return $item['water']+$diff;
            }
        }
        return $fertilizer;
    }

    public function getLightFromWater($water) {
        foreach($this->waterToLight as $item) {
            if ($water >= $item['water'] && $water <= $item['water']+$item['range']) {
                $diff = $water - $item['water'];
                return $item['light']+$diff;
            }
        }
        return $water;
    }

    public function getTemperatureFromLight($light) {
        foreach($this->lightToTemperature as $item) {
            if ($light >= $item['light'] && $light <= $item['light']+$item['range']) {
                $diff = $light - $item['light'];
                return $item['temperature']+$diff;
            }
        }
        return $light;
    }

    public function getHumidityFromTemperature($temperature) {
        foreach($this->tempteratureToHumidity as $item) {
            if ($temperature >= $item['temperature'] && $temperature <= $item['humidity']+$item['range']) {
                $diff = $temperature - $item['temperature'];
                return $item['humidity']+$diff;
            }
        }
        return $temperature;
    }

    public function getLocationFromHumidity($humidity) {
        foreach($this->humidityToLocation as $item) {
            if ($humidity >= $item['humidity'] && $humidity <= $item['humidity']+$item['range']) {
                $diff = $humidity - $item['humidity'];
                return $item['location']+$diff;
            }
        }
        return $humidity;
    }

    public function getLowestLocation() {
        $lowest = null;
        foreach($this->seeds as $seed) {
            print "checking $seed" . PHP_EOL;
            $soil = $this->getSoilFromSeed($seed);
            $fertilizer = $this->getFertilizerFromSoil($soil);
            $water = $this->getWaterFromFertilizer($fertilizer);
            $light = $this->getLightFromWater($water);
            $temperature = $this->getTemperatureFromLight($light);
            $humidity = $this->getHumidityFromTemperature($temperature);
            $location = $this->getLocationFromHumidity($humidity);
            print "Seed: $seed -> Soil: $soil -> Fertilizer: $fertilizer -> Water: $water -> Light: $light -> Temperature: $temperature -> Humidity: $humidity -> Location: $location"  . PHP_EOL;
            if ($location < $lowest || $lowest === null) {
                $lowest = $location;
            }
        }
        return $lowest;
    }
}
