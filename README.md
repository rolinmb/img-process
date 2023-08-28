main.py -> Uisng pillow library to create FX using bit-wise logical operations after converting to grayscale.

png.c -> Partially following a tutorial from [Tsoading Daily on youtube](https://www.youtube.com/watch?v=M9ZwuIv3xz8) to learn how to read .png file signatures/bytes.

gpt_png.c -> Reading .png file bytes proceedurally; program written using ChatGPT's help.

[main.go]
may need ot zoom out a lot to see the full image generated. Will compartmentalize this feature as a separate function. Initially inspired by the tutorial [here.](https://golangdocs.com/golang-image-processing)

bg1i.png and cramer.png supplied as demo images; bg1i_terminal_06202023_440PMCST.png is the terminal output from processing bg1i.png

main.go cmdPrint() Example (6/20/2023 11:50PM CST): <br /><br />
main.go Test 1 Input = ![here](demo_results/bg1i.png) <br />
main.go Test 1 Output = ![here](demo_results/bg1i_terminal_06202023_440PMCST.png) <br />
main.go Test 2 Input = ![here](demo_results/cramer.png) <br />
main.go Test 2 Output = ![here](demo_results/cramer_terminal_05202023_8PMCST.png) <br /><br />
main.go trippyPng() Example (6/20/2023): <br /><br />
![here](demo_results/trippy16_06202023.png) <br /><br />
main.go noisyPng() Example(6/20/2023): <br /><br />
![here](demo_results/noisy4_06202023.png) <br /><br /> 

mb.go -> generates a .png of the [mandlebrot set](https://en.wikipedia.org/wiki/Mandelbrot_set)<br />
Example (6/21/2023): <br /><br />
![here](demo_results/mandelbrot.png) <br /><br />
julia.go -> generates a .png of the [julia set](https://en.wikipedia.org/wiki/Julia_set)<br />
Example (6/21/2023): <br /><br />
![here](demo_results/julia_set.png) <br /><br /> 
newton.go -> generates a .png of the [newton fractal](https://www.unf.edu/~ddreibel/teaching/newton/index.html)<br /> 
Example (6/21/2023): <br /><br />
![here](demo_results/newton_fractal.png) <br /><br />
bs.go -> generates a .png of the [burning ship fractal](https://en.wikipedia.org/wiki/Burning_Ship_fractal). This one probably needs the most rendering to give an interesting picture. <br />
Example (6/21/2023): <br /><br />
![here](demo_results/burning_ship.png) <br /><br />
mandelbulb.go -> generates a .png of a [mandelbulb fractal](https://en.wikipedia.org/wiki/Mandelbulb)<br /> 
Example (8/27/2023): <br /><br />
![here](demo_results/mandelbulb.png) <br /><br />
multibrot.go -> generates a .png of a [multibrot fractal](https://en.wikipedia.org/wiki/Multibrot_set)<br />
Example (8/27/2023): <br /><br />
![here](demo_results/multibrot.png) <br /><br />
mandelbox.go -> generates a .png of a [mandelbox fractal](https://en.wikipedia.org/wiki/Mandelbox)<br />
Example (8/27/2023): <br /><br />
TODO: doesn't generate at the moment
lyapunov.go -> generates a .png of a [lyapunov fractal](https://en.wikipedia.org/wiki/Lyapunov_fractal)<br />
Example (8/27/2023): <br /><br />
TODO: doesn't generate at the moment
gosper.go -> generates a .png of a [gosper curve](https://en.wikipedia.org/wiki/Gosper_curve)<br />
Example (8/27/2023): <br /><br />
![here](demo_results/gosper.png) <br /><br />
tchebi.go -> generates a .png of [tchebichef fractal](https://arxiv.org/abs/2102.10640)<br />
TODO: tchebi.go not currently functioning; takes too long to finish execution (infinite loop perhaps? or i'm not patient enough?)