main.py -> Uisng pillow library to create FX using bit-wise logical operations after converting to grayscale.

png.c -> Partially following a tutorial from (Tsoading Daily on youtube)[https://www.youtube.com/watch?v=M9ZwuIv3xz8] to learn how to read .png file signatures/bytes.

gpt_png.c -> Reading .png file bytes proceedurally; program written using ChatGPT's help. Couldn't complete the decompression of .PNG IDAT bytes due to zlib.h not
being properly inlcuded/compiled. I think this is primarily due to developing in C on windows in MinGW64.