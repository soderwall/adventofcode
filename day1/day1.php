<?php
$txt_file  = file_get_contents('day1.txt');
$array = str_split($txt_file);
$length = count($array);
$sum = 0;
$mid_point = $length/2;
for ($x = 0; $x <= $length-1; $x++) {
    if ($mid_point == $length){
        $mid_point = 0;
    }
    if ($array[$x] == $array[$mid_point]) {
        $sum += $array[$x];
    }
      $mid_point++;
}
print_r($sum);