import createCanvas from './modules/createCanvas'
import axios from 'axios'

export const actions = {
  setMap() {
    axios.get('/map')
    .then(response => {
      createCanvas(response.data.SomMap)
    })
    .catch(e => {
      console.error('error:', e)
    })

  }
}



