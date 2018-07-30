import Vuex from 'vuex'
import createCanvas from './modules/createCanvas'

const store = new Vuex.Store({
  state: {
  },
  mutations: {
  },
  actions: {
    createFunction () {
      createCanvas()
    }
  }
})

export default store