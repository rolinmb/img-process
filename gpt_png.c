#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <zlib.h>

void view_all_bytes(long n_bytes,FILE* fp){
  printf("\nTotal .png file size (bytes): %ld\n",n_bytes);
  // allocating memory to store read bytes
  unsigned char* bufr = (unsigned char*)malloc(n_bytes);
  if(!bufr){
    printf("\nFailed to allocate memory to read .PNG file.\n\n");
		fclose(fp);
    exit(1);
	}
	// read bytes into bufr
	size_t bytes_read = fread(bufr,1,n_bytes,fp);
	if(bytes_read != n_bytes){
		printf("\nFailed to read the entire PNG file into bufr.\n\n");
		free(bufr);
		fclose(fp);
    exit(1);
	}
	printf("\nDisplaying entire .PNG file bytes:\n\n");
	for(long i = 0; i < n_bytes; ++i){
		printf("%02X",bufr[i]);
		if((i+1)%16 == 0){
			printf("\n");
		}
	}
	free(bufr);
  rewind(fp);
  printf("\n");
}
// Must compile with -lz tag at the end of args
int main(void){
  char* fname = "bg1i.png";
	FILE *fp = fopen(fname,"rb");
	if(!fp){
		printf("Failed to open the PNG file.\n`");
		return 1;
	}
	fseek(fp,0,SEEK_END);
	long FILE_SIZE = ftell(fp);
	rewind(fp);
	view_all_bytes(FILE_SIZE, fp);
  unsigned char png_sig[8]; // SIG chunk is first 8 bytes
  fread(png_sig,1,8,fp);
  printf("\nDisplaying .PNG SIG bytes:\n\n");
  for(size_t i = 0; i < 8; ++i){
    printf("%02X",png_sig[i]);
    if((i+1)%16 == 0){
      printf("\n");
    }
  }
  unsigned char png_ihdr[25]; // IHDR chunk is next 25 bytes
  fread(png_ihdr,1,25,fp);
  printf("\n\n\nDisplaying .PNG IHDR bytes:\n\n");
  for(size_t i = 0; i < 25; ++i){
    printf("%02X",png_ihdr[i]);
    if((i+1)%16 == 0){
      printf("\n");
    }
  }
  unsigned int png_width = png_ihdr[3] << 24 | png_ihdr[4] << 16 | png_ihdr[5] << 8 | png_ihdr[6];
  unsigned int png_height = png_ihdr[7] << 24 | png_ihdr[8] << 16 | png_ihdr[9] << 8 | png_ihdr[10];
  unsigned char png_bit_depth = png_ihdr[12];
  unsigned char png_color_type = png_ihdr[13];
  unsigned char chunk_type[5];
  unsigned int idat_len;
  while(1){ // Determine how many bytes IDAT chunk is
    fread(&idat_len,4,1,fp);
    fread(chunk_type,1,4,fp);
    if(memcmp(chunk_type,"IDAT",4) == 0){
      break;
    }
    fseek(fp,idat_len+4,SEEK_CUR);
  }
  unsigned char* compressed = (unsigned char*)malloc(idat_len);
  if(!compressed){
    printf("\nFailed to allocate memory for reading IDAT chunk.\n\n");
    fclose(fp);
    return 1;
  }
  fread(compressed,1,idat_len,fp);
  printf("\n\nDisplaying .PNG compressed IDAT bytes:\n\n");
  for(size_t i = 0; i< idat_len; ++i){
    printf("%02X",compressed[i]);
    if((i+1)%16 == 0){
      printf("\n");
    }
  }
  printf("\nPreparing INFLATE to decompress IDAT chunk using zlib.n z_stream:\n\n");
  // NEW
  size_t dcomp_buf_size = png_width*png_height*4;
  unsigned char* dcompressed = (unsigned char*)malloc(dcomp_buf_size);
  if(!dcompressed){
    printf("\nFailed to allocate memory for decompressing IDAT chunk.\n\n");
    free(compressed);
    fclose(fp);
    return 1;
  }
  z_stream stream;
  stream.zalloc = Z_NULL;
  stream.zfree = Z_NULL;
  stream.opaque = Z_NULL;
  stream.avail_in = (uInt)idat_len;
  stream.next_in = compressed;
  stream.avail_out = (uInt)dcomp_buf_size;
  stream.next_out = dcompressed;
  // perform INFLATE
  if(inflateInit(&stream) != Z_OK){
    printf("\nFailed to initialize zlib z_stream.\n\n");
    free(compressed);
    free(dcompressed);
    fclose(fp);
    return 1;
  }
  if(inflate(&stream,Z_FINISH) != Z_STREAM_END){
    printf("\nFailed to INFLATE/decompress .PNG IDAT chunk bytes. Error code %s\n\n",stream.msg);
    inflateEnd(&stream);
    free(compressed);
    free(dcompressed);
    fclose(fp);
    return 1;
  }
  inflateEnd(&stream);
  for(size_t i = 0; i < dcomp_buf_size; ++i){
    printf("%02X",dcompressed[i]);
    if((i+1)%16 == 0){
      printf("\n");
    }
  }
  printf("\n\n");
  // cleaning up and exiting
  free(compressed);
  free(dcompressed);
  fclose(fp);
  return 0;
}
