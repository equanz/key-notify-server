<template id="app">
  <article id="app">
    <div id="status">
      <Status :room_name=room_name :from_time=from_time :status=room_status></Status>
    </div>

    <div id="contents">
      <Section :section_name=section_name.weekly>
        <WeeklyChart :raw_data=weekly_raw_data :last_week_state=last_week_data.State :select_date=select_date.weekly></WeeklyChart>
      </Section>

      <Section :section_name=section_name.yearly>
        <YearlyChart :raw_data=yearly_raw_data :fiscal_year=fiscal_year :select_date=select_date.yearly></YearlyChart>
      </Section>
    </div>
  </article>
</template>

<script>
  import WeeklyChart from './weekly_chart.vue'
  import YearlyChart from './yearly_chart.vue'
  import Status from './status.vue'
  import Section from './section.vue'
  import APIGet from '../js/api_get.js'
  export default {
    components:
    { // add to local component
      WeeklyChart: WeeklyChart,
      YearlyChart: YearlyChart,
      Status: Status,
      Section: Section
    },
    data() {
      return {
        weekly_raw_data: [],
        last_week_data: {},
        yearly_raw_data: [],
        fiscal_year: 0,
        room_name: "Server Room",
        from_time: "",
        room_status: "",
        section_name: {
          weekly: 'Weekly Status',
          yearly: 'Yearly Status'
        },
        select_date: {
          weekly: '',
          yearly: '',
        },
        update_interval: 60 * 60 * 1000 // 60 min
      }
    },
    methods: {
      updateWeekly: function(date) {
        APIGet.GetStatistics_Week(date).then((res) => {
          // update raw_data
          // not in res
          if(! res){
            this.weekly_raw_data = []
          } else{ // in res
            this.weekly_raw_data = res
          }
        }).catch(() => {
          console.log('caught some error!')
        })
      },
      updateYearly: function(date) {
        APIGet.GetStatistics_Year(date).then((res) => {
          // update fiscal year
          this.fiscal_year = this.calcFiscalYear(date)
          // update raw_data
          // not in res
          if(! res){
            this.yearly_raw_data = []
          } else{ // in res
            this.yearly_raw_data = res
          }
        }).catch(() => {
          console.log('caught some error!')
        })
      },
      updateLast: function() {
        APIGet.GetStatistics_Last().then((res) => {
          let last_date = new Date(res.Time)
          // update last data
          this.room_status = res.State
          this.from_time = `${last_date.getFullYear()}/${this.toZeroFill(last_date.getMonth() + 1, 2)}/${this.toZeroFill(last_date.getDate(), 2)} ${this.toZeroFill(last_date.getHours(), 2)}:${this.toZeroFill(last_date.getMinutes(), 2)}`
        }).catch(() => {
          console.log('caught some error!')
        })
      },
      updateLastWeek: function(date) {
        APIGet.GetStatistics_Before(date).then((res) => {
          // update last week data
          this.last_week_data = res
        }).catch(() => {
          console.log('caught some error!')
        })
      },
      calcFiscalYear: function(date) { // calculate fiscal year
        // Jan-Mar
        if(date.getMonth() < 3){
          return date.getFullYear() - 1
        } else{
          return date.getFullYear()
        }
      },
      updateAll: function(date) { // update all data
        this.select_date.weekly = date
        this.select_date.yearly = date
        this.updateWeekly(date)
        this.updateYearly(date)
        this.updateLast()
        this.updateLastWeek(date)
      },
      toZeroFill: function(num, place) { // zero fill by place value
        const fill_code = '0'
        if (num != 0){ // num is not 0
          let num_place = Math.floor(Math.log10(num)) + 1 // calc place

          if (num_place < place){
            // zero fill
            return `${fill_code.repeat(place - num_place)}${num.toString()}`
          } else{
            // not fill
            return num.toString()
          }
        } else{ // num is 0
          return `${fill_code.repeat(place)}`
        }
      }
    },
    created() {
      let now = new Date() // last week, year only
      this.updateAll(now)

      // time update event
      setInterval(() => {
        let now = new Date() // last week, year only
        this.updateAll(now)
      }, this.update_interval)
    }
  }
</script>

<style lang="less">
  @import "../less/main.less";
  @import url(https://fonts.googleapis.com/earlyaccess/notosansjp.css);

  body{
    margin: 0;
    color: @text-black-color;
    font-family: 'Noto Sans JP', sans-serif;
    background-color: @back-color;
  }

  article#app{
    padding-top: @nav-height + 20px;
  }

  div#status{
    display: flex;
    justify-content: center;
  }

  div#contents{
    margin-top: 40px;
    display: flex;
    justify-content: space-around;
    flex-direction: column;
    align-items: center;
    section{
      margin-top: 20px;
    }

    .screen-sm({
      flex-direction: row;
      section{
        margin-top: 0px;
      }
    });

    .screen-md({
      flex-direction: row;
      section{
        margin-top: 0px;
      }
    });

    .screen-lg({
      flex-direction: row;
      section{
        margin-top: 0px;
      }
    });
  }
</style>
