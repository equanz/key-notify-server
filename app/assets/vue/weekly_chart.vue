<script>
  import {HorizontalBar} from 'vue-chartjs'

  export default {
    extends: HorizontalBar,
    mounted() {
      this.renderChart({
        labels: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'], // day label
        datasets: compilePlotData(),
      }, { // options
        scales: {
          xAxes: [{
            stacked: true,
            ticks: {
              beginAtZero: true,
              max: 24 // max value in x-axes
            }
          }],
          yAxes: [{
            stacked: true
          }]
        },
        tooltips: {
          enabled: false // hide tooltips
        },
        legend: {
          display: false // hide legend
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
    return [
      {
        label: '',
        backgroundColor: '#333333',
        data: [24, 0, 0, 0, 0, 0]
      },
      {
        label: '',
        backgroundColor: '#333333',
        data: [0, 16, 0, 0, 0, 0]
      },
      {
        label: '',
        backgroundColor: '#f87979',
        data: [0, 3, 0, 0, 0, 0, 0]
      },
      {
        label: '',
        backgroundColor: '#333333',
        data: [0, 5, 0, 0, 0, 0]
      },
      {
        label: '',
        backgroundColor: '#333333',
        data: [0, 0, 24, 24, 24, 24, 24]
      }
    ]
  }
</script>
