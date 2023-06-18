main.py -> Uisng pillow library to create FX using bit-wise logical operations after converting to grayscale.

png.c -> Partially following a tutorial from (Tsoading Daily on youtube)[https://www.youtube.com/watch?v=M9ZwuIv3xz8] to learn how to read .png file signatures/bytes.

gpt_png.c -> Reading .png file bytes proceedurally; program written using ChatGPT's help. Uses "miniz.c" and "miniz.h"from the repo (miniz)[https://github.com/richgel999/miniz]
since I am developing on Windows and zlib.h is not linking/compiling properly at this moment. 

must compile & run using the powershell command(s): <b>gcc gpt_png.c miniz.c -o gpt_png</b>