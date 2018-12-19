<template>
  <v-app>
    <v-navigation-drawer app>
      <v-toolbar flat>
        <v-list>
          <v-list-tile>
            <v-list-tile-title class="title">
              Time sheets
            </v-list-tile-title>
          </v-list-tile>
        </v-list>
      </v-toolbar>

      <v-divider />

      <v-list dense class="pt-0">
        <v-list-tile v-for="item in items" :key="item.title" :to="item.route">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>

          <v-list-tile-content>
            <v-list-tile-title>{{ item.title }}</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>

      <v-divider />

      <v-list>
        <v-list-tile>
          <div class="header ">
            Change consultant
          </div>
        </v-list-tile>
        <v-list-tile>
          <v-select :items="consultants" item-text="name" item-value="name" @change="changeConsultant" />
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app>
      <v-toolbar-title class="headline">
        <span>{{ page }}</span>
      <!-- <span class="font-weight-light">  management</span> -->
      </v-toolbar-title>
      <v-spacer />
      <v-btn flat to="help">
        <span class="mr-2">
          Help
        </span>
      </v-btn>
    </v-toolbar>
    <v-content>
      <v-container fluid>
        <router-view>
          <report-table />
        </router-view>
      </v-container>
    </v-content>
    <v-footer app>
      <v-flex text-xs-center>
        2018 &copy; <strong>DataArch</strong>
      </v-flex>
    </v-footer>
  </v-app>
</template>

<script>
  import ReportTable from './ReportTable'
  import { mapState } from 'vuex'

  export default {
    name: 'App',
    components: {
      ReportTable
    },

    data () {
      return {
        items: [
          { title: 'Report my work', icon: 'dashboard', route: 'report' },
          { title: 'Show reported work', icon: 'question_answer', route: 'reported' }
        ],
        right: null
      }
    },

    computed: {
      ...mapState({
        consultants: state => state.consultants.all,
        // selectedConsultants: state => state.consultants.selected,
        selectedConsultants: {
          set (newValue) {
            this.$store.dispatch('consultants/setSelected', newValue)
          },
          get () {
            return this.consultants.selected
          }
        },
        page: state => state.context.page
      })
    },

    created () {
      this.$store.dispatch('consultants/getConsultants')
      this.$store.dispatch('projects/getProjects')
      this.$store.dispatch('rates/getRates')
      let month = (new Date().getMonth() + 1).toString()
      this.$store.dispatch('reportedHours/getReportedHours', month)
    },

    methods: {
      changeConsultant (ids) {
        console.log('changeConsultant ' + ids) /* eslint-disable-line no-console */
        this.$store.dispatch('consultants/setSelected', ids)
      }
    }
  }
</script>
