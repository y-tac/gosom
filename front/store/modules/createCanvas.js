const createCanvas = (value) => {
    console.log('export')
    draw(value)
  }
  
  function draw (value) {
    const canvasWidth = value.length
    const canvasHeight = value.length
  
    const canvas = document.getElementById('gosomcanvas')
    canvas.width = canvasWidth
    canvas.height = canvasHeight
  
    if (canvas.getContext) {
      const ctx = canvas.getContext('2d')
      var imageData = ctx.createImageData(canvasWidth, canvasHeight);
      for (let x = 0; x < canvasWidth; x++) {
        for (let y = 0; y < canvasWidth; y++) {
            var pixelIndex = x + y * canvasWidth;
            var dataIndex = pixelIndex * 4;
            var max = 255;
            var cmax = 0xff;
            var rgb = [0x81 , 0xa7 ,0xf3];
            var rgbMax = Math.max.apply(null,rgb);
            var rgbMin = Math.min.apply(null,rgb);
            var fixs = cmax - rgbMax + rgbMin;
            imageData.data[dataIndex + 0] = fixs * value[x][y].Red/max + rgb[0];
            imageData.data[dataIndex + 1] = fixs * value[x][y].Green/max + rgb[1] ;
            imageData.data[dataIndex + 2] = fixs * value[x][y].Blue/max + rgb[2]  ;
            imageData.data[dataIndex + 3] = max;
        }
      }
      ctx.putImageData(imageData, 0, 0);
    }
  }
  


 


  export default createCanvas