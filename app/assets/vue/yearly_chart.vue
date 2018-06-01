<script>
  import {Bar} from 'vue-chartjs'

  export default {
    extends: Bar,
    props: ['raw_data'],
    mounted() {
      this.renderChart({
        datasets: compilePlotData(this.raw_data),
      }, { // options
        tooltips: {
          enabled: false
        },
           legend: {
             display: false
           }
      })
    }
  }

  /**
   * compile to plot data from raw data   (caution: not implemented)
   * @param raw_data {Array} - raw data about status
   * @return {Array} return plot data
   */
  let compilePlotData = (raw_data) => {
    let plot_data = []
    for(let i = 0; i < 52; i++){
      plot_data.push(
        {
          label: 'foo',
          backgroundColor: '#f87979',
          data: [50]
        }
      )
    }
    return plot_data
  }
</script>
