from PIL import Image, ImageEnhance, ImageFilter, ImageOps, ImageChops
import cv2
from skimage import filters
FSB = Image.FLOYDSTEINBERG
import time

def checImgkModes(img1,img2):
    if img1.mode != '1':
        img1 = img1.convert('1')
    if img2.mode != '1':
        img2 = img2.convert('1')
    return img1, img2

'''
Ditherting using Image.convert, mode "P" by default(can't change), and given parameter colors "c"
'''
def dithering(img,out,c=256):
    newImg = img.convert(mode="P",colors=c,dither=FSB,palette=Image.ADAPTIVE)
    newImg.save(out)
'''
Adjust brightness of image, 0 (Black Image) to ?(the larger the closer to White Image) 
'''
def adjBrightness(img,n,out):
    fx = ImageEnhance.Brightness(img)
    newImg = fx.enhance(n)
    newImg.save(out)
'''
Gaussian Blur for radius r
'''
def gBlur(img,r,out):
    newImg = img.filter(ImageFilter.GaussianBlur(r))
    newImg.save(out)
'''
Inverts color of image
'''  
def invertColor(img,out):
    img = Image.open(img)
    if img.mode == 'RGBA':
        r,g,b,a = img.split()
        rgb_image = Image.merge('RGB',(r,g,b))
        inverted_img = ImageOps.invert(rgb_image)
        r2,g2,b2 = inverted_img.split()
        final_transparent_image = Image.merge('RGBA',(r2,g2,b2,a))
        final_transparent_image.save(out)
    else:
        newImg = ImageOps.invert(img)
        newImg.save(out)
'''
Adds two images, dividing the result by scale and adding the offset. If omitted, scale defaults to 1.0, and offset to 0.0.
'''
def addComp(f1,f2,out,scale=1.0,offset=0):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.add(img1,img2,scale,offset)
    newImg.save(out)
'''
Add two images, without clipping the result.
'''
def addCompMod(f1,f2,out):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.add_modulo(img1,img2)
    newImg.save(out)
'''
Subtracts two images, dividing the result by scale and adding the offset. If omitted, scale defaults to 1.0, and offset to 0.0.
'''
def subComp(f1,f2,out,scale=1.0,offset=0):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.subtract(img1,img2,scale,offset)
    newImg.save(out)
'''
Subtract two images, without clipping the result.
'''
def subCompMod(f1,f2,out):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.subtract_modulo(img1,img2)
    newImg.save(out)
'''
Composes two images using AND, converts both image modes to '1' as needed
'''
def andComp(f1,f2,out):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.logical_and(img1,img2)
    newImg.save(out)
'''
Composes two images using OR, converts both image modes to '1' as needed        
'''
def orComp(f1,f2,out):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.logical_or(img1,img2)
    newImg.save(out)
''' 
Composes two images using XOR, converts both image modes '1' as needed
'''
def xorComp(f1,f2,out):
    img1, img2 = checkImgModes(Image.open(f1), Image.open(f2))
    newImg = ImageChops.logical_xor(img1,img2)
    newImg.save(out)
'''
Cropping f_in image to dimensions (d1,d2,d3,d4) and saving to output filename
'''
def cropping(f_in,out,d1,d2,d3,d4):
    img = Image.open(f_in)
    img.crop((d1,d2,d3,d4)).save(out)
'''
Edge detection with opencv and scikit-image
'''
def edgeDetect(f_in,out):
    img = cv2.imread(f_in, 0)
    edges = filters.roberts(img)
    cv2.imwrite(out, edges)

if __name__ == "__main__":
    start = time.time()
    #raw = "IMG_0766_crop.png"
    raw2 = "cramer_crop.png"
    raw = "IMG_0766_i.png"
    #raw2 = "cramer_i.png"
    out_name1 = "cramer_0766I_and.png"
    out_name2 = "cramer_0766I_or.png"
    out_name3 = "cramer_0766I_xor.png"
    #cropping(raw,out_name,0,0,720,720)
    #invertColor(raw,out_name)
    andComp(raw,raw2,out_name1)
    orComp(raw,raw2,out_name2)
    xorComp(raw,raw2,out_name3)
    #edgeDetect(raw,out_name)
    print('Execution Time: %s seconds'%round(time.time()-start,2))