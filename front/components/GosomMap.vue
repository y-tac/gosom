<template>
  <div id="gosommap">
   <h1>{{ title }}</h1>
   <canvas id="gosomcanvas" style="paddling:0;"></canvas>
  </div>
</template>

<style>
#gosommap {
  display: inline-block;
  position: relative;
  overflow: hidden;
  height: 300px;
  width: 300px;
}

</style>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      info:"", title:"test"
    }
  },
  created () {
    axios.get('/map', {
      params: {
        page: 1,
        per_page: 6,
        query: 'nuxt.js'
      }
    })
    .then(response => {
      this.info = response.data
    })
    .catch(e => {
      console.error('error:', e)
    })
  },
  mounted: function () {
    console.log('created')
    this.$store.dispatch('createFunction')
  }

}
</script>
