<template>
  <div id="app">
    <Status :room_name=room_name :from_time=from_time :status=room_status></Status>

    <article>
      <Section :section_name=section_name.weekly>
        <WeeklyChart :raw_data=weekly_raw_data :last_week_state=last_week_data.State></WeeklyChart>
      </Section>

      <Section :section_name=section_name.yearly>
        <YearlyChart :raw_data=yearly_raw_data :fiscal_year=fiscal_year></YearlyChart>
      </Section>
    </article>
  </div>
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
    data(){
      return {
        // test data
        weekly_raw_data: [{"Time":"2018-05-28 15:52:56","State":"ON","Key_info_id":1},{"Time":"2018-05-28 17:43:55","State":"OFF","Key_info_id":2},{"Time":"2018-05-29 15:52:56","State":"ON","Key_info_id":3},{"Time":"2018-05-29 17:43:55","State":"OFF","Key_info_id":4},{"Time":"2018-05-30 15:52:56","State":"ON","Key_info_id":5},{"Time":"2018-05-30 17:43:55","State":"OFF","Key_info_id":6},{"Time":"2018-05-31 15:52:56","State":"ON","Key_info_id":7},{"Time":"2018-05-31 17:43:55","State":"OFF","Key_info_id":8},{"Time":"2018-06-01 16:50:06","State":"ON","Key_info_id":9}],
        last_week_data: {"Time":"2018-06-01 16:50:06","State":"ON","Key_info_id":9},
        yearly_raw_data: [{"Time":"2018-04-25 15:52:56","State":"ON","Key_info_id":1},{"Time":"2018-04-25 17:43:55","State":"OFF","Key_info_id":2},{"Time":"2018-05-29 15:52:56","State":"ON","Key_info_id":3},{"Time":"2018-05-29 17:43:55","State":"OFF","Key_info_id":4},{"Time":"2018-05-30 15:52:56","State":"ON","Key_info_id":5},{"Time":"2018-05-30 17:43:55","State":"OFF","Key_info_id":6},{"Time":"2018-05-31 15:52:56","State":"ON","Key_info_id":7},{"Time":"2018-05-31 17:43:55","State":"OFF","Key_info_id":8},{"Time":"2018-06-01 16:50:06","State":"ON","Key_info_id":9}],
        fiscal_year: 2018,
        room_name: "Server Room",
        from_time: "14:30",
        room_status: "ON",
        section_name: {
          weekly: 'Weekly Status',
          yearly: 'Yearly Status'
        }
      }
    }
  }
</script>

<style lang="less">
  @import "../less/main.less";
  @import url(//fonts.googleapis.com/earlyaccess/notosansjp.css);

  body{
    margin: 0;
    color: @text-black-color;
    font-family: 'Noto Sans JP', sans-serif;
    background-color: @back-color;
  }

  div#app{
    padding-top: @nav-height;
  }

  article{
    display: flex;
  }
</style>
