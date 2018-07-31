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
}

</style>

<script>

export default {
  data () {
    return {
      info:"", title:"test"
    }
  },
  created () {
  },
  mounted: function () {
    var store = this.$store;
    store.dispatch('som/setMap')
    setInterval(function () {
      store.dispatch('som/setMap')
    }, 5000)

  }

}
</script>
