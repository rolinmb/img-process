#include <stdio.h>
#include <stdlib.h>
#include <string.h>
//#include <zlib.h>

#define SIG_SIZE 8
#define IHDR_SIZE 25

int main(void){
	char* fname = "bg1i.png";
	FILE *fp = fopen(fname,"rb");
	if(!fp){
		printf("Failed to open the PNG file.\n`");
		return 1;
	}
	/*// determine .png file total size
	fseek(fp,0,SEEK_END);
	long FILE_SIZE = ftell(fp);
	rewind(fp);
	// allocating memory to store read bytes
	unsigned char* bufr = (unsigned char*)malloc(FILE_SIZE);
	if(!bufr){
		printf("\n\nFailed to allocate memory to read .PNG file.\n\n");
		fclose(fp);
		return 1;
	}
	// read bytes into bufr
	size_t bytes_read = fread(bufr,1,FILE_SIZE,fp);
	if(bytes_read != FILE_SIZE){
		printf("\n\nFailed to read the entire PNG file into bufr.\n\n");
		free(bufr);
		fclose(fp);
		return 1;
	}
	printf("\n\nDisplaying entire .PNG file bytes:\n\n");
	for(long i = 0; i < FILE_SIZE; ++i){
		printf("%02X",bufr[i]);
		if((i+1)%16 == 0){
			printf("\n");
		}
	}
	free(bufr);
	*/
	unsigned char png_sig[SIG_SIZE];
	fread(png_sig,1,SIG_SIZE,fp);
	// sig = 89 50 4E 47 0D 0A 1A 0A (hexadecimal)
	// sig = 137 80 78 71 232 46 180 0 (integer)
	printf("\n\nDisplaying .PNG signature bytes:\n\n");
	for(size_t i = 0; i < SIG_SIZE; ++i){
		printf("%02X",png_sig[i]);
		if((i+1)%16 == 0){
			printf("\n");
		}
	}
	unsigned char png_ihdr[IHDR_SIZE];
	fread(png_ihdr,1,IHDR_SIZE,fp);
	printf("\n\nDisplaying .PNG IHDR bytes:\n\n");
	for(size_t i = 0; i < IHDR_SIZE; ++i){
		printf("%02X",png_ihdr[i]);
		if((i+1)%16 == 0){
			printf("\n");
		}
	}
	unsigned int png_width = png_ihdr[3] << 24 | png_ihdr[4] << 16 | png_ihdr[5] << 8 | png_ihdr[6];
	unsigned int png_height = png_ihdr[7] << 24 | png_ihdr[8] << 16 | png_ihdr[9] << 8 | png_ihdr[10];
	unsigned char png_bit_depth = png_ihdr[12];
	unsigned char png_color_type = png_ihdr[13];
	// Skip over additional chunks until IDAT chunk is encountered
	unsigned char chunk_type[5];
	unsigned int chunk_len;
	while(1){
		fread(&chunk_len,4,1,fp);
		fread(chunk_type,1,4,fp);
		if(memcmp(chunk_type,"IDAT",4) == 0){
			break;
		}
		fseek(fp,chunk_len+4,SEEK_CUR);
	}
	// Extract image data from IDAT chunk
	unsigned char* compressed_data = (unsigned char*)malloc(chunk_len);
	if(!compressed_data){
		printf("Failed to allocate memory for compressed_data.\n`");
		free(compressed_data);
		fclose(fp);
		return 1;
	}
	size_t idat_bytes = fread(compressed_data,1,chunk_len,fp);
	if(idat_bytes != chunk_len){
		printf("Failed to read the PNG into compressed_bytes.\n");
		free(compressed_data);
		fclose(fp);
		return 1;
	}
	printf("\n\nDisplaying compressed .PNG IDAT bytes:\n\n");
	for(size_t i = 0; i < idat_bytes; ++i){
		printf("%02X",compressed_data[i]);
		if((i+1)%16 == 0){
			printf("\n");
		}
	}
	// can't decompress without implementation of DEFLATE/INFLATE algorithm
	/*
	// Instantiate z_stream to decompress the data (doesn't work on windows...)
	z_stream stream;
	stream.zalloc = Z_NULL;
	stream.zfree = Z_NULL;
	stream.opaque = Z_NULL;
	stream.avail_in = chunk_len;
	stream.nex_in = compressed_data;
	*/
	free(compressed_data);
	fclose(fp);
	return 0;
}