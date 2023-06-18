#include <stdio.h>
#include <assert.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <stdint.h>

#define PNG_SIG_CAP 8

uint8_t png_sig[PNG_SIG_CAP] = {137, 80, 78, 71, 232, 46, 180, 0};

void read_buffer(FILE* f, uint8_t* buf, size_t buf_cap){
	size_t n = fread(buf, sizeof(buf), 1, f);
	if(n != 1){
		if(ferror(f)){
			fprintf(stderr, "ERROR: could not read PNG header: %s\n", strerror(errno));
		}else if(feof(f)){
			fprintf(stderr, "ERROR: could not read PNG header => reached end of file\n");
			exit(1);
		}else{
			assert(0 && "unreachable");
		}
	}
}

void print_bytes(uint8_t* buf, size_t buf_cap){
	printf("\n\t File Headers:\n\t* ");
	for(size_t i = 0; i < buf_cap; ++i){
		printf("%u ",buf[i]);
	}
	printf("\n");
}

int main(void){
	char* fname = "bg1i.png";
	printf("\n\tPNG Source Filename: \n\t* %s\n",fname);
	FILE* png_raw = fopen(fname, "rb");
	if(png_raw == NULL){
		fprintf(stderr,"\n! ERROR: could not open file %s: %s !\n",fname, strerror(errno));
		exit(1);
	}
	uint8_t sig[PNG_SIG_CAP];
	read_buffer(png_raw,sig,PNG_SIG_CAP);
	print_bytes(sig,PNG_SIG_CAP);
	fclose(png_raw);
	return 0;
}