<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <div>
    <h1>Timesheet</h1>
    <p>Self-hosted web application for weekly reporting your consulting hours on projects using selected rates.</p>
    <p>2018 - {{ (new Date()).getFullYear() }} &copy; <strong>Stanislav Valasek</strong> Version {{ version }}</p>

    <!-- <p>
      Report an <a href src="https://github.com/valasek/timesheet/issues">
        issue
      </a> or request a <a href src="https://github.com/valasek/timesheet/issues">
        new feature
      </a>.
    </p> -->

    <h2>How to Start</h2>

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
          You are <router-link to="/">
            good to go
          </router-link>
        </li>
      </ol>
    </p>
    <h2>Commercial usage</h2>
    <p>
      Pay for the commertial usage of the application to support further development and maintenance via <a href="https://www.patreon.com/valasek">
        Patreon
      </a> or <a href="https://paypal.me/StanislavValasek">
        PayPal
      </a>.
    </p>
    <v-img
      src="./settings-image.jpg"
      class="grey lighten-2 image"
    />
  </div>
</template>

<script>
  export default {
    name: 'Help',
    data () {
      return {
      }
    },

    created () {
      this.$store.commit('context/SET_PAGE', 'Help')
    }
  }
</script>

<style scoped>
.image {
  border-radius: 10px;
}
</style>
