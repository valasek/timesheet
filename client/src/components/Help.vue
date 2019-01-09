<template>
  <div>
    <h1>Timeheet</h1>
    <p>Self-hosted web application for weekly reporting your consulting hours on projects using selected rates.</p>

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

    <h2>Requirements</h2>

    <ul>
      <li>Linux, Windows or MacOS</li>
      <li>DB connection to PostgreSQL - configure in <code>timesheet.yaml</code></li>
      <li>Initialize the DB using command <code>timesheet db --clean</code></li>
      <li>Import initial data in csv format using command <code>timesheet db --load all</code></li>
      <ul>
        <li>(dates should be entered in ISO format YYYY-MM-DD HH:MM:SS)</li>
        <li>consultants, projects, rates, holidays, initial reported_records - optional</li>
      </ul>
    </ul>
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
