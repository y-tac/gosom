const createCanvas = () => {
    console.log('export')
    draw()
  }
  
  function draw () {
    const canvasWidth = 400
    const canvasHeight = 300
  
    const canvas = document.getElementById('tutorial')
    canvas.width = canvasWidth
    canvas.height = canvasHeight
  
    if (canvas.getContext) {
      const ctx = canvas.getContext('2d')
      // 全体の透明度
      ctx.globalAlpha = 0.5
      // 丸の数の乱数を取得
      let roundMin = 8
      let roundMax = 20
      const roundNum = Math.floor( Math.random() * roundMax + 1 - roundMin ) + roundMin
      for (let i = 0; i < roundNum; i++) {
        drowRound(ctx, canvasWidth, canvasHeight)
      }
    }
  }
  
  function drowRound (ctx, canvasWidth, canvasHeight) {
    // 色の乱数
    const r = Math.floor(Math.random() * 256)
    const g = Math.floor(Math.random() * 256)
    const b = Math.floor(Math.random() * 256)
  
    ctx.fillStyle = 'rgb(' + r + ',' + g + ',' + b + ')'
  
    // 丸の数の乱数を取得
    const roundMin = 5
    const roundMax = 20
    const roundNum = Math.floor( Math.random() * roundMax + 1 - roundMin ) + roundMin
    const roundX = Math.random() * canvasWidth
    const roundY = Math.random() * canvasHeight
    const radiusMin = 10
    const radiusMax = 100
    const radiusNum = Math.floor( Math.random() * radiusMax + 1 - radiusMin ) + radiusMin
  
    ctx.beginPath()
    ctx.arc(roundX, roundY, radiusNum, 0, Math.PI*2.0,true)
    ctx.fill()
  }
  
  export default createCanvas