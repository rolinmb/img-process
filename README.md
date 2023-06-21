main.py -> Uisng pillow library to create FX using bit-wise logical operations after converting to grayscale.

png.c -> Partially following a tutorial from [Tsoading Daily on youtube](https://www.youtube.com/watch?v=M9ZwuIv3xz8) to learn how to read .png file signatures/bytes.

gpt_png.c -> Reading .png file bytes proceedurally; program written using ChatGPT's help.

main.go -> Reading .png files with golang's native libaries; so far just takes an image and converts it to ascii representation output to your terminal;
may need ot zoom out a lot to see the full image generated. Will compartmentalize this feature as a separate function. Initially inspired by the tutorial [here.](https://golangdocs.com/golang-image-processing)

bg1i.png and cramer.png supplied as demo images; bg1i_terminal_06202023_440PMCST.png is the terminal output from processing bg1i.png

main.go cmdPrint() Example (6/20/2023 11:50PM CST): <br /> <br />
main.go Test 1 Input = ![here](demo_results/bg1i.png) <br />
main.go Test 1 Output = ![here](demo_results/bg1i_terminal_06202023_440PMCST.png) <br />
main.go Test 2 Input = ![here](demo_results/cramer.png) <br />
main.go Test 2 Output = ![here](demo_results/cramer_terminal_05202023_8PMCST.png) <br /> <br />
main.go trippyPng() Example (6/20/2023 11:50PM CST): <br /> <br />
![here](demo_results/trippy16_06202023.png) <br /> <br />
main.go noisyPng() Example(6/20/2023 11:50PM CST): <br /> <br />
![here](demo_results/noisy4_06202023.png) <br /> <br /> 

mb.go -> generates a .png of the [mandlebrot set](https://en.wikipedia.org/wiki/Mandelbrot_set) (in the location of your choosing) <br />
mb.go main() Example (6/21/2023 12:05AM CST): <br /><br />
![here](demo_results/mandelbrot.png) <br /><br />
julia.go -> generates a .png of the [julia set](https://en.wikipedia.org/wiki/Julia_set) (in the location of your choosing) <br />
mb.go main() Example (6/21/2023 12:30AM CST): <br /><br />
![here](demo_results/julia_set.png) <br /><br />