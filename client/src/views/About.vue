<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <h1>Timesheet</h1>
    <p>Self-hosted web application for weekly reporting your consulting hours on projects using selected rates.</p>

    <h2>Quick Start</h2>
    <p>
      <ol>
        <li>
          Update <router-link to="/administration">
            Application settings
          </router-link>, <router-link to="administration">
            Vacation settings
          </router-link> and <router-link to="administration">
            Warning limits
          </router-link> to meet your needs
        </li>
        <li>
          Use <router-link to="/administration">
            Backup & Restore
          </router-link>
          to export demo data. Update csv files and import back your projects, rates, consultants and holidays.
        </li>
        <li>
          Enjoy. You are <router-link to="/">
            good to go
          </router-link>
        </li>
      </ol>
    </p>
    <h2>Documentation</h2>
    <p>
      Read <router-link to="/documentation">
        full documentation.
      </router-link>
    </p>
    <h2>Request New Feature</h2>
    <p>
      Do you like what you see but mising an important feature? <a href="mailto:valasek@gmail.com">Describe your need</a>
      and next business day we will let you know if and when it will be available.
    </p>
    <v-img
      src="./settings-image.jpg"
      class="grey lighten-2 image"
    />
    <br>
    <p>2018 - {{ (new Date()).getFullYear() }} &copy; <strong><a href="https://valasek.wordpress.com/">Stanislav Valasek</a></strong> Version {{ version }}, server: {{ url }}:{{ port }}</p>
  </div>
</template>

<script>
  import { mapState } from 'vuex'

  export default {

    data () {
      return {
        url: process.env.VUE_APP_URL,
        port: process.env.VUE_APP_PORT
      }
    },

    computed: {
      ...mapState({
        version: state => state.settings.version
      })
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'About')
    }
  }
</script>

<style scoped>
.image {
  border-radius: 10px;
}
</style>
