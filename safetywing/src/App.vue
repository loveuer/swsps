<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
}
</script>

<style>
#app {
  width: 100%;
  min-height: 100%;
}
.ql-size-small {
  font-size: 0.75em;
}
.ql-size-huge {
  font-size: 2.5em;
}
.ql-size-large {
  font-size: 1.5em;
}
</style>
