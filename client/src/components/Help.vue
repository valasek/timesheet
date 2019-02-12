<template>
  <div>
    <h1>TimeSheet</h1>
    <p>Self-hosted web application for weekly reporting your consulting hours on projects using selected rates.</p>

    <p>
      Report an <a href src="https://github.com/valasek/timesheet/issues">
        issue
      </a> or request a <a href src="https://github.com/valasek/timesheet/issues">
        new feature
      </a>.
    </p>

    <p>
      Donate to cover the spent time and support further development via <a src="https://www.patreon.com/valasek">
        Patreon
      </a> or <a src="https://paypal.me/StanislavValasek">
        PayPal
      </a>.
    </p>

    <h2>Rationale</h2>

    <p>Automation is based on a premise, that reporting and billing process includes three separated steps with well-defined data which flows in between:</p>

    <h3>1. Time reporting</h3>
    <p><v-label>All company employees</v-label></p>
    <p>
      <strong>Covered by this app</strong><br>
      Data are safe, eliminates manual errors, shows important data for consultants, data entry is as easy as possible.
    </p>

    <h3>2. Billing preparation</h3>
    <p><v-label>Project manager or administrative person</v-label></p>
    <p>
      Confirming/editing reported hours for billing.
      <br><strong>Covered partially by this app</strong>. If you bill 1:1 with your reporting then it has all you need.<br>
      Data are available in the DB which can be exported by this app read or exported by any DB tool.
    </p>

    <h3>3. Exporting reported and billed hours</h3>
    <p>
      In the internal accounting system for the invoicing to the clients and inputs for consultant's salaries.
      <br><strong>Not covered by this app.</strong><br>
      Data are available in the DB which can be exported by this app read or exported by any DB tool.
    </p>
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
</style>
