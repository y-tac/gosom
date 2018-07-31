import createCanvas from './modules/createCanvas'

export const actions = {
  setMap(context, value) {
    createCanvas(value)
  }
}



