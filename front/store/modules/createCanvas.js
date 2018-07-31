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
            imageData.data[dataIndex + 0] = value[x][y].Red;
            imageData.data[dataIndex + 1] = value[x][y].Green;
            imageData.data[dataIndex + 2] = value[x][y].Blue;
            imageData.data[dataIndex + 3] = max;
        }
      }
      ctx.putImageData(imageData, 0, 0);
    }
  }
  


 


  export default createCanvas