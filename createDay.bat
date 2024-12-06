@echo off
set /p day="Input day as number: "

mkdir day%day%

cd day%day%

copy NUL example.txt
copy NUL input.txt

echo package day%day% >"part1.go"
echo.>> "part1.go"
echo func Day%day%Part1() {>> "part1.go"
echo.>> "part1.go"
echo }>> "part1.go"

echo package day%day% >"part2.go"
echo.>> "part2.go"
echo func Day%day%Part2() {>> "part2.go"
echo.>> "part2.go"
echo }>> "part2.go"