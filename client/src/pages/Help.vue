<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <h5>Timesheet</h5>
    <p>Self-hosted web application for weekly reporting your consulting hours on projects using selected rates.</p>

    <h5>Quick Start</h5>
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
    <p>
      Read <router-link to="/documentation">
        full documentation.
      </router-link>
    </p>
    <h5>Request New Feature</h5>
    <p>
      Do you like what you see but mising an important feature? <a href="https://docs.google.com/forms/d/e/1FAIpQLSeh5Gw8iH9DlLQN3tjF66rrDLn5dkPnbi3V6vSjKAu0QPXnag/viewform?usp=sf_link">Describe your need</a>
      and next business day we will let you know if and when it will be available.
    </p>
    <h5>Report a Bug</h5>
    <p>
      Did you find an issue? <a href="https://docs.google.com/forms/d/e/1FAIpQLSfExGW2_D3YI4SjQ3-7SC7gys-b1AwywY-TBmITlDX644-Upg/viewform?usp=sf_link">Report it</a>
      and next business day we will let you know when the fix will be available.
    </p>
    <q-img
      src="statics/settings-image.jpg"
      class="image"
    />
    <my-footer />
  </q-page>
</template>

<script>

export default {

  components: {
    /* webpackChunkName: "core" */
    'my-footer': () => import('components/MyFooter')
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Help')
    this.$store.commit('context/SET_PAGE_ICON', 'help_outline')
  }
}
</script>

<style scoped>
.image {
  border-radius: 10px;
}
</style>
