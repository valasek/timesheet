import { Bar, mixins } from 'vue-chartjs'
import ChartDataLabels from 'chartjs-plugin-datalabels'

const { reactiveProp } = mixins

export default {
  extends: Bar,
  mixins: [reactiveProp],
  plugins: [ChartDataLabels],
  props: ['options'],

  mounted () {
    this.renderChart(this.chartData, this.options)
  }
}
