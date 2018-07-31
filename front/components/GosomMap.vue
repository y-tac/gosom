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
  },
  mounted: function () {
    axios.get('/map')
    .then(response => {
      this.$store.dispatch('som/setMap',response.data.SomMap)
    })
    .catch(e => {
      console.error('error:', e)
    })

  }

}
</script>
