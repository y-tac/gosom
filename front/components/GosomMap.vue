<template>
  <div id="gosommap">
   <h1>{{ title }}</h1>
   {{ info }}
  </div>
</template>

<style>
#gosommap {
  display: inline-block;
  animation: turn 2s linear forwards 1s;
  transform: rotateX(180deg);
  position: relative;
  overflow: hidden;
  height: 245px;
  width: 245px;
}

</style>

<script>
export default {
  async asyncData ({ params }) {
    let { data } = await axios.get(`http://localhost:3306/map`)
    return { info: data }
  }

}
</script>
