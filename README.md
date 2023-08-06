# compare2images
A small Go library which compares 2 images pixel by pixel.  
  
You can call Compare2Images from your Go code supplying it with 2 *image.Image and it returns a struct:  

type compResult struct {  
	Different uint64  
	Percent   float64  
	RedGreen  *image.NRGBA  
	Faded     *image.NRGBA  
}  
  
Where Different is the number of pixels which are different in the 2 images.  
Percent is the percentage of pixels which are different in the images.  
RedGreen is an image which has the different pixels rendered in red color and the same pixels in green color.  
Faded is an image where different pixels are rendered in red color and pixels which are the same accross the 2 images are rendered as they are, only with lower alpha. 
  
If you would like to try a complete GUI program using my library you can give [ImageCompare](https://github.com/Z1-Gamer/ImageCompare) a try!
