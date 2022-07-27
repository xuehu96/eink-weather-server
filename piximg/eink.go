package piximg

import "image"

//GenerateEinkBytes 生成最终的图片序列
func GenerateEinkBytes(img image.Image) []byte {
	einkData := make([]byte, 200*25)
	for y := 0; y < 200; y++ {
		for l := 0; l < 25; l++ {
			for i := 0; i < 8; i++ {
				r, _, _, _ := img.At(l*8+i, y).RGBA()
				einkData[y*25+l] >>= 1
				if r >= 127 {
					einkData[y*25+l] |= 0x80
				}
			}
		}
	}
	return einkData
}

/*
//生成最终的图片序列
fn generate_eink_bytes(img: &GrayImage)->Vec<u8>{
    let mut r:Vec<u8> = Vec::with_capacity((HEIGHT*WIDTH/8) as usize);//存结果
    for y in 0..HEIGHT {
        for l in 0..WIDTH/8 {
            let mut temp:u8 = 0;
            for i in 0..8 {
                let p:u8 = img.get_pixel(l*8+i,y)[0];
                //匹配像素点颜色
                let t = if p > 127 {0}else{1};
                temp+=t<<i;
            }
            r.push(temp);
        }
    }
    r
}
**/
