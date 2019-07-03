<!-- Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com> -->

<template>
  <q-page padding>
    <div class="q-pa-md row items-start q-gutter-md">
      <q-card
        class="text-white"
        style="background: radial-gradient(circle, #35a2ff 0%, #014a88 100%)">
        <q-card-section>
          <div class="text-h6">Timesheet</div>
          <div class="text-subtitle2">Self-hosted web application for weekly reporting</div>
        </q-card-section>
        <q-card-section>
          Report your consulting hours on projects with further segmentation via rates.<br/>
          You can download all data in csv format, modify as required and import modified data.<br/>
          Timesheet provides API to automate data exports.
        </q-card-section>
      </q-card>
    </div>
    <div class="text-h6">
      Top 10 projects in {{ month }} {{ year }}
      <q-btn flat icon="skip_previous" @click="previousMonth" />
      <q-btn flat icon="skip_next" @click="nextMonth" />
    </div>
    <div class="row items-top">
      <div class="col-6">
        <project-chart :chartData="projectChart" :options="projectChartOptions"></project-chart>
      </div>
      <q-card flat>
        <q-card-section>
          <q-table :columns="columns" :data="topProjects" :pagination="myPagination" hide-bottom dense flat/>
        </q-card-section>
      </q-card>
    </div>
    <div class="row items-center">
      <div class="col-4">
        <q-card flat>
          <q-card-section>
            <div class="text-h6">Managed data</div>
          </q-card-section>
          <q-card-section>
            <project-chart :chartData="dataChart" :options="projectChartOptions"></project-chart>
          </q-card-section>
        </q-card>
      </div>
    </div>
    <my-footer/>
  </q-page>
</template>

<script>
import { mapState } from 'vuex'
import { format, addMonths, subMonths } from 'date-fns'

export default {

  components: {
    /* webpackChunkName: "core" */
    'my-footer': () => import('components/MyFooter'),
    'project-chart': () => import('components/ProjectChart')
  },

  data () {
    return {
      url: process.env.VUE_APP_URL,
      port: process.env.VUE_APP_PORT,
      columns: [
        { label: 'Project', align: 'left', field: 'project', sortable: false },
        { label: 'Hours', align: 'left', field: 'hours', sortable: false }
      ],
      myPagination: { 'rowsPerPage': 50 }
    }
  },

  computed: {
    ...mapState({
      version: state => state.settings.version,
      projects: state => state.projects.all,
      rates: state => state.rates.all,
      consultants: state => state.consultants.all,
      selectedMonth: state => state.settings.selectedMonth,
      reportedHoursSummary: state => state.reportedHours.summary,
      assignedProjects: state => state.projects.all
    }),
    projectChart () {
      return {
        labels: this.topProjects.map(p => p.project),
        datasets: [
          {
            label: 'Total reported hours',
            backgroundColor: '#66BB6A',
            data: this.topProjects.map(p => p.hours)
          }
        ]
      }
    },
    projectChartOptions () {
      return {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          xAxes: [{
            gridLines: {
              display: false
            }
          }]
        },
        plugins: {
          datalabels: {
            color: '#FFFFFF',
            anchor: 'center',
            clamp: true
          }
        }
      }
    },
    dataChart () {
      return {
        labels: ['Projects', 'Rates', 'Consultants'],
        datasets: [
          {
            label: '# of managed records',
            backgroundColor: '#66BB6A',
            data: [this.assignedProjects.length, this.rates.length, this.consultants.length]
          }
        ]
      }
    },
    month () {
      return format(this.selectedMonth, 'MMMM')
    },
    year () {
      return format(this.selectedMonth, 'yyyy')
    },
    topProjects () {
      var allProjects = this.reportedHoursSummary.map(record => ({ project: record.project, month: record.month, year: record.year, hours: parseFloat(record.hours) }))
      const m = format(this.selectedMonth, 'M')
      const y = format(this.selectedMonth, 'yyyy')
      var inputProjects = allProjects.filter(function (obj) {
        return obj.month === m && obj.year === y
      })
      var projectTotal = []
      this.assignedProjects.forEach(function (p) {
        projectTotal.push({ project: p.name, hours: 0 })
      })
      inputProjects.forEach(function (p) {
        const index = projectTotal.findIndex(obj => obj.project === p.project)
        if (index > 0) {
          projectTotal[index].hours = projectTotal[index].hours + p.hours
        }
      })
      projectTotal.sort((a, b) => (a.hours < b.hours) ? 1 : -1)
      return projectTotal.slice(0, 10)
    }
  },

  created () {
    this.$store.commit('context/SET_PAGE', 'Timesheet')
    this.$store.commit('context/SET_PAGE_ICON', 'home')
    this.$store.dispatch('reportedHours/getYearlySummary', this.selectedMonth)
  },

  methods: {
    previousMonth () {
      const d = subMonths(this.selectedMonth, 1)
      this.$store.dispatch('settings/jumpToWeek', d)
    },
    nextMonth () {
      const d = addMonths(this.selectedMonth, 1)
      this.$store.dispatch('settings/jumpToWeek', d)
    }
  }

}
</script>

<style scoped>
</style>
