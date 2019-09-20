<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <q-page-sticky
      position="top-right"
      :offset="[18, 18]"
    >
      <q-list
        dense
        bordered
        padding
        class="rounded-borders"
      >
        <q-item
          v-for="(header, id) in toc"
          :key="id"
          v-ripple
          clickable
        >
          <q-item-section>
            <a :href="getLink(header.id)">{{ header.label }}</a>
          </q-item-section>
        </q-item>
      </q-list>
    </q-page-sticky>
    <q-markdown
      :src="markdown"
      toc
      style="width:80%"
      @data="onToc"
    />
  </q-page>
</template>

<script>
import api from '../api/axiosSettings'

export default {
  data () {
    return {
      markdown: 'Loading documentation ...',
      toc: {}
    }
  },

  computed: {},

  created () {
    this.$store.commit('context/SET_PAGE', 'Documentation')
    this.$store.commit('context/SET_PAGE_ICON', 'help_outline')
    this.getDocumentation()
  },

  methods: {
    getDocumentation () {
      const comp = this
      api.apiClient
        .get('/api/download/docs')
        .then(response => {
          comp.markdown = response.data
        })
        .catch(function (e) {
          this.$q.notify({
            message: "Couldn't find documentation: " + e.response.data,
            color: 'negative',
            icon: 'report_problem'
          })
          console.log(e, e.response) /* eslint-disable-line no-console */
        })
    },
    onToc (toc) {
      this.toc = toc.filter(u => {
        return u.level === 1
      })
    },
    getLink (id) {
      return '/documentation#' + id
    }
  }
}
</script>

<style scoped>
</style>
