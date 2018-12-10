<template>
  <div></div>
</template>

<script>
export default {
  name: 'Help',
  data () {
    return {
    }
  },

  created () {
    this.$store.commit('context/setPage', "Help")
  }
}
</script>

<style scoped>
</style>